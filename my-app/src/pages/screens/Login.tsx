import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from "next/link";
import { getAuth , signInWithEmailAndPassword } from 'firebase/auth'
import { auth } from '../../firebase/firebase'
import Image from 'next/image';

// Import the styles for this component
const MENU_LIST = [
    { text: "Home", href: "/" },
    { text: "About Us", href: "/about" },
    { text: "Contact", href: "/contact" },
];

const Login = () => {

  

const getUser = async () => {
  const user = auth.currentUser;
  console.log(user?.uid);
  console.log(user?.email);

  return user;
};

const getUid = async () => {
  const user = await getUser();
  const fireID = user?.uid;

  return fireID;
};

// Wrap the code in an async function to use 'await' inside
const checkAuthentication = async () => {
  const uid = await getUid();

  if (uid !== "") {
    console.log('User is Authorized');
    router.push('/screens/HomePage'); // Use Next.js router to go to the HomePage screen
  } else {
    console.log('not authorized');
    router.push('/screens/SignupForm'); // Use Next.js router to go to the SignupForm screen
  }
};

// Call the function to check authentication


    const router = useRouter(); // Add this line to use Next.js router
    let userIsAuth = null;
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
  
    const handleSubmit = async (e: { preventDefault: () => void; }) => {
      e.preventDefault();
  
      if (email.length && password.length) {
        try {
          const userCredential = await signInWithEmailAndPassword(auth, email, password);
          userIsAuth = userCredential.user;
          setEmail('');
          setPassword('');
          
          if (await checkAuthentication) {
            console.log(userIsAuth);
            console.log('User is Authorized ');
            router.push('/Home'); // Use Next.js router to go to the Profile screen
          } else {
            console.log('not authorized');
            router.push('/screens/SignupForm'); // Use Next.js router to go to the SignupForm screen
          }
        } catch {
          console.log('error', e);
        }
      }
    };
    // Render the form
    return (
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center bg-gradient-to-r from-blue-200 via-green-200 to-purple-200">
          <Link href="/">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
          <div className="bg-white p-8 rounded-xl shadow-md">
            <div className="flex justify-center">
              <Image src="/images/Logo.png" alt="Logo" width="249" height="107"/>
            </div>
            <h1 className="text-3xl font-medium mb-6 text-purple-600 text-center">Welcome to Freel</h1> 
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
                    <div className="flex justify-center mt-4">
              
              <button className="bg-indigo-500 text-white py-2 px-6 rounded-lg hover:bg-indigo-600"
              onClick={handleSubmit}
              >
                Log in
              </button>
          </div>
          <div className="text-center mt-4">
            <Link href="/screens/SignUpForm">
              <button className="text-indigo-500">
                Dont have an account? Sign up
              </button>
            </Link>
          </div>
          <div className="text-center mt-4">
            <Link href="/screens/GoogleSignIn">
              <button className="text-indigo-500">
                Sign In with Google
              </button>
            </Link>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;