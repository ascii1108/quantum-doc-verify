package main

import (
	"fmt"
	"log"
	"os"

	"quantum-doc-verify/pkg/crypto"
)

func main() {
	// Create a new Dilithium signer
	signer := crypto.NewDilithiumSigner()

	// Generate a keypair
	fmt.Println("Generating quantum-resistant Dilithium keypair...")
	pubKey, privKey, err := signer.GenerateKeypair()
	if err != nil {
		log.Fatalf("Failed to generate keypair: %v", err)
	}

	// Save keys to files
	err = signer.SaveKeys(pubKey, privKey, "public_key.bin", "private_key.bin")
	if err != nil {
		log.Fatalf("Failed to save keys: %v", err)
	}
	fmt.Println("Keys saved to public_key.bin and private_key.bin")

	// Create a test document
	testDoc := "test_document.txt"
	err = os.WriteFile(testDoc, []byte("This is a test document to be signed with Dilithium."), 0644)
	if err != nil {
		log.Fatalf("Failed to create test document: %v", err)
	}

	// Sign the document
	fmt.Println("Signing document...")
	signature, err := signer.SignDocument(testDoc, privKey)
	if err != nil {
		log.Fatalf("Failed to sign document: %v", err)
	}

	// Save signature to a file
	err = os.WriteFile("signature.bin", signature, 0644)
	if err != nil {
		log.Fatalf("Failed to save signature: %v", err)
	}
	fmt.Println("Document signed. Signature saved to signature.bin")

	// Verify the signature
	fmt.Println("Verifying signature...")
	valid, err := signer.VerifySignature(testDoc, signature, pubKey)
	if err != nil {
		log.Fatalf("Error during verification: %v", err)
	}

	if valid {
		fmt.Println("Signature is valid! Document has not been tampered with.")
	} else {
		fmt.Println("Signature verification failed! Document may have been tampered with.")
	}

	// Get document hash (will be used later for blockchain storage)
	docHash, err := signer.GetDocumentHash(testDoc)
	if err != nil {
		log.Fatalf("Failed to get document hash: %v", err)
	}
	fmt.Printf("Document hash (to be stored on blockchain): %s\n", docHash)
}
