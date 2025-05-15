package crypto

import (
    "crypto/sha256"
)

// DeriveEncryptionKey derives an AES-256 encryption key from a Dilithium private key
// This ensures that only the document owner can decrypt the document
func DeriveEncryptionKey(dilithiumKey []byte) []byte {
    // Use SHA-256 to derive a 32-byte key (suitable for AES-256)
    // In a production system, use a proper key derivation function like HKDF
    hasher := sha256.New()
    hasher.Write(dilithiumKey)
    return hasher.Sum(nil) // Returns a 32-byte key
}