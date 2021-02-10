import { gql } from '@apollo/client';

export default gql`
    query {
      messages {
        id
        message
      }
    }
`
