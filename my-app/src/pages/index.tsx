import Head from 'next/head';
import { useState, useEffect } from 'react';
import Link from 'next/link';
import Header from '@/components/Header';
import Footer from '@/components/Footer';

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

      <hr className="border-t-2 border-gray-300" />

      <div className="bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 min-h-screen flex flex-col justify-center items-center text-white">
        <div className="text-center space-y-4">
          <h1 className="text-6xl font-serif">Welcome to FREEL</h1>
          <h2 className="text-2xl font-semibold">Find your perfect photographer</h2>
        </div>
        <div className="mt-10">
          {isLoggedIn ? (
            <Link href="/compare">
              <span className="inline-block bg-green-500 hover:bg-green-600 py-4 px-6 text-2xl font-semibold rounded-lg transition-colors duration-300">
                Compare Now!
              </span>
            </Link>
          ) : (
            <Link href="/screens/SignUpForm">
              <span className="inline-block bg-green-500 hover:bg-green-600 py-4 px-6 text-2xl font-semibold rounded-lg transition-colors duration-300">
                Get started
              </span>
            </Link>
          )}
        </div>
      </div>

      <Footer isLoggedIn={isLoggedIn} />
    </div>
  );
}
