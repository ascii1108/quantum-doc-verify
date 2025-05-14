// contracts/DocumentRegistry.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract DocumentRegistry {
    struct Document {
        address owner;
        string documentHash;
        string ipfsCID;
        uint256 timestamp;
        bool exists;
    }
    
    // Document hash => Document details
    mapping(string => Document) public documents;
    
    // Owner address => List of document hashes
    mapping(address => string[]) public ownerDocuments;
    
    // Events
    event DocumentRegistered(string documentHash, string ipfsCID, address owner);
    event DocumentVerified(string documentHash, address verifier, bool verified);
    
    // Register a document
    function registerDocument(string memory documentHash, string memory ipfsCID) public {
        require(!documents[documentHash].exists, "Document already registered");
        
        documents[documentHash] = Document({
            owner: msg.sender,
            documentHash: documentHash,
            ipfsCID: ipfsCID,
            timestamp: block.timestamp,
            exists: true
        });
        
        ownerDocuments[msg.sender].push(documentHash);
        
        emit DocumentRegistered(documentHash, ipfsCID, msg.sender);
    }
    
    // Verify document ownership
    function verifyDocumentOwnership(string memory documentHash, address claimedOwner) public view returns (bool) {
        return documents[documentHash].exists && documents[documentHash].owner == claimedOwner;
    }
    
    // Check if document exists
    function documentExists(string memory documentHash) public view returns (bool) {
        return documents[documentHash].exists;
    }
    
    // Get document details
    function getDocumentDetails(string memory documentHash) public view returns (
        address owner,
        string memory ipfsCID,
        uint256 timestamp,
        bool exists
    ) {
        Document memory doc = documents[documentHash];
        return (doc.owner, doc.ipfsCID, doc.timestamp, doc.exists);
    }
    
    // Get count of documents owned by an address
    function getDocumentCount(address owner) public view returns (uint256) {
        return ownerDocuments[owner].length;
    }
    
    // Get document hash by index for an owner
    function getDocumentHashByIndex(address owner, uint256 index) public view returns (string memory) {
        require(index < ownerDocuments[owner].length, "Index out of bounds");
        return ownerDocuments[owner][index];
    }
    
    // Record verification
    function recordVerification(string memory documentHash, bool verified) public {
        require(documents[documentHash].exists, "Document does not exist");
        emit DocumentVerified(documentHash, msg.sender, verified);
    }
}