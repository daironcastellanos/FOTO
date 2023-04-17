import Head from 'next/head'
import { useState, useEffect } from 'react';
import { Inter } from '@next/font/google'
import Link from 'next/link';
import Login from './screens/Login';

import HomePage from './screens/HomePage';
import Profile from './screens/Profile';
import SignUpForm from './screens/SignUpForm';
import Testr from 'src/pages/screens/test';
import Sidebar from '@/components/Sidebar';
import Header from '@/components/Header';
import Footer from '@/components/Footer';



const inter = Inter({ subsets: ['latin'] })

export default function Home() {

  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    if (typeof window !== 'undefined' && window.sessionStorage) {
      setIsLoggedIn(sessionStorage.getItem('loggedIn') != null);
    } else {
      console.warn('sessionStorage is not available.');
    }
  }, []);
  return (
    <div className="">
      <Head>
        <title>Freel</title>
      </Head>

      <Header isLoggedIn={isLoggedIn} />
      
      <div className="flex justify-center items-center bg-blue-400 border-y border-black py-40">
        <div className="px-10 space-y-2">
          <h1 className="text-6xl max-w-xl font-serif">Welcome to FREEL</h1>
          <h2 className="text-xl">Find your perfect photographer</h2>

          <div className="max-w-xs py-10">
            {isLoggedIn ? (
              <Link href="/compare">
                <h1 className="text-3xl text-white bg-green-600 px-4 py-3 rounded text-center">
                  Compare Now!
                </h1>
              </Link>
            ) : (
              <Link href="/signup">
                <h1 className="text-3xl text-white bg-green-600 px-4 py-3 rounded text-center">
                  Get started
                </h1>
              </Link>
            )}
          </div>
        </div>

        <div></div>
      </div>
      <Footer isLoggedIn={isLoggedIn} />
     
    </div>
  );
}