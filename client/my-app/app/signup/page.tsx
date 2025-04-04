
"use client"

import React, {useContext, useEffect, useState} from 'react'
import { API_URL } from '../../constants'
import { useRouter } from 'next/navigation'
import { AuthContext , UserInfo } from '@/modules/auth_provider'


const Page = () => {

  const [email , setEmail] = useState('')
  const [password , setPassword] = useState('')
  const [username , setUsername] = useState('')
  const { authenticated } = useContext(AuthContext)
  const router = useRouter()

  useEffect(() => {
    if (authenticated){
      router.push('/')
      return
    }
  },[authenticated])

  const submitHandler = async ( e: React.SyntheticEvent) => {
     e.preventDefault()
      try{
        const res = await fetch(`${API_URL}/signup`, {
          method:"POST",
          headers:{
            'Content-Type':'application/json'
          },
          body: JSON.stringify({email, password, username})
        })
        const data = await res.json()

        if (res.ok){
          const user: UserInfo = {
            username: data.username,
            id: data.id
          }
        console.log("user created as follows: ", user)
        //   localStorage.setItem('user_info', JSON.stringify(user))
          return router.push('/login')
        }


      }catch(err){
        console.log(err)
      }
  }

  return (

    <div className="flex items-center justify-center min-w-full min-h-screen">
        <form className="flex flex-col md:w-1/5">
            <div className='text-3xl font-bold text-center'>
              <span className='text-blue'>
                welcome! Sign up here.
              </span>
            </div>
            <input placeholder='Enter username'    className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
            onChange={(e) => setUsername(e.target.value)} />
            <input placeholder='Enter email'    className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
            onChange={(e) => setEmail(e.target.value)} />
            <input type="password "placeholder='password' className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue'
            onChange={(e) => setPassword(e.target.value)} />
            <button
            className='p-3 mt-6 rounded-md bg-blue font-bold text-white'
          onClick={submitHandler}
            >Login</button>
        </form>
    </div>

  )
}

export default  Page
