'use client'
import { stringify } from 'querystring'
import React, { useState } from 'react'

const index = () => {

  const [rooms, setRooms] = useState<{ id: string, name: string }[]>([
    { id: '1', name: 'room1' },
    { id: '2', name: 'room2' },
    { id: '3', name: 'room3' },
  ])

  const [roomName, setRoomName] = useState('')
  return (
    <div className='my-8 px-4 md:mx-32 w-full h-full'>
      <div className='flex justify-center mt-3 p-5'>
        <input type='text' className='border border-grey p-2 rounded-md focus:outline-none focus:border-blue' placeholder="room name"
        value={roomName}
        onChange={(e)=>{
          setRoomName)
        }} />
        <button className='bg-blue border text-white rounded-md p-2 md:ml-4' > Create room</button>
      </div>
      <div className='mt-6'>
        <div className='font-bold'>Available Rooms</div>
        <div className='grid grid-cols-1 md:grid-cols-5 gap-4 mt-6'>
          { rooms.map((room, index) => (
            <div
              key={index}
              className='border border-blue p-4 flex items-center rounded-md w-full'
            >
              <div className='w-full'>
                <div className='text-sm'>room</div>
                <div className='text-blue font-bold text-lg'>{room.name}</div>
              </div>
              <div className=''>
                <button
                  className='px-4 text-white bg-blue rounded-md'
                  onClick={() => joinRoom(room.id)}
                >
                  join
                </button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}

export default index
