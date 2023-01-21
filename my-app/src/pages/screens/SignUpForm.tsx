import React, { useState } from 'react';

const Signup = () => {
    const [fullName, setFullName] = useState('');
    const [username, setUsername] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [dob, setDOB] = useState('');

    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        // Handle form submission here, such as sending a request to your server's signup endpoint
    }

    return (
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center">
            <div className="bg-white p-6 rounded-lg">
                <h1 className="text-2xl font-medium mb-4 text-purple-600">Sign up for Freel</h1>

                <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="fullName">
                            Full Name
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="text" 
                            id="fullName" 
                            value={fullName} 
                            onChange={e => setFullName(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="username">
                            Username
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="text" 
                            id="username" 
                            value={username} 
                            onChange={e => setUsername(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="email">
                            Email
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="email" 
                            id="email" 
                            value={email} 
                            onChange={e => setEmail(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="password">
                            Password
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="password" 
                            id="password" 
                            value={password} 
                            onChange={e => setPassword(e.target.value)} 
                            required 
                        />
                         </div>
                         <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="dob">
                            Date of Birth
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="date" 
                            id="dob" 
                            value={dob} 
                            onChange={e => setDOB(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className='space-x-2'>
                    <button className="bg-red-200 text-gray-800 py-2 px-7 rounded-lg hover:bg-indigo-600" onClick={() => window.history.back()}>
                        Back   
                    </button>
                    <button className="bg-indigo-500 text-white py-2 px-4 rounded-lg  hover:bg-indigo-600">
                        Register
                    </button>
                    </div>
                    

                    
                  </form>
            </div>
        </div>
    )
}

export default Signup;





