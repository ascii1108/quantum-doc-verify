package main

import (
    "context"
    "crypto/ecdsa" // Add this import for ecdsa type
    "fmt"
    "math/big"
    // "strings" - Remove or comment this out since it's not used
    "os"

    // "github.com/ethereum/go-ethereum/accounts/abi" - Remove or comment this out since it's not used
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"

    "quantum-doc-verify/pkg/blockchain"
)

func main() {
    // Configure logging
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    var rootCmd = &cobra.Command{
        Use:   "blockchain",
        Short: "Blockchain tool for quantum-resistant document verification",
        Long:  "A command-line tool to interact with blockchain for document verification",
    }

    // Add subcommands
    rootCmd.AddCommand(registerCmd())
    rootCmd.AddCommand(verifyCmd())
    rootCmd.AddCommand(detailsCmd())
    rootCmd.AddCommand(recordVerificationCmd())
    rootCmd.AddCommand(deployCmd())

    if err := rootCmd.Execute(); err != nil {
        log.Fatal().Err(err).Msg("Failed to execute command")
    }
}

func registerCmd() *cobra.Command {
    var nodeURL string
    var contractAddress string
    var privateKeyHex string
    var documentHash string
    var ipfsCID string
    var publicKeyPath string

    cmd := &cobra.Command{
        Use:   "register",
        Short: "Register a document on the blockchain",
        Run: func(cmd *cobra.Command, args []string) {
            registerDocument(nodeURL, contractAddress, privateKeyHex, documentHash, ipfsCID, publicKeyPath)
        },
    }

    cmd.Flags().StringVar(&nodeURL, "node", "http://localhost:8545", "Ethereum node URL")
    cmd.Flags().StringVar(&contractAddress, "contract", "", "Document registry contract address")
    cmd.Flags().StringVar(&privateKeyHex, "key", "", "Private key in hex format (with or without 0x prefix)")
    cmd.Flags().StringVar(&documentHash, "hash", "", "Document hash to register")
    cmd.Flags().StringVar(&ipfsCID, "cid", "", "IPFS CID of the document")
    cmd.Flags().StringVar(&publicKeyPath, "pubkey", "", "Path to public key file (optional)")
    
    cmd.MarkFlagRequired("contract")
    cmd.MarkFlagRequired("key")
    cmd.MarkFlagRequired("hash")
    cmd.MarkFlagRequired("cid")

    return cmd
}

func verifyCmd() *cobra.Command {
    var nodeURL string
    var contractAddress string
    var documentHash string
    var claimedOwner string

    cmd := &cobra.Command{
        Use:   "verify",
        Short: "Verify document ownership",
        Run: func(cmd *cobra.Command, args []string) {
            verifyDocumentOwnership(nodeURL, contractAddress, documentHash, claimedOwner)
        },
    }

    cmd.Flags().StringVar(&nodeURL, "node", "http://localhost:8545", "Ethereum node URL")
    cmd.Flags().StringVar(&contractAddress, "contract", "", "Document registry contract address")
    cmd.Flags().StringVar(&documentHash, "hash", "", "Document hash to verify")
    cmd.Flags().StringVar(&claimedOwner, "owner", "", "Claimed owner address")
    cmd.MarkFlagRequired("contract")
    cmd.MarkFlagRequired("hash")
    cmd.MarkFlagRequired("owner")

    return cmd
}

func detailsCmd() *cobra.Command {
    var nodeURL string
    var contractAddress string
    var documentHash string

    cmd := &cobra.Command{
        Use:   "details",
        Short: "Get document details from blockchain",
        Run: func(cmd *cobra.Command, args []string) {
            getDocumentDetails(nodeURL, contractAddress, documentHash)
        },
    }

    cmd.Flags().StringVar(&nodeURL, "node", "http://localhost:8545", "Ethereum node URL")
    cmd.Flags().StringVar(&contractAddress, "contract", "", "Document registry contract address")
    cmd.Flags().StringVar(&documentHash, "hash", "", "Document hash to get details for")
    cmd.MarkFlagRequired("contract")
    cmd.MarkFlagRequired("hash")

    return cmd
}

func recordVerificationCmd() *cobra.Command {
    var nodeURL string
    var contractAddress string
    var privateKeyHex string
    var documentHash string
    var verified bool

    cmd := &cobra.Command{
        Use:   "record-verification",
        Short: "Record document verification on blockchain",
        Run: func(cmd *cobra.Command, args []string) {
            recordDocumentVerification(nodeURL, contractAddress, privateKeyHex, documentHash, verified)
        },
    }

    cmd.Flags().StringVar(&nodeURL, "node", "http://localhost:8545", "Ethereum node URL")
    cmd.Flags().StringVar(&contractAddress, "contract", "", "Document registry contract address")
    cmd.Flags().StringVar(&privateKeyHex, "key", "", "Private key in hex format (with or without 0x prefix)")
    cmd.Flags().StringVar(&documentHash, "hash", "", "Document hash to record verification for")
    cmd.Flags().BoolVar(&verified, "verified", true, "Verification result (true/false)")
    cmd.MarkFlagRequired("contract")
    cmd.MarkFlagRequired("key")
    cmd.MarkFlagRequired("hash")

    return cmd
}

func deployCmd() *cobra.Command {
    var nodeURL string
    var privateKeyHex string
    
    cmd := &cobra.Command{
        Use:   "deploy",
        Short: "Deploy the document registry contract",
        Run: func(cmd *cobra.Command, args []string) {
            deployContract(nodeURL, privateKeyHex)
        },
    }
    
    cmd.Flags().StringVar(&nodeURL, "node", "http://localhost:8545", "Ethereum node URL")
    cmd.Flags().StringVar(&privateKeyHex, "key", "", "Private key in hex format (with or without 0x prefix)")
    cmd.MarkFlagRequired("key")
    
    return cmd
}

func registerDocument(nodeURL, contractAddress, privateKeyHex, documentHash, ipfsCID, publicKeyPath string) {
    log.Info().
        Str("hash", documentHash).
        Str("cid", ipfsCID).
        Msg("Registering document on blockchain...")

    // Create blockchain client
    client, err := blockchain.NewBlockchainClient(nodeURL, contractAddress)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create blockchain client")
    }

    // Load private key
    privateKey, err := blockchain.LoadPrivateKey(privateKeyHex)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load private key")
    }

    /*
    var publicKey []byte
    if publicKeyPath != "" {
        publicKey, err = os.ReadFile(publicKeyPath)
        if err != nil {
            log.Fatal().Err(err).Msg("Failed to read public key file")
        }
    }
		*/

    // Register document
    var dilithiumSignature []byte // Either get this from somewhere or use an empty signature
    txHash, err := client.RegisterDocument(privateKey, documentHash, ipfsCID, dilithiumSignature)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to register document")
    }

    log.Info().
        Str("txHash", txHash).
        Msg("Document registered on blockchain successfully!")
    
    fmt.Println("\nDocument registration transaction:")
    fmt.Printf("Transaction Hash: %s\n", txHash)
}

func verifyDocumentOwnership(nodeURL, contractAddress, documentHash, claimedOwner string) {
    log.Info().
        Str("hash", documentHash).
        Str("claimedOwner", claimedOwner).
        Msg("Verifying document ownership...")

    // Create blockchain client
    client, err := blockchain.NewBlockchainClient(nodeURL, contractAddress)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create blockchain client")
    }

    // Format the address
    ownerAddr := common.HexToAddress(claimedOwner)

    // Verify ownership
    isOwner, err := client.VerifyDocumentOwnership(documentHash, ownerAddr)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to verify document ownership")
    }

    if isOwner {
        log.Info().Msg("✅ Verification successful! Address is the document owner.")
    } else {
        log.Warn().Msg("❌ Verification failed! Address is not the document owner.")
    }

    fmt.Println("\nDocument Ownership:")
    fmt.Printf("Document Hash: %s\n", documentHash)
    fmt.Printf("Claimed Owner: %s\n", claimedOwner)
    fmt.Printf("Verified: %v\n", isOwner)
}

func getDocumentDetails(nodeURL, contractAddress, documentHash string) {
    log.Info().
        Str("hash", documentHash).
        Msg("Getting document details from blockchain...")

    // Create blockchain client
    client, err := blockchain.NewBlockchainClient(nodeURL, contractAddress)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create blockchain client")
    }

    // Get document details
    owner, ipfsCID, timestamp, exists, err := client.GetDocumentDetails(documentHash)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to get document details")
    }

    if exists {
        log.Info().Msg("Document found on blockchain!")
        fmt.Println("\nDocument Details:")
        fmt.Printf("Owner: %s\n", owner.Hex())
        fmt.Printf("IPFS CID: %s\n", ipfsCID)
        fmt.Printf("Timestamp: %s\n", timestamp.String())
    } else {
        log.Warn().Msg("❌ Document not found on blockchain!")
    }
}

func recordDocumentVerification(nodeURL, contractAddress, privateKeyHex, documentHash string, verified bool) {
    log.Info().
        Str("hash", documentHash).
        Bool("verified", verified).
        Msg("Recording document verification...")

    // Create blockchain client
    client, err := blockchain.NewBlockchainClient(nodeURL, contractAddress)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create blockchain client")
    }

    // Load private key
    privateKey, err := blockchain.LoadPrivateKey(privateKeyHex)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load private key")
    }

    // Record verification
    txHash, err := client.RecordVerification(privateKey, documentHash, verified)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to record verification")
    }

    log.Info().
        Str("txHash", txHash).
        Msg("Document verification recorded on blockchain successfully!")
    
    fmt.Println("\nVerification recording transaction:")
    fmt.Printf("Transaction Hash: %s\n", txHash)
    fmt.Printf("Document Hash: %s\n", documentHash)
    fmt.Printf("Verified: %v\n", verified)
}

func deployContract(nodeURL, privateKeyHex string) {
    log.Info().Msg("Deploying document registry contract...")
    
    // Connect to Ethereum node
    client, err := ethclient.Dial(nodeURL)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to connect to Ethereum node")
    }
    
    // Load private key
    privateKey, err := blockchain.LoadPrivateKey(privateKeyHex)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load private key")
    }
    
    // Get auth for transaction
    auth, err := createTransactionAuth(client, privateKey)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create transaction auth")
    }
    
    // Deploy contract
    contractAddr, tx, err := blockchain.DeployDocumentRegistryContract(auth, client)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to deploy contract")
    }
    
    log.Info().
        Str("contract", contractAddr.Hex()).
        Str("txHash", tx.Hash().Hex()).
        Msg("Contract deployed successfully!")
    
    fmt.Println("\nContract Deployment:")
    fmt.Printf("Contract Address: %s\n", contractAddr.Hex())
    fmt.Printf("Transaction Hash: %s\n", tx.Hash().Hex())
    fmt.Println("\nStore this contract address for future use with --contract flag")
}

// Helper function to create transaction auth
func createTransactionAuth(client *ethclient.Client, privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
    ctx := context.Background()
    
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }
    
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(ctx, fromAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to get nonce: %w", err)
    }
    
    gasPrice, err := client.SuggestGasPrice(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to suggest gas price: %w", err)
    }
    
    chainID, err := client.ChainID(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get chain ID: %w", err)
    }
    
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        return nil, fmt.Errorf("failed to create transactor: %w", err)
    }
    
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)      // in wei
    auth.GasLimit = uint64(3000000) // contract deployment needs more gas
    auth.GasPrice = gasPrice
    
    return auth, nil
}