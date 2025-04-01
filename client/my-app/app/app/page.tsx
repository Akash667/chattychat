"use client"

import { useState, useRef, useContext, useEffect } from 'react';
import ChatBody from './ChatBody';
import { WebSocketContext } from '@/modules/websocket_provider';
import { useRouter } from 'next/navigation';
import { API_URL } from '@/constants';
import autosize from 'autosize'
import { AuthContext } from '@/modules/auth_provider';


export type Message = {
  content: string
  client_id: string
  username: string
  room_id: string
  type: 'recv' | 'self'
}

const Page = () => {

  const [messages, setMessage] = useState<Array<Message>>([])
  const textarea = useRef<HTMLTextAreaElement>(null)
  const { conn } = useContext(WebSocketContext)
  const [users, setUsers] = useState<Array<{ username: string }>>([])
  const router = useRouter()

  const { user } = useContext(AuthContext)


  useEffect(() => {
    if (conn === null) {
      router.push('/')
      return
    }

    const roomId = conn.url.split('/')[5]

    async function getUsers() {
      try {
        console.log("fetching users")
        const res = await fetch(`${API_URL}/ws/getClients/${roomId}`,
          {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' }
          }
        )
        const data = await res.json()

        console.log('data fetched in getUsers: ' + JSON.stringify(data))
        setUsers(data)
      } catch (e) {
        console.log("error fetching users" + e)
      }
    }
    getUsers()

  }, [])    // get clients in the room


  useEffect(() => {
    console.log('useEffect triggered as state changed')
    console.log('current list of messages is as follows' + messages)
    if (textarea.current) {
      autosize(textarea.current)
      console.log('autosize triggered')
    }


    if (conn == null) {
      router.push('/')
      return
    }


    conn.onmessage = (message) => {
      console.log("useEffect triggered as message received")
      const m: Message = JSON.parse(message.data)
      if (m.content == "User has joined the room") {
        setUsers([...users, { username: m.username }])
      }


      if (m.content == "User has left the chat") {
        setUsers(users.filter((user) => user.username != m.username))
        setMessage([...messages, m])
        return
      }

      if (user?.username == m.username) {
        m.type = 'self';
      } else {
        m.type = 'recv';
      }
      setMessage([...messages, m])

    }


    conn.onclose = () => { }

    conn.onopen = () => { }
    conn.onerror = () => { }

  }, [textarea, messages, conn, users])     // handle websocket connection


  const sendMessage = () => {
    if (!textarea.current?.value) return
    // check connections
    if (conn == null) {
      router.push("/")
      return
    }
    console.log('sending message')
    conn.send(textarea.current.value)
    textarea.current.value = ''
  }
  return (
    <div className='flex flex-col w-full'>


      <div className='absolute top-4 right-4'>
        <button
          className='p-2 rounded-md bg-red-500 text-blue hover:bg-red-600'
          onClick={() => {
            // Clear user session and redirect to login page
            localStorage.removeItem('user_info')
            router.push('/');
          }}
        >
          Logout
        </button>
      </div>


      <div className='p-4 md:mx-6 mb-14' >
        <ChatBody data={messages} />
      </div>
      <div className='fixed bottom-0 mt-4 w-full'>
        <div className='flex md:flex-row px-4 py-2 bg-grey md:mx-4 rounded-md'>
          <div className='flex w-full mr-4 rounded-md border border-blue'>
            <textarea
              ref={textarea}
              placeholder='type your message here'
              className='w-full h-10 p-2 rounded-md focus:outline-none'
              style={{ resize: 'none' }}
            />
          </div>
          <div className='flex items-center'>
            <button className='p-2 rounded-md bg-blue text-white' onClick={sendMessage}>
              Send
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Page;
