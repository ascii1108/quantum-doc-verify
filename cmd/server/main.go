package main

import (
    "encoding/json"
    "flag"
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "golang.org/x/crypto/sha3"
    "encoding/hex"
    "crypto/sha256"

    "quantum-doc-verify/pkg/logger"
)

var (
    port           int
    ipfsNodeAddr   string
    contractAddr   string
    privateKeyPath string
    publicKeyPath  string
    uploadDir      string
    infuraEndpoint string
    loggerInstance *logger.Logger
    documentStore  = make(map[string]DocumentData)
)

// Document data structure
type DocumentData struct {
    Hash      string    `json:"hash"`
    CID       string    `json:"cid"`
    FileName  string    `json:"fileName"`
    MimeType  string    `json:"mimeType"`
    Size      int64     `json:"size"`
    Content   []byte    `json:"content,omitempty"`
    Signature string    `json:"signature"`
    TxHash    string    `json:"txHash"`
    Timestamp time.Time `json:"timestamp"`
}

func init() {
    flag.IntVar(&port, "port", 8080, "Port to run the server on")
    flag.StringVar(&ipfsNodeAddr, "ipfs", "localhost:5001", "IPFS node address")
    flag.StringVar(&contractAddr, "contract", "", "Document registry smart contract address")
    flag.StringVar(&privateKeyPath, "private-key", "./keys/dilithium_private.key", "Path to Dilithium private key")
    flag.StringVar(&publicKeyPath, "public-key", "./keys/dilithium_public.key", "Path to Dilithium public key")
    flag.StringVar(&uploadDir, "upload-dir", "./uploads", "Directory for temporary document uploads")
    flag.StringVar(&infuraEndpoint, "infura", "", "Infura endpoint for blockchain connection")
}

func main() {
    flag.Parse()

    // Initialize logger
    loggerInstance = logger.New("server")
    loggerInstance.Info("Starting Quantum-Doc-Verify API server")

    // Ensure upload directory exists
    if err := os.MkdirAll(uploadDir, 0755); err != nil {
        loggerInstance.Fatal("Failed to create upload directory", "error", err)
    }

    // Create router
    router := mux.NewRouter()

    // Health check endpoint
    router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "status":    "ok",
            "timestamp": time.Now().Format(time.RFC3339),
        })
    }).Methods("GET")

    // Document upload and sign endpoint
    router.HandleFunc("/api/documents", func(w http.ResponseWriter, r *http.Request) {
        handleDocumentUpload(w, r)
    }).Methods("POST")

    // Document verification endpoint
    router.HandleFunc("/api/documents/{hash}/verify", func(w http.ResponseWriter, r *http.Request) {
        handleDocumentVerify(w, r)
    }).Methods("GET")

    // Document retrieval endpoint
    router.HandleFunc("/api/documents/retrieve", func(w http.ResponseWriter, r *http.Request) {
        handleDocumentRetrieve(w, r)
    }).Methods("GET")

    // Create a CORS middleware
    corsMiddleware := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"}, // For development - restrict in production
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Accept", "Authorization"},
        AllowCredentials: true,
    })

    // Wrap the router with the CORS middleware
    handler := corsMiddleware.Handler(router)

    // Start server
    addr := fmt.Sprintf(":%d", port)
    loggerInstance.Info("Server listening", "port", port)
    if err := http.ListenAndServe(addr, handler); err != nil {
        loggerInstance.Fatal("Failed to start server", "error", err)
    }
}

func handleDocumentUpload(w http.ResponseWriter, r *http.Request) {
    // Parse multipart form (max 10MB)
    if err := r.ParseMultipartForm(10 << 20); err != nil {
        http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
        return
    }

    // Get the file from the request
    file, header, err := r.FormFile("document")
    if err != nil {
        http.Error(w, "Failed to get file: "+err.Error(), http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a temp file
    tempFilePath := filepath.Join(uploadDir, header.Filename)
    tempFile, err := os.Create(tempFilePath)
    if err != nil {
        http.Error(w, "Failed to create temp file: "+err.Error(), http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFilePath) // Clean up temp file after processing

    // Read the file content
    fileContent, err := io.ReadAll(file)
    if err != nil {
        http.Error(w, "Failed to read file: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Write to temp file
    tempFile.Write(fileContent)
    tempFile.Close() // Close so it can be reopened for reading

    // Get document hash - simulate using the binary
    documentHash := simulateHashDocument(tempFilePath)

    // Sign the document - simulate using the binary
    signature := simulateSignDocument(tempFilePath)

    // Upload to IPFS - simulate using the binary
    cid := simulateIPFSUpload(tempFilePath)

    // Register on blockchain - simulate using the binary
    txHash := simulateBlockchainRegistration(documentHash, cid, signature)

    // Store document data in memory
    documentStore[documentHash] = DocumentData{
        Hash:      documentHash,
        CID:       cid,
        FileName:  header.Filename,
        MimeType:  header.Header.Get("Content-Type"),
        Size:      header.Size,
        Content:   fileContent,
        Signature: signature,
        TxHash:    txHash,
        Timestamp: time.Now(),
    }

    // Also store with CID for retrieval by CID
    documentStore[cid] = documentStore[documentHash]

    loggerInstance.Info("Document stored",
        "hash", documentHash,
        "cid", cid,
        "filename", header.Filename,
        "size", header.Size)

    loggerInstance.Info("Document mappings",
        "hash_key", documentHash,
        "cid_key", cid,
        "hash_in_data", documentStore[documentHash].Hash,
        "cid_in_data", documentStore[documentHash].CID)

    // Prepare response
    response := map[string]string{
        "hash":      documentHash,
        "cid":       cid,
        "txHash":    txHash,
        "signature": signature,
    }

    // Send response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func handleDocumentVerify(w http.ResponseWriter, r *http.Request) {
    // Get document hash from URL
    vars := mux.Vars(r)
    hash := vars["hash"]

    // Simulate verification using the binary
    verified, owner, timestamp, cid := simulateDocumentVerification(hash)

    if !verified {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "verified": false,
            "message":  "Document not found on blockchain",
        })
        return
    }

    // Prepare response
    response := map[string]interface{}{
        "verified":  true,
        "owner":     owner,
        "timestamp": timestamp,
        "cid":       cid,
    }

    // Send response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func handleDocumentRetrieve(w http.ResponseWriter, r *http.Request) {
    // Get document hash and CID from request
    vars := mux.Vars(r)
    hash := vars["hash"]
    cid := r.URL.Query().Get("cid")

    if cid == "" {
        http.Error(w, "CID parameter is required", http.StatusBadRequest)
        return
    }

    // Look up document by CID
    docData, exists := documentStore[cid]
    if !exists {
        // If the document doesn't exist in our store, create a dummy document
        loggerInstance.Warn("Document not found in store, creating dummy document",
            "cid", cid, "hash", hash)

        // Create a dummy PDF file
        w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", hash+".pdf"))
        w.Header().Set("Content-Type", "application/pdf")

        dummyPDF := []byte("%PDF-1.4\n1 0 obj\n<</Type /Catalog /Pages 2 0 R>>\nendobj\n2 0 obj\n<</Type /Pages /Kids [3 0 R] /Count 1>>\nendobj\n3 0 obj\n<</Type /Page /Parent 2 0 R /Resources 4 0 R /MediaBox [0 0 612 792] /Contents 5 0 R>>\nendobj\n4 0 obj\n<</Font <</F1 6 0 R>>>>\nendobj\n5 0 obj\n<</Length 44>>\nstream\nBT /F1 24 Tf 100 700 Td (Quantum-Doc-Verify: " + hash + ") Tj ET\nendstream\nendobj\n6 0 obj\n<</Type /Font /Subtype /Type1 /BaseFont /Helvetica>>\nendobj\nxref\n0 7\n0000000000 65535 f\n0000000009 00000 n\n0000000056 00000 n\n0000000111 00000 n\n0000000212 00000 n\n0000000250 00000 n\n0000000344 00000 n\ntrailer\n<</Size 7 /Root 1 0 R>>\nstartxref\n406\n%%EOF")
        w.Write(dummyPDF)
        return
    }

    // Set response headers based on actual document data
    w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", docData.FileName))

    // Set appropriate content type or default to application/octet-stream
    contentType := docData.MimeType
    if contentType == "" {
        contentType = http.DetectContentType(docData.Content)
        if contentType == "" {
            contentType = "application/octet-stream"
        }
    }
    w.Header().Set("Content-Type", contentType)

    loggerInstance.Info("Serving document",
        "hash", hash,
        "cid", cid,
        "filename", docData.FileName,
        "size", len(docData.Content),
        "type", contentType)

    // Write the actual document content
    w.Write(docData.Content)
}

// Simulated functions that would normally call your binaries
func simulateHashDocument(filePath string) string {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return "0x" + hex.EncodeToString([]byte(filePath))
    }
    
    // Use SHA3-256 (Keccak) which is quantum-resistant
    hash := sha3.Sum256(content)
    return "0x" + hex.EncodeToString(hash[:])
}

func simulateSignDocument(filePath string) string {
    return "dilithium-" + fmt.Sprintf("%x", time.Now().UnixNano())
}

func simulateIPFSUpload(filePath string) string {
    content, err := os.ReadFile(filePath)
    if err != nil {
        return "Qm" + hex.EncodeToString([]byte(filePath))[:32]
    }
    
    // Create a deterministic CID based on content
    hash := sha256.Sum256(content)
    return "Qm" + hex.EncodeToString(hash[:])[:32]
}

func simulateBlockchainRegistration(hash, cid, signature string) string {
    hashInput := hash + cid + signature + fmt.Sprintf("%d", time.Now().UnixNano())
    hexString := fmt.Sprintf("%x", hashInput)

    maxLength := len(hexString)
    if maxLength > 64 {
        maxLength = 64
    }

    return "0x" + hexString[:maxLength]
}

func simulateDocumentVerification(hash string) (bool, string, string, string) {
    if strings.HasPrefix(hash, "0x") {
        hashInput := hash + fmt.Sprintf("%d", time.Now().UnixNano())
        hexString := fmt.Sprintf("%x", hashInput)

        ownerLength := len(hexString)
        if ownerLength > 40 {
            ownerLength = 40
        }

        cidLength := len(hexString)
        if cidLength > 32 {
            cidLength = 32
        }

        return true,
            "0x" + hexString[:ownerLength],
            time.Now().Add(-24 * time.Hour).Format(time.RFC3339),
            "Qm" + hexString[:cidLength]
    }
    return false, "", "", ""
}

func simulateDocumentRetrieval(cid string) (string, error) {
    tempFile := filepath.Join(uploadDir, "retrieved-"+cid)
    if _, err := os.Create(tempFile); err != nil {
        return "", err
    }
    return tempFile, nil
}