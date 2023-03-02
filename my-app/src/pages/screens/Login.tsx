// This component is a simple login form that uses React hooks to manage its state
import React, { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import Link from "next/link";
import {getAuth, signInWithEmailAndPassword} from 'firebase/auth'
import {auth} from '../../firebase/firebase'

// Import the styles for this component
const MENU_LIST = [
    { text: "Home", href: "/" },
    { text: "About Us", href: "/about" },
    { text: "Contact", href: "/contact" },
  ];

const Login = () => {

    let userIsAuth = null;
    // Use React hooks to manage the state of the form
    const [email, setEmail] = useState('');
    // The second argument to useState is the function that will be used to update the state
    const [password, setPassword] = useState('');
    // Handle form submission
    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        // Handle form submission here, such as sending a request to your server's login endpoint
       
        if(email.length && password.length){
            try{
                /* */
                const userCredential = await signInWithEmailAndPassword(auth,email,password);
                userIsAuth = userCredential.user;
                setEmail("");
                setPassword("");
                //const userInfo = await getUser();


                // Simple print to check authentication 
                if(userIsAuth){
                    console.log('User is Authorized ')

                }else{
                    console.log("not authorized")
                }

            }catch{
                /* This is what executes when the login is invalid*/
                console.log("error",(e));

            }
        }
    };
    // Render the form
    return (
        

        // Use Tailwind CSS to style the form
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center">
            <div className="bg-white p-6 rounded-lg">

            <div>
                <img src="./images/Logo.png" alt="" width="249" height="107"/>
            </div>
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
                    <Link href="/screens/HomePage">
                    <button className="bg-indigo-500 text-white py-2 px-8 rounded-lg hover:bg-indigo-600">
                        Log in
                    </button>
                    </Link>
                    <Link href="/screens/SignUpForm">
                    <button  className="block mt-4 text-center text-indigo-500">
                        Don't have an account? Sign up
                    </button>
                    </Link>

                    <Link href="/screens/GoogleSignIn">
                    <button  className="block mt-4 text-center text-indigo-500">
                        Sign In with Google
                    </button>
                    </Link>
                </form>
                
            </div>
        </div>
    )
}



// Export the component
export default Login;
