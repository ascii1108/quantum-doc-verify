package storage

import (
    "bytes"
    "crypto/rand"
    "testing"
)

func TestHybridEncryption(t *testing.T) {
    // Create some test data
    content := []byte("This is a secret document that needs quantum-resistant encryption")
    
    // Generate a dummy key for testing (in real use, this would be a valid Dilithium key)
    dummyKey := make([]byte, 32)
    rand.Read(dummyKey)
    
    // Encrypt the data
    encryptedData, err := EncryptDocument(content, dummyKey)
    if err != nil {
        t.Fatalf("Encryption failed: %v", err)
    }
    
    // Ensure encryption actually happened
    if bytes.Equal(content, encryptedData) {
        t.Fatalf("Encryption did not change the content")
    }
    
    // Decrypt the data
    decryptedData, err := DecryptDocument(encryptedData, dummyKey)
    if err != nil {
        t.Fatalf("Decryption failed: %v", err)
    }
    
    // Verify the decrypted data matches the original
    if !bytes.Equal(content, decryptedData) {
        t.Fatalf("Decryption did not restore original content")
    }
}