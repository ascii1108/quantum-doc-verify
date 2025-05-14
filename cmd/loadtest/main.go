package main

import (
    "crypto/rand"
    "flag"
    "fmt"
    "os"
    "path/filepath"
    "sync"
    "time"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    
    "quantum-doc-verify/pkg/crypto"
    "quantum-doc-verify/pkg/storage"
)

func main() {
    // Configure logging
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})

    // Parse command line flags
    concurrency := flag.Int("concurrency", 10, "Number of concurrent uploads")
    totalUploads := flag.Int("total", 100, "Total number of documents to upload")
    documentSize := flag.Int("size", 10, "Size of test documents in KB")
    ipfsGateway := flag.String("ipfs", "localhost:5001", "IPFS API gateway")
    signDocuments := flag.Bool("sign", true, "Sign documents with Dilithium (slower but realistic)")
    flag.Parse()

    log.Info().
        Int("concurrency", *concurrency).
        Int("totalUploads", *totalUploads).
        Int("documentSizeKB", *documentSize).
        Bool("signDocuments", *signDocuments).
        Str("ipfsGateway", *ipfsGateway).
        Msg("Starting load test")

    // Create IPFS client
    ipfs, err := storage.NewIPFSClient(*ipfsGateway)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create IPFS client")
    }

    // Generate a Dilithium keypair if signing is enabled
    var (
        signer  *crypto.DilithiumSigner
        privKey []byte
    )
    
    if *signDocuments {
        signer = crypto.NewDilithiumSigner()
        pubKey, privateKey, err := signer.GenerateKeypair()
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to generate Dilithium keypair")
        }
        privKey = privateKey
        
        // Save keys to temporary directory for reference
        tempDir := os.TempDir()
        pubKeyPath := filepath.Join(tempDir, "loadtest_dilithium_public.key")
        privKeyPath := filepath.Join(tempDir, "loadtest_dilithium_private.key")
        
        err = signer.SaveKeys(pubKey, privateKey, pubKeyPath, privKeyPath)
        if err != nil {
            log.Warn().Err(err).Msg("Failed to save keys but continuing with in-memory keys")
        } else {
            log.Info().
                Str("publicKey", pubKeyPath).
                Str("privateKey", privKeyPath).
                Msg("Dilithium keys saved")
        }
    }

    // Initialize counters and metrics
    var (
        wg             sync.WaitGroup
        mutex          sync.Mutex
        successCount   int
        failureCount   int
        totalIPFSTime  time.Duration
        totalSignTime  time.Duration
        minUploadTime  = time.Hour
        maxUploadTime  time.Duration
        uploadTimes    = make([]time.Duration, 0, *totalUploads)
    )

    // Create a channel to control concurrency
    semaphore := make(chan struct{}, *concurrency)
    
    // Generate base document content
    baseContent := generateRandomDocument(*documentSize * 1024)
    
    // Start the timer for the entire test
    startTime := time.Now()

    // Perform uploads
    for i := 0; i < *totalUploads; i++ {
        wg.Add(1)
        semaphore <- struct{}{} // Acquire semaphore
        
        go func(uploadNum int) {
            defer wg.Done()
            defer func() { <-semaphore }() // Release semaphore
            
            // Create a slightly different document for each upload
            docContent := modifyContent(baseContent, uploadNum)
            
            // Start timing this upload
            uploadStart := time.Now()
            
            // Sign document if requested
            var signTime time.Duration
            
            if *signDocuments {
                signStart := time.Now()
                
                // Create a temporary file for the document
                tempFile, err := os.CreateTemp("", fmt.Sprintf("loadtest-doc-%d-*.txt", uploadNum))
                if err != nil {
                    log.Error().Err(err).Int("uploadNum", uploadNum).Msg("Failed to create temp file")
                    mutex.Lock()
                    failureCount++
                    mutex.Unlock()
                    return
                }
                tempFilePath := tempFile.Name()
                defer os.Remove(tempFilePath)
                
                // Write document to temp file
                _, err = tempFile.Write(docContent)
                tempFile.Close() // Close file after writing
                if err != nil {
                    log.Error().Err(err).Int("uploadNum", uploadNum).Msg("Failed to write document to temp file")
                    mutex.Lock()
                    failureCount++
                    mutex.Unlock()
                    return
                }
                
                // Sign the document (don't store signature in a variable since we're not using it)
                _, err = signer.SignDocument(tempFilePath, privKey)
                if err != nil {
                    log.Error().Err(err).Int("uploadNum", uploadNum).Msg("Failed to sign document")
                    mutex.Lock()
                    failureCount++
                    mutex.Unlock()
                    return
                }
                
                signTime = time.Since(signStart)
            }
            
            // Upload to IPFS
            ipfsStart := time.Now()
            cid, err := ipfs.Store(docContent)
            ipfsTime := time.Since(ipfsStart)
            
            // Calculate total upload time
            uploadTime := time.Since(uploadStart)
            
            // Update statistics
            mutex.Lock()
            defer mutex.Unlock()
            
            if err != nil {
                log.Error().
                    Err(err).
                    Int("uploadNum", uploadNum).
                    Dur("uploadTime", uploadTime).
                    Msg("Upload failed")
                failureCount++
            } else {
                log.Info().
                    Int("uploadNum", uploadNum).
                    Str("cid", cid).
                    Dur("totalTime", uploadTime).
                    Dur("signTime", signTime).
                    Dur("ipfsTime", ipfsTime).
                    Int("size", len(docContent)).
                    Bool("signed", *signDocuments).
                    Msg("Upload successful")
                
                successCount++
                totalIPFSTime += ipfsTime
                totalSignTime += signTime
                uploadTimes = append(uploadTimes, uploadTime)
                
                if uploadTime < minUploadTime {
                    minUploadTime = uploadTime
                }
                if uploadTime > maxUploadTime {
                    maxUploadTime = uploadTime
                }
            }
        }(i)
    }

    // Wait for all uploads to complete
    wg.Wait()
    
    // Calculate total test duration
    totalTime := time.Since(startTime)
    
    // Calculate averages and throughput
    var avgUploadTime, avgIPFSTime, avgSignTime time.Duration
    if successCount > 0 {
        avgUploadTime = time.Duration(int64(totalTime) / int64(successCount))
        avgIPFSTime = time.Duration(int64(totalIPFSTime) / int64(successCount))
        if *signDocuments {
            avgSignTime = time.Duration(int64(totalSignTime) / int64(successCount))
        }
    }
    
    // Calculate the median upload time
    median := calculateMedian(uploadTimes)
    
    // Calculate the 95th percentile upload time
    p95 := calculatePercentile(uploadTimes, 95)
    
    // Print results
    fmt.Println("\n------- QUANTUM-DOC-VERIFY LOAD TEST RESULTS -------")
    fmt.Printf("Total time: %v\n", totalTime)
    fmt.Printf("Concurrency level: %d\n", *concurrency)
    fmt.Printf("Document size: %d KB\n", *documentSize)
    fmt.Printf("Documents signed: %v\n", *signDocuments)
    fmt.Printf("Total uploads attempted: %d\n", *totalUploads)
    fmt.Printf("Successful uploads: %d (%.1f%%)\n", successCount, float64(successCount)*100/float64(*totalUploads))
    fmt.Printf("Failed uploads: %d (%.1f%%)\n", failureCount, float64(failureCount)*100/float64(*totalUploads))
    
    if successCount > 0 {
        fmt.Printf("\nPerformance metrics:\n")
        fmt.Printf("Average upload time: %v\n", avgUploadTime)
        fmt.Printf("Median upload time: %v\n", median)
        fmt.Printf("95th percentile upload time: %v\n", p95)
        fmt.Printf("Min upload time: %v\n", minUploadTime)
        fmt.Printf("Max upload time: %v\n", maxUploadTime)
        fmt.Printf("Average IPFS storage time: %v\n", avgIPFSTime)
        
        if *signDocuments {
            fmt.Printf("Average Dilithium signing time: %v\n", avgSignTime)
        }
        
        fmt.Printf("\nThroughput: %.2f documents/second\n", float64(successCount)/totalTime.Seconds())
        fmt.Printf("Throughput: %.2f KB/second\n", float64(successCount*(*documentSize))/totalTime.Seconds())
    }
    fmt.Println("----------------------------------------------------")
}

// generateRandomDocument creates a random document of specified size
func generateRandomDocument(size int) []byte {
    const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789 .,;:!?-()[]{}'\"\n\t"
    
    doc := make([]byte, size)
    _, err := rand.Read(doc)
    if err != nil {
        // Fall back to a less random approach if crypto/rand fails
        for i := 0; i < size; i++ {
            doc[i] = chars[i%len(chars)]
        }
    }
    
    return doc
}

// modifyContent creates a unique document by modifying the base content
func modifyContent(base []byte, id int) []byte {
    result := make([]byte, len(base))
    copy(result, base)
    
    // Add a unique header so each document is different
    header := fmt.Sprintf("Document ID: %d\nTimestamp: %s\n\n", id, time.Now().Format(time.RFC3339))
    
    // Ensure we don't exceed original size
    if len(header) < len(result) {
        copy(result, []byte(header))
    } else {
        // If header is somehow larger than document, truncate
        copy(result, []byte(header[:len(result)]))
    }
    
    return result
}

// calculateMedian finds the median upload time
func calculateMedian(times []time.Duration) time.Duration {
    if len(times) == 0 {
        return 0
    }
    
    // Sort the times
    sortDurations(times)
    
    // Find median
    middle := len(times) / 2
    if len(times)%2 == 0 {
        return (times[middle-1] + times[middle]) / 2
    }
    return times[middle]
}

// calculatePercentile calculates the Nth percentile
func calculatePercentile(times []time.Duration, percentile int) time.Duration {
    if len(times) == 0 {
        return 0
    }
    
    // Sort the times
    sortDurations(times)
    
    // Calculate index
    index := int(float64(len(times)) * float64(percentile) / 100)
    if index >= len(times) {
        index = len(times) - 1
    }
    
    return times[index]
}

// sortDurations sorts a slice of time.Duration
func sortDurations(times []time.Duration) {
    // Simple bubble sort for simplicity
    for i := 0; i < len(times); i++ {
        for j := i + 1; j < len(times); j++ {
            if times[i] > times[j] {
                times[i], times[j] = times[j], times[i]
            }
        }
    }
}