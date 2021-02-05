import React, { useEffect, useState, useCallback } from 'react';
import { useSubscription, useMutation } from '@apollo/client';
import MESSAGE_CREATED from './graphql/subscriptions/messageCreated';
import CREATE_MESSAGE from './graphql/mutations/createMessage';
import { Input, Text, Stack, Flex, Button } from "@chakra-ui/react"

type Message = {
  message: string
}

export const Component: React.VFC = () => {
  const { data, loading } = useSubscription<Message>(MESSAGE_CREATED);
  const [createMessage] = useMutation<Message>(CREATE_MESSAGE);
  const [messages, setMessages] = useState<string[]>(['First message']);
  const [inputValue, setInputValue] = useState('')

  useEffect(() => {
    if (data?.message) setMessages(m => [...m, data?.message])
  }, [data?.message])
  console.log(data, loading)

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
          {messages.map(m => <Text key={m}>{m}</Text>)}
        </Stack>
      </Stack>
    </Flex>
  );
}
