// web/src/components/DocumentUpload.js
import React, { useState } from 'react';
import { 
  Box, Button, FormControl, FormLabel, Input, 
  VStack, Heading, Text, useToast, Progress 
} from '@chakra-ui/react';
import { apiService } from '../services/api';

function DocumentUpload() {
  const [file, setFile] = useState(null);
  const [isUploading, setIsUploading] = useState(false);
  const [uploadResult, setUploadResult] = useState(null);
  const toast = useToast();
  
  // Use contract address from environment variable or fall back to default
  const contractAddress = process.env.REACT_APP_CONTRACT_ADDRESS || "0xBc59F6A37b6283889Bd25405b822909ab03d0f6B";

  const handleFileChange = (e) => {
    if (e.target.files.length > 0) {
      setFile(e.target.files[0]);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    
    if (!file) {
      toast({
        title: "No file selected",
        status: "warning",
        duration: 3000,
        isClosable: true,
      });
      return;
    }

    setIsUploading(true);
    
    try {
      // Use our API service for uploading
      const result = await apiService.uploadDocument(
        file, 
        contractAddress, 
        null // For demo, we're not using a private key in the frontend
      );
      
      if (!result.success) {
        throw new Error(result.error);
      }
      
      setUploadResult(result);
      
      toast({
        title: "Document uploaded successfully",
        status: "success",
        duration: 5000,
        isClosable: true,
      });
    } catch (error) {
      console.error('Upload failed:', error);
      
      toast({
        title: "Upload failed",
        description: error.message || "Something went wrong",
        status: "error",
        duration: 5000,
        isClosable: true,
      });
    } finally {
      setIsUploading(false);
    }
  };

  return (
    <Box maxW="600px" mx="auto" p={5}>
      <VStack spacing={5} align="stretch">
        <Heading as="h1" size="lg">Quantum-Resistant Document Upload</Heading>
        
        <form onSubmit={handleSubmit}>
          <VStack spacing={4}>
            <FormControl isRequired>
              <FormLabel>Select Document</FormLabel>
              <Input 
                type="file" 
                onChange={handleFileChange}
                variant="filled"
                p={1}
              />
            </FormControl>
            
            <Button 
              type="submit" 
              colorScheme="blue" 
              isLoading={isUploading}
              loadingText="Uploading and signing..."
              isDisabled={!file}
              width="full"
            >
              Upload & Register Document
            </Button>
          </VStack>
        </form>
        
        {isUploading && (
          <Box mt={4}>
            <Text mb={2}>Processing document with quantum-resistant signatures...</Text>
            <Progress size="sm" isIndeterminate colorScheme="blue" />
          </Box>
        )}
        
        {uploadResult && (
          <Box 
            mt={5} 
            p={4} 
            borderWidth={1} 
            borderRadius="md" 
            bgColor="blue.50"
          >
            <Heading as="h3" size="md" mb={3}>
              Document Registered Successfully
            </Heading>
            
            <VStack align="start" spacing={2}>
              <Text><strong>Document Hash:</strong> {uploadResult.hash}</Text>
              <Text><strong>IPFS CID:</strong> {uploadResult.cid}</Text>
              {uploadResult.txHash && (
                <Text><strong>Blockchain Transaction:</strong> {uploadResult.txHash}</Text>
              )}
            </VStack>
          </Box>
        )}
      </VStack>
    </Box>
  );
}

export default DocumentUpload;