// pages/features.tsx
import Head from 'next/head';
import Link from 'next/link';
import Header from '../components/Header';

function Features() {

  let isLoggedIn: boolean = false;

  if (typeof window !== 'undefined' && window.sessionStorage) {
    isLoggedIn = sessionStorage.getItem('loggedIn') != null;
  } else {
    console.warn('sessionStorage is not available.');
  }

  return (
    <div className="max-w-7xl mx-auto">
      <Head>
        <title>Features - Freel</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Header isLoggedIn={isLoggedIn}/>

      <div className="p-10">
        <h1 className="text-4xl font-bold mb-10">Freel's Best Features</h1>
        <div className="grid grid-cols-1 md:grid-cols-3 gap-10">
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-medium mb-4">Location-Based Services</h2>
            <p className="text-lg mb-4">
              Find photographers near you, wherever you are in the world. Our location-based services make it easy for clients to discover and hire photographers in their immediate vicinity.
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-medium mb-4">Portfolio Showcase</h2>
            <p className="text-lg mb-4">
              Photographers can showcase their work through personalized profiles. Clients can view portfolios, ensuring they find the perfect photographer for their needs.
            </p>
          </div>
          <div className="bg-white p-6 rounded-lg shadow-md">
            <h2 className="text-2xl font-medium mb-4">Integrated Booking System</h2>
            <p className="text-lg mb-4">
              Simplify the booking process with our integrated system. Photographers can display their availability, and clients can easily book their services directly through the platform.
            </p>
          </div>
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

export default Features;
