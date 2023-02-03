// This component is a simple login form that uses React hooks to manage its state
import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import Link from "next/link";
import {
    getAuth,
    signInWithEmailAndPassword,
    setPersistence,
    browserSessionPersistence,
  } from "firebase/auth";
  import { app } from "../../config/firebaseConfig";
//import {app} from

// Import the styles for this component
const MENU_LIST = [
    { text: "Home", href: "/" },
    { text: "About Us", href: "/about" },
    { text: "Contact", href: "/contact" },
  ];

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

            <Link href="./screens/Login">
                <img src="./images/Logo.png" alt="" width="249" height="107"/>
            </Link>
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
                    <button className="bg-indigo-500 text-white py-2 px-8 rounded-lg hover:bg-indigo-600">
                        Log in
                    </button>
                    <Link href="/screens/SignUpForm">
                    <button  className="block mt-4 text-center text-indigo-500">
                        Don't have an account? Sign up
                    </button>
                    </Link>
                </form>
            </div>
        </div>
    )
}



// Export the component
export default Login;
