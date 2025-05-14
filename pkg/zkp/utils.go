package zkp

import (
    "crypto/sha256"
    "fmt"
    "os"
    "path/filepath"
)

// HashDocument creates a SHA-256 hash of a document
func HashDocument(document []byte) []byte {
    hash := sha256.Sum256(document)
    return hash[:]
}

// LoadDocument loads a document from a file and returns its content
func LoadDocument(path string) ([]byte, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return nil, fmt.Errorf("failed to read document: %w", err)
    }
    return content, nil
}

// SaveProofToFile saves a ZK proof to a file
func SaveProofToFile(proof []byte, path string) error {
    // Ensure the directory exists
    dir := filepath.Dir(path)
    if dir != "" && dir != "." {
        if err := os.MkdirAll(dir, 0755); err != nil {
            return fmt.Errorf("failed to create directory for proof: %w", err)
        }
    }
    return os.WriteFile(path, proof, 0644)
}

// LoadProofFromFile loads a ZK proof from a file
func LoadProofFromFile(path string) ([]byte, error) {
    return os.ReadFile(path)
}

// LoadProvingKey loads a proving key from a file
func LoadProvingKey(path string) ([]byte, error) {
    return os.ReadFile(path)
}

// LoadVerifyingKey loads a verifying key from a file
func LoadVerifyingKey(path string) ([]byte, error) {
    return os.ReadFile(path)
}

// GenerateAndSaveProof is a convenience function that generates a proof and saves it
func GenerateAndSaveProof(
    prover DocumentProverInterface, // Changed from DocumentProver to DocumentProverInterface
    documentPath string,
    privateKeyPath string,
    documentHashBytes []byte,
    publicKeyPath string,
    signaturePath string,
    proofPath string,
) error {
    // Load document
    document, err := LoadDocument(documentPath)
    if err != nil {
        return fmt.Errorf("failed to load document: %w", err)
    }
    
    // Load private key
    privateKey, err := LoadDocument(privateKeyPath)
    if err != nil {
        return fmt.Errorf("failed to load private key: %w", err)
    }
    
    // Load public key
    publicKey, err := LoadDocument(publicKeyPath)
    if err != nil {
        return fmt.Errorf("failed to load public key: %w", err)
    }
    
    // Load signature
    signature, err := LoadDocument(signaturePath)
    if err != nil {
        return fmt.Errorf("failed to load signature: %w", err)
    }
    
    // Generate proof
    proof, err := prover.GenerateProof(
        document,
        privateKey,
        documentHashBytes,
        publicKey,
        signature,
    )
    if err != nil {
        return fmt.Errorf("failed to generate proof: %w", err)
    }
    
    // Save proof
    return SaveProofToFile(proof, proofPath)
}

// LoadAndVerifyProof is a convenience function that loads a proof and verifies it
func LoadAndVerifyProof(
    prover DocumentProverInterface,
    proofPath string,
    documentHashBytes []byte,
    publicKeyPath string,
    signaturePath string,
) (bool, error) {
    // Load proof
    proof, err := LoadProofFromFile(proofPath)
    if err != nil {
        return false, fmt.Errorf("failed to load proof: %w", err)
    }
    
    // Load public key
    publicKey, err := LoadDocument(publicKeyPath)
    if err != nil {
        return false, fmt.Errorf("failed to load public key: %w", err)
    }
    
    // Load signature
    signature, err := LoadDocument(signaturePath)
    if err != nil {
        return false, fmt.Errorf("failed to load signature: %w", err)
    }
    
    // Verify proof
    return prover.VerifyProof(
        proof,
        documentHashBytes,
        publicKey,
        signature,
    )
}