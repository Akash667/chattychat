import React from 'react'

const page = () => {
  return (
    <div className="flex items-center justify-center min-w-full min-h-screen">
        <form className="flex flex-col md:w-1/5">
            <input placeholder='email'    className='p-3 mt-8 rounded-md border-2 border-grey focus:outline-none focus:border-blue' />
            <input type="password "placeholder='password' className='p-3 mt-4 rounded-md border-2 border-grey focus:outline-none focus:border-blue' />
            <button
            className='p-3 mt-6 rounded-md bg-blue font-bold text-white'
            >Login</button>
        </form>
    </div>
  )
}

export default page
