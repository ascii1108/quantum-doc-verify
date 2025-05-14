package crypto

import (
    
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"

    "github.com/cloudflare/circl/sign"
    "github.com/cloudflare/circl/sign/dilithium/mode2"
)

type DilithiumSigner struct {
    privateKey sign.PrivateKey
    publicKey  sign.PublicKey
    scheme     sign.Scheme
}

// NewDilithiumSigner creates a new DilithiumSigner instance
func NewDilithiumSigner() *DilithiumSigner {
    return &DilithiumSigner{
        scheme: mode2.Scheme(),
    }
}

// GenerateKeypair generates a new keypair and returns the public and private key bytes
func (ds *DilithiumSigner) GenerateKeypair() ([]byte, []byte, error) {
    pub, priv, err := ds.scheme.GenerateKey()
    if err != nil {
        return nil, nil, fmt.Errorf("failed to generate keypair: %w", err)
    }

    ds.publicKey = pub
    ds.privateKey = priv

    pubBytes, err := pub.MarshalBinary()
    if err != nil {
        return nil, nil, fmt.Errorf("failed to marshal public key: %w", err)
    }

    privBytes, err := priv.MarshalBinary()
    if err != nil {
        return nil, nil, fmt.Errorf("failed to marshal private key: %w", err)
    }

    return pubBytes, privBytes, nil
}

// SignDocument signs a document using a private key
func (ds *DilithiumSigner) SignDocument(docPath string, privateKeyBytes []byte) ([]byte, error) {
    // Load the private key if provided
    if privateKeyBytes != nil {
        err := ds.LoadPrivateKey(privateKeyBytes)
        if err != nil {
            return nil, err
        }
    }

    if ds.privateKey == nil {
        return nil, fmt.Errorf("private key not available")
    }

    content, err := os.ReadFile(docPath)
    if err != nil {
        return nil, fmt.Errorf("failed to read document: %w", err)
    }

    signature := ds.scheme.Sign(ds.privateKey, content, nil)
    return signature, nil
}

// VerifySignature verifies a document signature
func (ds *DilithiumSigner) VerifySignature(docPath string, signature, publicKeyBytes []byte) (bool, error) {
    content, err := os.ReadFile(docPath)
    if err != nil {
        return false, fmt.Errorf("failed to read document: %w", err)
    }

    // Load the public key from bytes
    publicKey, err := ds.scheme.UnmarshalBinaryPublicKey(publicKeyBytes)
    if err != nil {
        return false, fmt.Errorf("failed to unmarshal public key: %w", err)
    }

    valid := ds.scheme.Verify(publicKey, content, signature, nil)
    return valid, nil
}

// GetDocumentHash returns the SHA-256 hash of a document
func (ds *DilithiumSigner) GetDocumentHash(docPath string) (string, error) {
    content, err := os.ReadFile(docPath)
    if err != nil {
        return "", fmt.Errorf("failed to read document: %w", err)
    }

    hash := sha256.Sum256(content)
    return hex.EncodeToString(hash[:]), nil
}

// SaveKeys saves the keypair to files
func (ds *DilithiumSigner) SaveKeys(publicKeyBytes, privateKeyBytes []byte, pubKeyPath, privKeyPath string) error {
    if err := os.WriteFile(pubKeyPath, publicKeyBytes, 0644); err != nil {
        return fmt.Errorf("failed to save public key: %w", err)
    }
    
    if err := os.WriteFile(privKeyPath, privateKeyBytes, 0644); err != nil {
        return fmt.Errorf("failed to save private key: %w", err)
    }
    
    return nil
}

// ExportPublicKey returns the binary representation of the public key
func (ds *DilithiumSigner) ExportPublicKey() ([]byte, error) {
    if ds.publicKey == nil {
        return nil, fmt.Errorf("public key not available")
    }
    return ds.publicKey.MarshalBinary()
}

// ExportPrivateKey returns the binary representation of the private key
func (ds *DilithiumSigner) ExportPrivateKey() ([]byte, error) {
    if ds.privateKey == nil {
        return nil, fmt.Errorf("private key not available")
    }
    return ds.privateKey.MarshalBinary()
}

// LoadPrivateKey loads a private key from its binary representation
func (ds *DilithiumSigner) LoadPrivateKey(privateKeyBytes []byte) error {
    priv, err := ds.scheme.UnmarshalBinaryPrivateKey(privateKeyBytes)
    if err != nil {
        return fmt.Errorf("failed to unmarshal private key: %w", err)
    }
    ds.privateKey = priv
    return nil
}

// LoadPublicKey loads a public key from its binary representation
func (ds *DilithiumSigner) LoadPublicKey(publicKeyBytes []byte) error {
    pub, err := ds.scheme.UnmarshalBinaryPublicKey(publicKeyBytes)
    if err != nil {
        return fmt.Errorf("failed to unmarshal public key: %w", err)
    }
    ds.publicKey = pub
    return nil
}
