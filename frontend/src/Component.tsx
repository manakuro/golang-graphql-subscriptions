import React, { useEffect, useState, useCallback } from 'react';
import { useSubscription, useMutation, useQuery } from '@apollo/client';
import MESSAGE_CREATED from './graphql/subscriptions/messageCreated';
import CREATE_MESSAGE from './graphql/mutations/createMessage';
import MESSAGES from './graphql/query/messages';
import { Input, Text, Stack, Flex, Button } from "@chakra-ui/react"

type Message = {
  id: string
  message: string
}
type MessageSubscription = {
  messageCreated: Message
}
type MessagesQuery = {
  messages: Message[]
}

export const Component: React.VFC = () => {
  const { data } = useSubscription<MessageSubscription>(MESSAGE_CREATED);
  const [createMessage] = useMutation<Message>(CREATE_MESSAGE);
  const queryResult = useQuery<MessagesQuery>(MESSAGES)
  const [messages, setMessages] = useState<Message[]>([]);
  const [inputValue, setInputValue] = useState('')

  useEffect(() => {
    if (data?.messageCreated?.message) setMessages(m => [...m, data?.messageCreated])
  }, [data?.messageCreated])

  useEffect(() => {
    if (queryResult.data?.messages) setMessages(queryResult.data.messages)
  }, [queryResult.data?.messages])

  const handleChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setInputValue(e.target.value)
  }, [])

  const handleClick = useCallback(async (e) => {
    e.preventDefault()
    await createMessage({ variables: { message: inputValue } })
  }, [inputValue, createMessage])

  return (
    <Flex maxW="1180px" minH="100vh" m="0 auto" p={4} justifyContent="center">
      <Stack spacing={4} mt={12}>
        <Flex>
          <Input placeholder="enter message" w="400px" value={inputValue} onChange={handleChange} mr={3} />
          <Button onClick={handleClick}>Submit</Button>
        </Flex>

        <Stack spacing={2} p={4}>
          {messages.map((m, i) => <Text key={m.id}>{m.message}</Text>)}
        </Stack>
      </Stack>
    </Flex>
  );
}
