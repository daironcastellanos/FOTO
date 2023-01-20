// This component is a simple login form that uses React hooks to manage its state
import React, { useState } from 'react';

// Import the styles for this component
const Login = () => {
    // Use React hooks to manage the state of the form
    const [email, setEmail] = useState('');
    // The second argument to useState is the function that will be used to update the state
    const [password, setPassword] = useState('');
    // Handle form submission
    const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        // Handle form submission here, such as sending a request to your server's login endpoint
    }
    // Render the form
    return (
        

        // Use Tailwind CSS to style the form
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center">
            <div className="bg-white p-6 rounded-lg">

            
                <h1 className="text-2xl font-medium mb-4 text-purple-600 ">Welcome to Freel</h1>
      
                
                <form onSubmit={handleSubmit}>
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
                    <button className="bg-indigo-500 text-white py-2 px-4 rounded-lg hover:bg-indigo-600">
                        Log in
                    </button>
                </form>
            </div>
        </div>
    )
}

// Export the component
export default Login;
