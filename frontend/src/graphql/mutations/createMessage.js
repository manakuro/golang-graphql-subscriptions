import { gql } from '@apollo/client';

export default gql`
  mutation($message: String!) {
    createMessage(message: $message) {
      message
    }
  }
`
