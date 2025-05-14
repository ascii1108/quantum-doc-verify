// web/src/App.js
import React from 'react';
import { ChakraProvider, Box, Flex, Spacer, Link, Container } from '@chakra-ui/react';
import { BrowserRouter as Router, Route, Routes, NavLink } from 'react-router-dom';
import DocumentUpload from './components/DocumentUpload';
import DocumentVerify from './components/DocumentVerify';

function App() {
  return (
    <ChakraProvider>
      <Router>
        <Box>
          <Flex 
            as="nav" 
            align="center" 
            wrap="wrap" 
            padding="1.5rem" 
            bg="teal.500" 
            color="white"
          >
            <Box>
              <Link 
                as={NavLink} 
                to="/" 
                fontSize="xl" 
                fontWeight="bold"
                _hover={{ textDecoration: 'none' }}
              >
                Quantum-Doc-Verify
              </Link>
            </Box>
            <Spacer />
            <Box display="flex" gap={4}>
              <Link 
                as={NavLink} 
                to="/" 
                p={2} 
                borderRadius="md"
                _activeLink={{ bg: 'teal.700' }}
                _hover={{ bg: 'teal.600' }}
              >
                Upload
              </Link>
              <Link 
                as={NavLink} 
                to="/verify" 
                p={2} 
                borderRadius="md"
                _activeLink={{ bg: 'teal.700' }}
                _hover={{ bg: 'teal.600' }}
              >
                Verify
              </Link>
            </Box>
          </Flex>

          <Container maxW="container.lg" py={8}>
            <Routes>
              <Route path="/" element={<DocumentUpload />} />
              <Route path="/verify" element={<DocumentVerify />} />
            </Routes>
          </Container>
        </Box>
      </Router>
    </ChakraProvider>
  );
}

export default App;