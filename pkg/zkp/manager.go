package zkp

import (
    "fmt"
    "os"
    "path/filepath"
)

// ZKPManager manages zero-knowledge proofs for documents
type ZKPManager struct {
    prover       DocumentProverInterface // Changed to use the interface
    keysDir      string
    proofsDir    string
    initialized  bool
}

// NewZKPManager creates a new ZK proof manager
func NewZKPManager(keysDir string, proofsDir string) (*ZKPManager, error) {
    // Create the necessary directories
    if err := ensureDir(keysDir); err != nil {
        return nil, fmt.Errorf("failed to create keys directory: %w", err)
    }
    
    if err := ensureDir(proofsDir); err != nil {
        return nil, fmt.Errorf("failed to create proofs directory: %w", err)
    }
    
    // Use SimpleDocumentProver instead of DocumentProver
    prover, err := NewSimpleDocumentProver()
    if err != nil {
        return nil, fmt.Errorf("failed to create document prover: %w", err)
    }
    
    return &ZKPManager{
        prover:      prover,
        keysDir:     keysDir,
        proofsDir:   proofsDir,
        initialized: true,
    }, nil
}

// ensureDir makes sure a directory exists
func ensureDir(dir string) error {
    return os.MkdirAll(dir, 0755)
}

// InitializeKeys initializes and saves the proving and verifying keys
func (zkm *ZKPManager) InitializeKeys() error {
    if !zkm.initialized {
        return fmt.Errorf("ZKP manager not initialized")
    }
    
    // Save the proving key
    provingKeyPath := filepath.Join(zkm.keysDir, "proving_key.bin")
    err := zkm.prover.SaveProvingKey(provingKeyPath)
    if err != nil {
        return fmt.Errorf("failed to save proving key: %w", err)
    }
    
    // Save the verifying key
    verifyingKeyPath := filepath.Join(zkm.keysDir, "verifying_key.bin")
    err = zkm.prover.SaveVerifyingKey(verifyingKeyPath)
    if err != nil {
        return fmt.Errorf("failed to save verifying key: %w", err)
    }
    
    return nil
}

// GenerateProof generates a ZK proof for a document
func (zkm *ZKPManager) GenerateProof(
    documentPath string,
    privateKeyPath string,
    documentHashBytes []byte,
    publicKeyPath string,
    signaturePath string,
    proofName string,
) (string, error) {
    if !zkm.initialized {
        return "", fmt.Errorf("ZKP manager not initialized")
    }
    
    // Generate proof filename
    proofPath := filepath.Join(zkm.proofsDir, proofName+".zkp")
    
    // Generate and save the proof
    err := GenerateAndSaveProof(
        zkm.prover,
        documentPath,
        privateKeyPath,
        documentHashBytes,
        publicKeyPath,
        signaturePath,
        proofPath,
    )
    if err != nil {
        return "", fmt.Errorf("failed to generate proof: %w", err)
    }
    
    return proofPath, nil
}

// VerifyProof verifies a ZK proof
func (zkm *ZKPManager) VerifyProof(
    proofPath string,
    documentHashBytes []byte,
    publicKeyPath string,
    signaturePath string,
) (bool, error) {
    if !zkm.initialized {
        return false, fmt.Errorf("ZKP manager not initialized")
    }
    
    // Verify the proof
    return LoadAndVerifyProof(
        zkm.prover,
        proofPath,
        documentHashBytes,
        publicKeyPath,
        signaturePath,
    )
}

// GetProver returns the underlying DocumentProver
func (zkm *ZKPManager) GetProver() DocumentProverInterface {
    return zkm.prover
}