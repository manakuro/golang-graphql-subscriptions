import React from 'react';
import { ApolloProvider } from '@apollo/client';
import { client } from './lib/apolloClient';
import { ChakraProvider } from "@chakra-ui/react"
import { Component } from './Component';

function App() {
  return (
    <ApolloProvider client={client}>
      <ChakraProvider resetCSS>
        <Component />
      </ChakraProvider>
    </ApolloProvider>
  );
}

export default App;
