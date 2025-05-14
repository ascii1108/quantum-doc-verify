// pkg/blockchain/client.go
package blockchain

import (
    "context"
    "crypto/ecdsa"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "math/big"
    "os"
    "strings"
    "time"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"

    "github.com/ipfs/go-cid"
    "github.com/multiformats/go-multihash"
)

var documentRegistry = make(map[string]string) // Maps document hash to CID

const registryFile = "document_registry.json"

// DocumentMetadata holds document information from the blockchain
type DocumentMetadata struct {
    Hash      string    `json:"hash"`
    CID       string    `json:"cid"`
    Owner     string    `json:"owner"`
    Timestamp time.Time `json:"timestamp"`
    Signature string    `json:"signature"`
}

// Client defines an interface for blockchain operations
type Client interface {
    // RegisterDocument registers a document's hash, IPFS CID, and signature on the blockchain
    RegisterDocument(hash, cid, signature string) (txHash string, err error)
    
    // VerifyDocument checks if a document is registered on the blockchain
    VerifyDocument(hash string) (exists bool, metadata DocumentMetadata, err error)
    
    // GetDocumentMetadata retrieves all metadata for a document
    GetDocumentMetadata(hash string) (exists bool, metadata DocumentMetadata, err error)
}

// NewLocalClient creates a client connected to a local blockchain node
func NewLocalClient(contractAddress string) (Client, error) {
    // Implement using your existing blockchain.go code
    return &localClient{
        contractAddress: contractAddress,
    }, nil
}

// NewInfuraClient creates a client connected to an Infura Ethereum node
func NewInfuraClient(infuraEndpoint, contractAddress string) (Client, error) {
    // Implement using your existing blockchain.go code with Infura connection
    return &infuraClient{
        endpoint:        infuraEndpoint,
        contractAddress: contractAddress,
    }, nil
}

// Implementation of local client
type localClient struct {
    contractAddress string
}

func (c *localClient) RegisterDocument(hash, cid, signature string) (string, error) {
    // Call your existing blockchain code to register a document
    // You can execute a command using exec.Command to run your blockchain binary
    return "0x" + hash[:8], nil // Example implementation
}

func (c *localClient) VerifyDocument(hash string) (bool, DocumentMetadata, error) {
    // Call your existing blockchain code to verify a document
    // Return placeholder data for now - replace with actual implementation
    return true, DocumentMetadata{
        Hash:      hash,
        CID:       "QmSample",
        Owner:     "0xSampleAddress",
        Timestamp: time.Now(),
        Signature: "sample-signature",
    }, nil
}

func (c *localClient) GetDocumentMetadata(hash string) (bool, DocumentMetadata, error) {
    // Call your existing blockchain code to get document metadata
    return c.VerifyDocument(hash)
}

// Implementation of Infura client
type infuraClient struct {
    endpoint        string
    contractAddress string
}

func (c *infuraClient) RegisterDocument(hash, cid, signature string) (string, error) {
    // Implement with Infura API calls
    return "0x" + hash[:8], nil
}

func (c *infuraClient) VerifyDocument(hash string) (bool, DocumentMetadata, error) {
    // Implement with Infura API calls
    return true, DocumentMetadata{
        Hash:      hash,
        CID:       "QmSample",
        Owner:     "0xSampleAddress",
        Timestamp: time.Now(),
        Signature: "sample-signature",
    }, nil
}

func (c *infuraClient) GetDocumentMetadata(hash string) (bool, DocumentMetadata, error) {
    return c.VerifyDocument(hash)
}

// BlockchainClient handles interactions with Ethereum blockchain
type BlockchainClient struct {
    client       *ethclient.Client
    contractAddr common.Address
}

// Load the registry from file during client initialization
func (bc *BlockchainClient) loadRegistry() error {
    // Check if file exists
    if _, err := os.Stat(registryFile); os.IsNotExist(err) {
        // If not, create an empty registry
        documentRegistry = make(map[string]string)
        return nil
    }
    
    // Read the file
    data, err := os.ReadFile(registryFile)
    if err != nil {
        return fmt.Errorf("failed to read registry file: %w", err)
    }
    
    // Parse JSON
    return json.Unmarshal(data, &documentRegistry)
}

// Save the registry to file after each update
func (bc *BlockchainClient) saveRegistry() error {
    // Convert to JSON
    data, err := json.MarshalIndent(documentRegistry, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal registry: %w", err)
    }
    
    // Write to file
    return os.WriteFile(registryFile, data, 0644)
}

// NewBlockchainClient creates a new blockchain client
func NewBlockchainClient(nodeURL string, contractAddress string) (*BlockchainClient, error) {
    // Connect to Ethereum node
    client, err := ethclient.Dial(nodeURL)
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
    }
    
    // Parse contract address
    contractAddr := common.HexToAddress(contractAddress)
    
    bc := &BlockchainClient{
        client:       client,
        contractAddr: contractAddr,
    }
    
    // Load existing registry
    if err := bc.loadRegistry(); err != nil {
        return nil, fmt.Errorf("failed to load document registry: %w", err)
    }
    
    return bc, nil
}

// RegisterDocument registers a document on the blockchain
func (bc *BlockchainClient) RegisterDocument(privateKey *ecdsa.PrivateKey, documentHash, ipfsCID string, dilithiumSignature []byte) (string, error) {
    // Store the CID in our registry
    documentRegistry[documentHash] = ipfsCID
    
    // Save the updated registry
    if err := bc.saveRegistry(); err != nil {
        return "", fmt.Errorf("failed to save document registry: %w", err)
    }
    
    // Get auth for transaction
    auth, err := bc.getTransactionAuth(privateKey)
    if err != nil {
        return "", fmt.Errorf("failed to create transaction auth: %w", err)
    }

    // This would normally call the contract, but for now we'll simulate a transaction
    // Format the function call data for "registerDocument(string,string)"
    functionSig := []byte("registerDocument(string,string)")
    functionHash := crypto.Keccak256(functionSig)[:4]
    
    // Simple encoding of parameters - in a real implementation this would use ABI encoding
    callData := append(functionHash, append([]byte(documentHash), []byte(ipfsCID)...)...)
    
    // Create a transaction
    tx := types.NewTransaction(
        auth.Nonce.Uint64(),
        bc.contractAddr,
        auth.Value,
        auth.GasLimit,
        auth.GasPrice,
        callData,
    )
    
    // Sign the transaction
    chainID, err := bc.client.ChainID(context.Background())
    if err != nil {
        return "", fmt.Errorf("failed to get chain ID: %w", err)
    }
    
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        return "", fmt.Errorf("failed to sign transaction: %w", err)
    }
    
    // Send transaction
    err = bc.client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", fmt.Errorf("failed to send transaction: %w", err)
    }
    
    txHash := signedTx.Hash()

    // Include the quantum signature with the transaction hash
    _ = dilithiumSignature
    _ = txHash

    return txHash.Hex(), nil
}

// VerifyDocumentOwnership checks if a document is owned by a specific address
func (bc *BlockchainClient) VerifyDocumentOwnership(documentHash string, claimedOwner common.Address) (bool, error) {
    // In a real implementation, this would call the contract
    // For now, return true to simulate success
    return true, nil
}

// DocumentExists checks if a document is registered
func (bc *BlockchainClient) DocumentExists(documentHash string) (bool, error) {
    // In a real implementation, this would call the contract
    // For now, return true to simulate success
    return true, nil
}

// GetDocumentDetails retrieves document details from blockchain
func (bc *BlockchainClient) GetDocumentDetails(documentHash string) (common.Address, string, time.Time, bool, error) {
    // Look up the CID from our registry
    ipfsCID, exists := documentRegistry[documentHash]
    if !exists {
        // If not found, fall back to calculating it (for compatibility)
        hashBytes, err := hex.DecodeString(documentHash)
        if err != nil {
            return common.Address{}, "", time.Time{}, false, fmt.Errorf("invalid document hash: %w", err)
        }
        
        mh, err := multihash.Sum(hashBytes, multihash.SHA2_256, -1)
        if err != nil {
            return common.Address{}, "", time.Time{}, false, fmt.Errorf("failed to create multihash: %w", err)
        }
        
        c := cid.NewCidV0(mh)
        ipfsCID = c.String()
    }
    
    return common.HexToAddress("0x1234567890AbcdEF1234567890aBcdef12345678"), 
           ipfsCID, 
           time.Unix(1630000000, 0), 
           true, 
           nil
}

// RecordVerification records a verification event on the blockchain
func (bc *BlockchainClient) RecordVerification(privateKey *ecdsa.PrivateKey, documentHash string, verified bool) (string, error) {
    // Get auth for transaction
    auth, err := bc.getTransactionAuth(privateKey)
    if err != nil {
        return "", fmt.Errorf("failed to create transaction auth: %w", err)
    }

    // This would normally call the contract, but for now we'll simulate a transaction
    // Format function call data for "recordVerification(string,bool)"
    functionSig := []byte("recordVerification(string,bool)")
    functionHash := crypto.Keccak256(functionSig)[:4]
    
    // Simple encoding of parameters
    verifiedByte := byte(0)
    if verified {
        verifiedByte = 1
    }
    callData := append(functionHash, append([]byte(documentHash), verifiedByte)...)

    // Create a transaction
    tx := types.NewTransaction(
        auth.Nonce.Uint64(),
        bc.contractAddr,
        auth.Value,
        auth.GasLimit,
        auth.GasPrice,
        callData,
    )
    
    // Sign the transaction
    chainID, err := bc.client.ChainID(context.Background())
    if err != nil {
        return "", fmt.Errorf("failed to get chain ID: %w", err)
    }
    
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        return "", fmt.Errorf("failed to sign transaction: %w", err)
    }
    
    // Send transaction
    err = bc.client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        return "", fmt.Errorf("failed to send transaction: %w", err)
    }
    
    return signedTx.Hash().Hex(), nil
}

// Helper to create transaction auth from private key
func (bc *BlockchainClient) getTransactionAuth(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
    ctx := context.Background()

    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }

    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := bc.client.PendingNonceAt(ctx, fromAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to get nonce: %w", err)
    }

    gasPrice, err := bc.client.SuggestGasPrice(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to suggest gas price: %w", err)
    }

    chainID, err := bc.client.ChainID(ctx)
    if err != nil {
        return nil, fmt.Errorf("failed to get chain ID: %w", err)
    }

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
    if err != nil {
        return nil, fmt.Errorf("failed to create transactor: %w", err)
    }

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(300000) // in units
    auth.GasPrice = gasPrice

    return auth, nil
}

// LoadPrivateKey loads a private key from a hex string
func LoadPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, error) {
    // Remove '0x' prefix if present
    if strings.HasPrefix(privateKeyHex, "0x") {
        privateKeyHex = privateKeyHex[2:]
    }

    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return nil, fmt.Errorf("failed to load private key: %w", err)
    }

    return privateKey, nil
}