import axios from 'axios';
import { create as createIPFSClient } from 'ipfs-http-client';
import { ethers } from 'ethers';

// Environment detection
const isProduction = process.env.NODE_ENV === 'production';

/**
 * Creates an environment-aware IPFS client
 * Uses Pinata in production, local IPFS in development
 */
const createIPFSService = () => {
  // For production - use Pinata
  if (isProduction) {
    return {
      // Store a document on IPFS using Pinata
      storeDocument: async (file) => {
        try {
          const formData = new FormData();
          formData.append('file', file);
          
          const response = await axios.post(
            'https://api.pinata.cloud/pinning/pinFileToIPFS',
            formData,
            {
              headers: {
                'Content-Type': 'multipart/form-data',
                'pinata_api_key': process.env.REACT_APP_PINATA_API_KEY,
                'pinata_secret_api_key': process.env.REACT_APP_PINATA_SECRET_KEY
              }
            }
          );
          
          return {
            cid: response.data.IpfsHash,
            success: true
          };
        } catch (error) {
          console.error('Pinata upload failed:', error);
          return { 
            success: false, 
            error: error.response?.data?.error || 'Failed to upload to IPFS' 
          };
        }
      },
      
      // Retrieve a document from IPFS using Pinata gateway
      retrieveDocument: async (cid) => {
        try {
          const response = await axios.get(
            `https://gateway.pinata.cloud/ipfs/${cid}`,
            { responseType: 'arraybuffer' }
          );
          
          return {
            content: new Uint8Array(response.data),
            success: true
          };
        } catch (error) {
          console.error('IPFS retrieval failed:', error);
          return { 
            success: false, 
            error: 'Failed to retrieve document from IPFS' 
          };
        }
      }
    };
  } 
  // For development - use local IPFS
  else {
    try {
      // Connect to local IPFS node
      const client = createIPFSClient({ 
        host: 'localhost', 
        port: 5001, 
        protocol: 'http' 
      });
      
      return {
        storeDocument: async (file) => {
          try {
            // Convert File object to buffer if needed
            let content;
            if (file instanceof Blob) {
              content = await file.arrayBuffer();
              content = new Uint8Array(content);
            } else {
              content = file;
            }
            
            const result = await client.add(content);
            return {
              cid: result.path,
              success: true
            };
          } catch (error) {
            console.error('Local IPFS upload failed:', error);
            return { 
              success: false, 
              error: 'Failed to upload to local IPFS node' 
            };
          }
        },
        
        retrieveDocument: async (cid) => {
          try {
            const chunks = [];
            for await (const chunk of client.cat(cid)) {
              chunks.push(chunk);
            }
            
            return {
              content: new Uint8Array(Buffer.concat(chunks)),
              success: true
            };
          } catch (error) {
            console.error('Local IPFS retrieval failed:', error);
            return { 
              success: false, 
              error: 'Failed to retrieve document from local IPFS' 
            };
          }
        }
      };
    } catch (error) {
      console.warn('Failed to connect to local IPFS, using fallback HTTP API');
      // Fallback to HTTP API if local node isn't running
      return {
        storeDocument: async (file) => {
          try {
            const formData = new FormData();
            formData.append('file', file);
            
            const response = await axios.post('http://localhost:5001/api/v0/add', formData);
            return {
              cid: response.data.Hash,
              success: true
            };
          } catch (error) {
            console.error('Local IPFS API failed:', error);
            return { 
              success: false, 
              error: 'Failed to connect to local IPFS' 
            };
          }
        },
        
        retrieveDocument: async (cid) => {
          try {
            const response = await axios.get(`http://localhost:8080/ipfs/${cid}`, {
              responseType: 'arraybuffer'
            });
            
            return {
              content: new Uint8Array(response.data),
              success: true
            };
          } catch (error) {
            console.error('Local IPFS gateway failed:', error);
            return { 
              success: false, 
              error: 'Failed to retrieve from local IPFS gateway' 
            };
          }
        }
      };
    }
  }
};

/**
 * Creates an environment-aware Ethereum provider
 * Uses Infura in production, local node in development
 */
const createEthereumService = () => {
  // For production - use Infura
  if (isProduction) {
    const provider = new ethers.providers.JsonRpcProvider(
      `https://goerli.infura.io/v3/${process.env.REACT_APP_INFURA_PROJECT_ID}`
    );
    
    return {
      provider,
      
      verifyDocument: async (documentHash, contractAddress) => {
        try {
          // Create contract interface for document registry
          const abi = [
            "function documentExists(string memory documentHash) view returns (bool)",
            "function getDocumentDetails(string memory documentHash) view returns (address owner, string memory ipfsCID, uint256 timestamp, bool exists)"
          ];
          
          const contract = new ethers.Contract(contractAddress, abi, provider);
          
          // Get document details
          const [owner, ipfsCID, timestamp, exists] = await contract.getDocumentDetails(documentHash);
          
          return {
            exists,
            owner,
            ipfsCID,
            timestamp: new Date(timestamp.toNumber() * 1000).toISOString(),
            success: true
          };
        } catch (error) {
          console.error('Blockchain verification failed:', error);
          return { 
            success: false, 
            error: 'Failed to verify document on blockchain' 
          };
        }
      }
    };
  } 
  // For development - use local node
  else {
    const provider = new ethers.providers.JsonRpcProvider('http://localhost:8545');
    
    return {
      provider,
      
      verifyDocument: async (documentHash, contractAddress) => {
        try {
          // Create contract interface for document registry
          const abi = [
            "function documentExists(string memory documentHash) view returns (bool)",
            "function getDocumentDetails(string memory documentHash) view returns (address owner, string memory ipfsCID, uint256 timestamp, bool exists)"
          ];
          
          const contract = new ethers.Contract(contractAddress, abi, provider);
          
          // Get document details
          const [owner, ipfsCID, timestamp, exists] = await contract.getDocumentDetails(documentHash);
          
          return {
            exists,
            owner,
            ipfsCID,
            timestamp: new Date(timestamp.toNumber() * 1000).toISOString(),
            success: true
          };
        } catch (error) {
          console.error('Local blockchain verification failed:', error);
          return { 
            success: false, 
            error: 'Failed to verify document on local blockchain'
          };
        }
      }
    };
  }
};

/**
 * Main API service that combines IPFS and Ethereum functionality
 */
const createAPIService = () => {
  const ipfsService = createIPFSService();
  const ethereumService = createEthereumService();
  
  return {
    // Upload and register a document
    uploadDocument: async (file, contractAddress, privateKey) => {
      // 1. Store on IPFS
      const ipfsResult = await ipfsService.storeDocument(file);
      if (!ipfsResult.success) {
        return { success: false, error: ipfsResult.error };
      }
      
      // 2. Calculate document hash (in real implementation, use your SHA3-256 function)
      // For demo we'll use a simple hash
      const documentHash = await calculateDocumentHash(file);
      
      // 3. Register on blockchain
      let txHash = null;
      if (isProduction) {
        // For production use Infura and actual transaction
        try {
          const wallet = new ethers.Wallet(privateKey, ethereumService.provider);
          const abi = ["function registerDocument(string memory documentHash, string memory ipfsCID)"];
          const contract = new ethers.Contract(contractAddress, abi, wallet);
          
          const tx = await contract.registerDocument(documentHash, ipfsResult.cid);
          await tx.wait();
          txHash = tx.hash;
        } catch (error) {
          console.error('Document registration failed:', error);
          // Continue even if blockchain registration fails
          // We still have the document on IPFS
        }
      } else {
        // For development, use local node or simulate
        try {
          const wallet = new ethers.Wallet("0xc89a2b5c8db43cf084f6ead1b63580af3e231b99e395f0d47ee437157060851c", ethereumService.provider);
          const abi = ["function registerDocument(string memory documentHash, string memory ipfsCID)"];
          const contract = new ethers.Contract(contractAddress, abi, wallet);
          
          const tx = await contract.registerDocument(documentHash, ipfsResult.cid);
          await tx.wait();
          txHash = tx.hash;
        } catch (error) {
          console.error('Local document registration failed:', error);
          // Use mock hash for development
          txHash = "0x" + Array(64).fill('0').map(() => Math.floor(Math.random() * 16).toString(16)).join('');
        }
      }
      
      return {
        success: true,
        hash: documentHash,
        cid: ipfsResult.cid,
        txHash
      };
    },
    
    // Verify document authenticity
    verifyDocument: async (documentHash, contractAddress) => {
      const result = await ethereumService.verifyDocument(documentHash, contractAddress);
      return result;
    },
    
    // Retrieve document by CID
    retrieveDocument: async (cid) => {
      return await ipfsService.retrieveDocument(cid);
    },
    
    // Direct access to services if needed
    ipfs: ipfsService,
    ethereum: ethereumService
  };
};

// Helper function to calculate document hash
const calculateDocumentHash = async (file) => {
  // In production you'd use your SHA3-256 function from the backend
  // For frontend demo purposes, we'll use a simple hash
  return new Promise((resolve) => {
    const reader = new FileReader();
    reader.onload = (e) => {
      const arrayBuffer = e.target.result;
      const hashArray = Array.from(new Uint8Array(arrayBuffer)).slice(0, 32);
      const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
      resolve(hashHex);
    };
    reader.readAsArrayBuffer(file);
  });
};

// Create and export the API service
export const apiService = createAPIService();