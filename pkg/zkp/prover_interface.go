package zkp

// DocumentProverInterface defines the interface for document provers
type DocumentProverInterface interface {
    GenerateProof(document, privateKey, documentHash, publicKey, signature []byte) ([]byte, error)
    VerifyProof(proof, documentHash, publicKey, signature []byte) (bool, error)
    SaveProvingKey(path string) error
    SaveVerifyingKey(path string) error
}