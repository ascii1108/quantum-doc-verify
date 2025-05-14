package ipfs

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io"
    "io/ioutil"
    "os"

    shell "github.com/ipfs/go-ipfs-api"
)

// IPFSClient handles interactions with IPFS
type IPFSClient struct {
    shell *shell.Shell
}

// NewIPFSClient creates a new IPFS client
func NewIPFSClient(apiURL string) *IPFSClient {
    return &IPFSClient{
        shell: shell.NewShell(apiURL),
    }
}

// EncryptDocument encrypts a document using AES-256-GCM
func EncryptDocument(document []byte, password []byte) ([]byte, error) {
    // Generate a strong encryption key from the password
    key := sha256.Sum256(password)
    
    // Create a new AES cipher block
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }
    
    // Create a GCM cipher mode (authenticated encryption)
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }
    
    // Generate a random nonce
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, fmt.Errorf("failed to generate nonce: %w", err)
    }
    
    // Encrypt and seal the data
    ciphertext := gcm.Seal(nonce, nonce, document, nil)
    
    return ciphertext, nil
}

// DecryptDocument decrypts a document using AES-256-GCM
func DecryptDocument(ciphertext []byte, password []byte) ([]byte, error) {
    // Generate key from password
    key := sha256.Sum256(password)
    
    // Create a new AES cipher block
    block, err := aes.NewCipher(key[:])
    if err != nil {
        return nil, fmt.Errorf("failed to create cipher: %w", err)
    }
    
    // Create a GCM cipher mode
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, fmt.Errorf("failed to create GCM: %w", err)
    }
    
    // Extract the nonce from the ciphertext
    nonceSize := gcm.NonceSize()
    if len(ciphertext) < nonceSize {
        return nil, fmt.Errorf("ciphertext too short")
    }
    
    nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
    
    // Decrypt the data
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return nil, fmt.Errorf("failed to decrypt: %w", err)
    }
    
    return plaintext, nil
}

// StoreDocument stores an encrypted document on IPFS
func (ic *IPFSClient) StoreDocument(documentPath string, password []byte) (string, error) {
    // Read the document
    document, err := os.ReadFile(documentPath)
    if err != nil {
        return "", fmt.Errorf("failed to read document: %w", err)
    }
    
    // Encrypt the document
    encryptedDoc, err := EncryptDocument(document, password)
    if err != nil {
        return "", fmt.Errorf("failed to encrypt document: %w", err)
    }
    
    // Store on IPFS
    cid, err := ic.shell.Add(bytes.NewReader(encryptedDoc))
    if err != nil {
        return "", fmt.Errorf("failed to store on IPFS: %w", err)
    }
    
    return cid, nil
}

// RetrieveDocument retrieves and decrypts a document from IPFS
func (ic *IPFSClient) RetrieveDocument(cid string, password []byte, outputPath string) error {
    // Get the encrypted document from IPFS
    reader, err := ic.shell.Cat(cid)
    if err != nil {
        return fmt.Errorf("failed to retrieve from IPFS: %w", err)
    }
    defer reader.Close()
    
    // Read the encrypted data
    encryptedDoc, err := ioutil.ReadAll(reader)
    if err != nil {
        return fmt.Errorf("failed to read from IPFS: %w", err)
    }
    
    // Decrypt the document
    document, err := DecryptDocument(encryptedDoc, password)
    if err != nil {
        return fmt.Errorf("failed to decrypt document: %w", err)
    }
    
    // Write to output file
    err = os.WriteFile(outputPath, document, 0644)
    if err != nil {
        return fmt.Errorf("failed to write output file: %w", err)
    }
    
    return nil
}

// GetDocumentHash returns the hash of an IPFS CID
// This can be used for blockchain registration
func (ic *IPFSClient) GetDocumentHash(cid string) (string, error) {
    // For simplicity, we'll use the CID as a base for the hash
    hash := sha256.Sum256([]byte(cid))
    return hex.EncodeToString(hash[:]), nil
}

// PinDocument pins a document to ensure it remains on IPFS
func (ic *IPFSClient) PinDocument(cid string) error {
    err := ic.shell.Pin(cid)
    if err != nil {
        return fmt.Errorf("failed to pin document: %w", err)
    }
    return nil
}

// Add uploads raw bytes to IPFS and returns the CID
func (ic *IPFSClient) Add(data []byte) (string, error) {
    // Upload to IPFS
    cid, err := ic.shell.Add(bytes.NewReader(data))
    if err != nil {
        return "", fmt.Errorf("failed to store on IPFS: %w", err)
    }
    
    return cid, nil
}

// Cat retrieves raw bytes from IPFS by CID
func (ic *IPFSClient) Cat(cid string) ([]byte, error) {
    // Get the data from IPFS
    reader, err := ic.shell.Cat(cid)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve from IPFS: %w", err)
    }
    defer reader.Close()
    
    // Read all data
    data, err := io.ReadAll(reader)
    if err != nil {
        return nil, fmt.Errorf("failed to read from IPFS: %w", err)
    }
    
    return data, nil
}

// Client defines an interface for IPFS operations
type Client interface {
    // UploadFile uploads a file to IPFS and returns its CID
    UploadFile(filePath string) (cid string, err error)
    
    // DownloadFile downloads a file from IPFS by its CID
    DownloadFile(cid string, outputPath string) error
}

// NewClient creates a new IPFS client
func NewClient(nodeAddr string) (Client, error) {
    // Implement using your existing IPFS code
    return &ipfsClient{
        nodeAddr: nodeAddr,
    }, nil
}

// Implementation of ipfsClient
type ipfsClient struct {
    nodeAddr string
}

func (c *ipfsClient) UploadFile(filePath string) (string, error) {
    // Call your existing IPFS code to upload a file
    // You can execute a command using exec.Command to run your IPFS binary
    return "QmSampleCID", nil // Example implementation
}

func (c *ipfsClient) DownloadFile(cid string, outputPath string) error {
    // Call your existing IPFS code to download a file
    return nil // Example implementation
}