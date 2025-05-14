# Quantum-Doc-Verify

A blockchain-based quantum-resistant document verification framework that combines post-quantum cryptography with Ethereum smart contracts and IPFS decentralized storage.

## Overview

Quantum-Doc-Verify provides a complete solution for securing document verification systems against threats from quantum computing. Using Dilithium, a quantum-resistant lattice-based digital signature scheme, this framework enables users to sign, store, verify, and manage documents with cryptographic guarantees that remain secure even against quantum attacks.

## Features

- **Quantum-Resistant Signatures**: Implementation of Dilithium (NIST PQC finalist) for robust post-quantum security
- **Blockchain Integration**: Ethereum smart contracts for immutable verification records
- **Decentralized Storage**: IPFS integration for distributed document storage
- **Complete Workflow**: End-to-end process from document signing to verification
- **Performance Optimized**: Benchmarked for high throughput and scalability
- **Command-Line Interface**: Simple tools for all document management functions

## Architecture

![Architecture Diagram](https://example.com/architecture-diagram.png)

The system consists of four main components:

1. **Cryptographic Layer**: Handles Dilithium key generation, document signing, and signature verification
2. **Storage Layer**: Manages document storage and retrieval through IPFS
3. **Blockchain Layer**: Interacts with Ethereum smart contracts for tamper-proof record keeping
4. **Integration Layer**: Provides CLI tools and APIs that bring all components together

## Installation

### Prerequisites

- Go 1.18+
- Node.js 14+ (for Ethereum interaction)
- IPFS daemon (for local development)
- Ethereum client or access to an Ethereum network

### Building from Source

```bash
# Clone the repository
git clone https://github.com/yourusername/quantum-doc-verify.git
cd quantum-doc-verify

# Build all components
./build_server.sh
```

## Usage

### Key Generation

```bash
./bin/quantum-doc-verify generate-keys
```

### Document Signing and Registration

```bash
./bin/quantum-doc-verify store-register --file=document.pdf --contract=0x12345... --eth-key=your_private_key
```

### Document Verification

```bash
./bin/quantum-doc-verify verify-retrieve --hash=document_hash --cid=ipfs_cid --contract=0x12345... --out=retrieved_document.pdf
```

### Full Demo

```bash
./demo.sh
```

## Benchmarks

The system has been benchmarked for performance under various conditions:

- **Throughput**: 50+ documents/second with 20 concurrent clients
- **Scalability**: Linear scaling up to 50 concurrent clients
- **Document Size Impact**: Minimal impact on performance with documents up to 1MB
- **Signature Performance**: Dilithium signing takes approximately 3-5ms per document

## Smart Contract

The Ethereum smart contract provides the following functions:

- Document registration with owner, hash, and IPFS CID
- Ownership verification
- Existence checking
- Retrieval of document metadata

## Security Considerations

- Private keys should be securely stored and never shared
- For production use, ensure proper access controls to the document registration interface
- Consider additional encryption for highly sensitive documents
- The system provides authentication and integrity but not confidentiality by default

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Citation

If you use this system in your research, please cite:

```
J. Mathew, "Quantum-Doc-Verify: A Blockchain-based Quantum-Resistant Document Verification Framework," 
in IEEE International Conference on Quantum Computing Applications, 2025.
```

## Acknowledgements

- NIST Post-Quantum Cryptography Standardization Project
- Ethereum Foundation
- IPFS Project
- Go Ethereum (geth) team