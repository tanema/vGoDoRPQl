import React from 'react';
import ReactDOM from 'react-dom';
import { ApolloProvider } from 'react-apollo';
import ApolloClient from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { InMemoryCache } from 'apollo-cache-inmemory';

import App from './components/App';

let apiURL = process.env.NODE_ENV === 'production' ?
  'http://localhost:5000/graphql' :
  'http://localhost:5000/graphql';

const apolloClient = new ApolloClient({
  link: new HttpLink({uri: apiURL}),
  cache: new InMemoryCache(),
});

ReactDOM.render(
  <ApolloProvider client={apolloClient}>
    <App />
  </ApolloProvider>,
  document.getElementById('root'),
);
