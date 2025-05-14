package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
    
    "quantum-doc-verify/pkg/storage"
)

func main() {
    // Configure logging
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    var rootCmd = &cobra.Command{
        Use:   "ipfs",
        Short: "IPFS tool for quantum-resistant document storage",
        Long:  "A command-line tool to interact with IPFS for document storage",
    }

    // Add subcommands
    rootCmd.AddCommand(storeCmd())
    rootCmd.AddCommand(retrieveCmd())

    if err := rootCmd.Execute(); err != nil {
        log.Fatal().Err(err).Msg("Failed to execute command")
    }
}

func storeCmd() *cobra.Command {
    var filePath string
    var encrypt bool
    var publicKeyPath string
    var ipfsGateway string

    cmd := &cobra.Command{
        Use:   "store",
        Short: "Store a document on IPFS",
        Run: func(cmd *cobra.Command, args []string) {
            storeDocument(filePath, encrypt, publicKeyPath, ipfsGateway)
        },
    }

    cmd.Flags().StringVar(&filePath, "file", "", "Path to document file")
    cmd.Flags().BoolVar(&encrypt, "encrypt", false, "Encrypt document before storing")
    cmd.Flags().StringVar(&publicKeyPath, "pubkey", "", "Path to recipient's public key (required for encryption)")
    cmd.Flags().StringVar(&ipfsGateway, "gateway", "localhost:5001", "IPFS gateway address")
    cmd.MarkFlagRequired("file")

    return cmd
}

func retrieveCmd() *cobra.Command {
    var cid string
    var outputPath string
    var decrypt bool
    var privateKeyPath string
    var ipfsGateway string

    cmd := &cobra.Command{
        Use:   "retrieve",
        Short: "Retrieve a document from IPFS",
        Run: func(cmd *cobra.Command, args []string) {
            retrieveDocument(cid, outputPath, decrypt, privateKeyPath, ipfsGateway)
        },
    }

    cmd.Flags().StringVar(&cid, "cid", "", "IPFS CID of the document")
    cmd.Flags().StringVar(&outputPath, "out", "", "Output path for retrieved document")
    cmd.Flags().BoolVar(&decrypt, "decrypt", false, "Decrypt document after retrieval")
    cmd.Flags().StringVar(&privateKeyPath, "privkey", "", "Path to recipient's private key (required for decryption)")
    cmd.Flags().StringVar(&ipfsGateway, "gateway", "localhost:5001", "IPFS gateway address")
    cmd.MarkFlagRequired("cid")
    cmd.MarkFlagRequired("out")

    return cmd
}

func storeDocument(filePath string, encrypt bool, publicKeyPath string, ipfsGateway string) {
    log.Info().
        Str("file", filePath).
        Bool("encrypt", encrypt).
        Msg("Storing document on IPFS...")

    // Create IPFS client
    ipfs, err := storage.NewIPFSClient(ipfsGateway)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create IPFS client")
    }

    // Read file content
    content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to read file")
    }

    // Handle encryption if requested
    if encrypt {
        if publicKeyPath == "" {
            log.Fatal().Msg("Public key path is required for encryption")
        }
        
        // Read public key
        pubKey, err := os.ReadFile(publicKeyPath)
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to read public key file")
        }
        
        // Encrypt content (you'll need to implement this)
        content, err = storage.EncryptDocument(content, pubKey)
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to encrypt document")
        }
    }

    // Store on IPFS
    cid, err := ipfs.Store(content)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to store document on IPFS")
    }

    log.Info().
        Str("cid", cid).
        Msg("Document stored on IPFS successfully!")
    
    // Calculate document hash for blockchain registration
    hash := storage.CalculateDocumentHash(content)
    
    fmt.Println("\nDocument Storage:")
    fmt.Printf("IPFS CID: %s\n", cid)
    fmt.Printf("Document Hash: %s\n", hash)
    fmt.Println("\nYou can register this document on the blockchain with:")
    fmt.Printf("./bin/blockchain register --contract=YOUR_CONTRACT_ADDRESS --key=YOUR_PRIVATE_KEY --hash=%s --cid=%s\n", hash, cid)
}

func retrieveDocument(cid string, outputPath string, decrypt bool, privateKeyPath string, ipfsGateway string) {
    log.Info().
        Str("cid", cid).
        Bool("decrypt", decrypt).
        Msg("Retrieving document from IPFS...")

    // Create IPFS client
    ipfs, err := storage.NewIPFSClient(ipfsGateway)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create IPFS client")
    }

    // Retrieve from IPFS
    content, err := ipfs.Retrieve(cid)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to retrieve document from IPFS")
    }

    // Handle decryption if requested
    if decrypt {
        if privateKeyPath == "" {
            log.Fatal().Msg("Private key path is required for decryption")
        }
        
        // Read private key
        privKey, err := os.ReadFile(privateKeyPath)
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to read private key file")
        }
        
        // Decrypt content (you'll need to implement this)
        content, err = storage.DecryptDocument(content, privKey)
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to decrypt document")
        }
    }

    // Ensure output directory exists
    err = os.MkdirAll(filepath.Dir(outputPath), 0755)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create output directory")
    }

    // Write to output file
    err = os.WriteFile(outputPath, content, 0644)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to write output file")
    }

    // Calculate document hash for verification
    hash := storage.CalculateDocumentHash(content)

    log.Info().
        Str("output", outputPath).
        Msg("Document retrieved from IPFS successfully!")
    
    fmt.Println("\nDocument Retrieval:")
    fmt.Printf("IPFS CID: %s\n", cid)
    fmt.Printf("Document Hash: %s\n", hash)
    fmt.Printf("Output File: %s\n", outputPath)
}