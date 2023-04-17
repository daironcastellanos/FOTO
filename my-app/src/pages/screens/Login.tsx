import React, { useState } from 'react';
import { useRouter } from 'next/router';
import Link from "next/link";
import { getAuth, signInWithEmailAndPassword } from 'firebase/auth'
import { auth } from '../../firebase/firebase'
import Image from 'next/image';

// Import the styles for this component
const MENU_LIST = [
    { text: "Home", href: "/" },
    { text: "About Us", href: "/about" },
    { text: "Contact", href: "/contact" },
];

const Login = () => {
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
          if (userIsAuth) {
            console.log(userIsAuth);
            console.log('User is Authorized ');

            router.push('/screens/HomePage'); // Use Next.js router to go to the Profile screen
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


        // Use Tailwind CSS to style the form
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center">
            <div className="bg-white p-6 rounded-lg">

                <div>
                <Image src="/images/Logo.png" alt="Your alt text" width={249} height={107} />
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
                    <button className="bg-indigo-500 text-white py-2 px-8 rounded-lg hover:bg-indigo-600">
                        Log in
                    </button>
                    <Link href="/screens/SignUpForm">
                        <button className="block mt-4 text-center text-indigo-500">
                            Dont have an account? Sign up
                        </button>
                    </Link>

                    <Link href="/screens/GoogleSignIn">
                        <button className="block mt-4 text-center text-indigo-500">
                            Sign In with Google
                        </button>
                    </Link>
                </form>

            </div>
        </div>
    );
};



// Export the component
export default Login;
