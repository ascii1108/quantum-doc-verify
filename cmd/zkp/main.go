package main

import (
    
    "os"
    "path/filepath"
    "encoding/json"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"

    "quantum-doc-verify/pkg/zkp"
)

func main() {
    // Configure logging
    log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

    var rootCmd = &cobra.Command{
        Use:   "zkp",
        Short: "Zero-Knowledge Proof tool for quantum-resistant document verification",
        Long:  "A command-line tool to manage Zero-Knowledge Proofs for quantum-resistant document verification",
    }

    // Add subcommands
    rootCmd.AddCommand(setupCmd())
    rootCmd.AddCommand(proveCmd())
    rootCmd.AddCommand(verifyCmd())

    if err := rootCmd.Execute(); err != nil {
        log.Fatal().Err(err).Msg("Failed to execute command")
    }
}

func setupCmd() *cobra.Command {
    var keysDir string
    var proofsDir string

    cmd := &cobra.Command{
        Use:   "setup",
        Short: "Set up the ZK proof system",
        Run: func(cmd *cobra.Command, args []string) {
            setupZKP(keysDir, proofsDir)
        },
    }

    cmd.Flags().StringVar(&keysDir, "keys-dir", "./zkp_keys", "Directory to store ZKP keys")
    cmd.Flags().StringVar(&proofsDir, "proofs-dir", "./zkp_proofs", "Directory to store ZKP proofs")

    return cmd
}

func proveCmd() *cobra.Command {
    var docPath string
    var privKeyPath string
    var pubKeyPath string
    var sigPath string
    var proofName string

    cmd := &cobra.Command{
        Use:   "prove",
        Short: "Generate a ZK proof for a document",
        Run: func(cmd *cobra.Command, args []string) {
            generateProof(docPath, privKeyPath, pubKeyPath, sigPath, proofName)
        },
    }

    cmd.Flags().StringVar(&docPath, "doc", "", "Path to the document")
    cmd.Flags().StringVar(&privKeyPath, "privkey", "", "Path to the private key")
    cmd.Flags().StringVar(&pubKeyPath, "pubkey", "", "Path to the public key")
    cmd.Flags().StringVar(&sigPath, "sig", "", "Path to the signature")
    cmd.Flags().StringVar(&proofName, "name", "proof", "Name for the proof file")

    cmd.MarkFlagRequired("doc")
    cmd.MarkFlagRequired("privkey")
    cmd.MarkFlagRequired("pubkey")
    cmd.MarkFlagRequired("sig")

    return cmd
}

func verifyCmd() *cobra.Command {
    var proofPath string
    var pubKeyPath string
    var sigPath string

    cmd := &cobra.Command{
        Use:   "verify",
        Short: "Verify a ZK proof",
        Run: func(cmd *cobra.Command, args []string) {
            verifyProof(proofPath, pubKeyPath, sigPath)
        },
    }

    cmd.Flags().StringVar(&proofPath, "proof", "", "Path to the proof")
    cmd.Flags().StringVar(&pubKeyPath, "pubkey", "", "Path to the public key")
    cmd.Flags().StringVar(&sigPath, "sig", "", "Path to the signature")

    cmd.MarkFlagRequired("proof")
    cmd.MarkFlagRequired("pubkey")
    cmd.MarkFlagRequired("sig")

    return cmd
}

func setupZKP(keysDir, proofsDir string) {
    log.Info().Msg("Setting up ZK proof system...")

    // Create a ZKP manager that uses the simplified prover
    zkpManager, err := zkp.NewZKPManager(keysDir, proofsDir)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create ZKP manager")
    }

    // Initialize keys
    err = zkpManager.InitializeKeys()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to initialize ZKP keys")
    }

    log.Info().Str("keys-dir", keysDir).Msg("ZK proving and verifying keys generated")
}

func generateProof(docPath, privKeyPath, pubKeyPath, sigPath, proofName string) {
    log.Info().Str("document", docPath).Msg("Generating ZK proof...")

    // Create a simplified prover directly instead of going through the deprecated NewDocumentProver
    prover, err := zkp.NewSimpleDocumentProver()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create prover")
    }

    // Load document and hash it
    doc, err := zkp.LoadDocument(docPath)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load document")
    }

    // Calculate document hash
    docHash := zkp.HashDocument(doc)

    // Generate proof
    proofDir := "./zkp_proofs"
    proofPath := filepath.Join(proofDir, proofName+".zkp")

    // Ensure proof directory exists
    if err := os.MkdirAll(proofDir, 0755); err != nil {
        log.Fatal().Err(err).Msg("Failed to create proofs directory")
    }

    // Generate and save the proof
    err = zkp.GenerateAndSaveProof(
        prover,
        docPath,
        privKeyPath,
        docHash,
        pubKeyPath,
        sigPath,
        proofPath,
    )
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to generate proof")
    }

    log.Info().Str("proof", proofPath).Msg("ZK proof generated and saved")
}

func verifyProof(proofPath, pubKeyPath, sigPath string) {
    log.Info().Str("proof", proofPath).Msg("Verifying ZK proof...")

    // Create a simplified prover directly
    prover, err := zkp.NewSimpleDocumentProver()
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to create prover")
    }

    // Load and parse the proof to extract the document hash
    proofData, err := os.ReadFile(proofPath)
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to read proof file")
    }
    
    // Parse the proof to extract document hash
    var simpleProof zkp.SimpleProof
    if err := json.Unmarshal(proofData, &simpleProof); err != nil {
        log.Fatal().Err(err).Msg("Failed to parse proof data")
    }
    
    // Use the document hash from the proof
    docHash := simpleProof.DocumentHash

    // Verify the proof
    valid, err := zkp.LoadAndVerifyProof(
        prover,
        proofPath,
        docHash,
        pubKeyPath,
        sigPath,
    )
    if err != nil {
        log.Fatal().Err(err).Msg("Error verifying proof")
    }

    if valid {
        log.Info().Msg("✅ Proof is valid! Document verified.")
    } else {
        log.Warn().Msg("❌ Proof is invalid! Document verification failed.")
    }
}