// pages/follow.tsx
import Head from 'next/head';
import Link from 'next/link';
import Header from '../components/Header';
import Image from 'next/image';

function Follow() {
  const developers = [
    {
      name: 'Sergio Arcila',
      position: 'Front-end Developer',
      bio: 'Sergio is a software engineer with over 10 years of experience. He is responsible for leading the front end team and ensuring the project stays on track.',
      image: '../images/serg.png',
    },
    {
      name: 'Dairon Castellanos',
      position: 'Front-end Developer',
      bio: 'Dairon is a skilled front-end developer with a strong background in creating efficient and responsive user interfaces. He has a deep understanding of React and works closely with the rest of the team to deliver high-quality results.',
      image: 'https://via.placeholder.com/150',
    },
    {
      name: 'Jose Simon',
      position: 'Back-end Developer',
      bio: 'Jose is a talented back-end developer who specializes in creating robust and scalable APIs. He has a strong background in various back-end technologies and is responsible for leading the back-end team.',
      image: 'https://via.placeholder.com/150',
    },
    {
      name: 'Eric Dequevedo',
      position: 'Back-end Developer',
      bio: 'Eric is a dedicated back-end developer with a deep understanding of server-side technologies and cloud infrastructure. He ensures the smooth deployment and maintenance of the project.',
      image: 'https://via.placeholder.com/150',
    },
  ];

  let isLoggedIn: boolean = false;

  if (typeof window !== 'undefined' && window.sessionStorage) {
    isLoggedIn = sessionStorage.getItem('loggedIn') != null;
  } else {
    console.warn('sessionStorage is not available.');
  }

  return (
    <div className="mx-auto">
      <Head>
        <title>Follow - WhatIF</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header isLoggedIn={isLoggedIn}/>

      <div className="p-10 max-w-7xl mx-auto">
        <h1 className="text-4xl font-bold mb-10">Our Developers</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
          {developers.map((developer, index) => (
            <div key={index} className="bg-white p-6 rounded-lg shadow-md">
              <Image
                className="mx-auto w-32 h-32 rounded-full mb-4"
                src={developer.image}
                alt={developer.name}
              />
              <h2 className="text-2xl font-medium mb-2">{developer.name}</h2>
              <h3 className="text-lg text-gray-500 mb-4">{developer.position}</h3>
              <p className="text-lg">{developer.bio}</p>
            </div>
          ))}
        </div>
        <div className="mt-10">
          <Link href="/">
            <h1 className="text-blue-600 hover:underline">Back to Home</h1>
          </Link>
        </div>
      </div>
    </div>
  );
}

export default Follow;
