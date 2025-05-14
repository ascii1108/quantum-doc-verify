// web/src/components/DocumentVerify.js
import React, { useState } from 'react';
import { 
  Box, Button, FormControl, FormLabel, Input, 
  VStack, Heading, Text, useToast, Alert, AlertIcon
} from '@chakra-ui/react';
import { apiService } from '../services/api';

function DocumentVerify() {
  const [documentHash, setDocumentHash] = useState('');
  const [isVerifying, setIsVerifying] = useState(false);
  const [verificationResult, setVerificationResult] = useState(null);
  const toast = useToast();
  
  // Use contract address from environment variable or fall back to default
  const contractAddress = process.env.REACT_APP_CONTRACT_ADDRESS || "0xBc59F6A37b6283889Bd25405b822909ab03d0f6B";

  const handleVerify = async (e) => {
    e.preventDefault();
    
    if (!documentHash) {
      toast({
        title: "Document hash required",
        status: "warning",
        duration: 3000,
        isClosable: true,
      });
      return;
    }

    setIsVerifying(true);
    
    try {
      // Use our API service for verification
      const result = await apiService.verifyDocument(documentHash, contractAddress);
      
      if (!result.success) {
        throw new Error(result.error);
      }
      
      setVerificationResult(result);
      
      if (!result.exists) {
        toast({
          title: "Document not found",
          description: "This document is not registered on the blockchain",
          status: "warning",
          duration: 5000,
          isClosable: true,
        });
      }
    } catch (error) {
      console.error('Verification failed:', error);
      
      toast({
        title: "Verification failed",
        description: error.message || "Something went wrong",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
      
      setVerificationResult(null);
    } finally {
      setIsVerifying(false);
    }
  };

  return (
    <Box maxW="600px" mx="auto" p={5}>
      <VStack spacing={5} align="stretch">
        <Heading as="h1" size="lg">Verify Document Authenticity</Heading>
        
        <form onSubmit={handleVerify}>
          <VStack spacing={4}>
            <FormControl isRequired>
              <FormLabel>Document Hash</FormLabel>
              <Input 
                value={documentHash}
                onChange={(e) => setDocumentHash(e.target.value)}
                placeholder="Enter document hash"
                variant="filled"
              />
            </FormControl>
            
            <Button 
              type="submit" 
              colorScheme="teal" 
              isLoading={isVerifying}
              loadingText="Verifying..."
              isDisabled={!documentHash}
              width="full"
            >
              Verify Document
            </Button>
          </VStack>
        </form>
        
        {verificationResult && (
          <Box mt={5}>
            {verificationResult.exists ? (
              <Alert status="success">
                <AlertIcon />
                <Box>
                  <Heading as="h3" size="sm">Document Verified Successfully</Heading>
                  <Text mt={2}><strong>Owner:</strong> {verificationResult.owner}</Text>
                  <Text><strong>Registration Date:</strong> {new Date(verificationResult.timestamp).toLocaleString()}</Text>
                  <Text><strong>IPFS CID:</strong> {verificationResult.ipfsCID}</Text>
                </Box>
              </Alert>
            ) : (
              <Alert status="error">
                <AlertIcon />
                <Box>
                  <Heading as="h3" size="sm">Document Verification Failed</Heading>
                  <Text mt={2}>This document is not registered on the blockchain.</Text>
                </Box>
              </Alert>
            )}
          </Box>
        )}
      </VStack>
    </Box>
  );
}

export default DocumentVerify;