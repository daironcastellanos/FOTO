// pages/about.tsx
import Head from 'next/head';
import Link from 'next/link';
import Header from '../components/Header';

function About() {

  let isLoggedIn: boolean = false;

  if (typeof window !== 'undefined' && window.sessionStorage) {
    isLoggedIn = sessionStorage.getItem('loggedIn') != null;
  } else {
    console.warn('sessionStorage is not available.');
  }

  return (
    <div className="max-w-7xl mx-auto bg-gray-50">
      <Head>
        <title>About Us - Freel</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div>
        <Header isLoggedIn={isLoggedIn}/>
      </div>

      <div className="p-10">
        <section className="bg-white py-8 px-6 rounded-lg shadow-md">
          <h2 className="text-3xl font-bold mb-6">Freel - The Future of Freelance Photography</h2>
          <p className="text-lg mb-4">
            At Freel, we believe in freedom. Freedom of life, freedom of choice, freedom for the freelancer. We aim to provide a service that facilitates and revolutionizes the way photographers conduct their lead generation. Our web-based application offers a unique platform for content creators to sign up and showcase their portfolio to a specific audience in a specific location.
          </p>
          <p className="text-lg mb-4">
            🌐 Imagine it as...
            A Google Maps for available photographers near you, but the difference is it can be anywhere in the world.
            Users in need of a photographer can sign up and search for available photographers in their immediate surrounding. This benefits the user by allowing them to find and hire photographers at a moment's notice while viewing their work through their profile, which showcases the photographer's portfolio.
          </p>
          <p className="text-lg mb-4">
            For photographers, our services are nothing shy of revolutionary. Now, content creators have the ability to travel the world and generate new business anywhere they are. Our location-based services allow well-known photographers to travel to the other side of the world, showcase their portfolio to users in that area, and create new business.
          </p>
          <p className="text-lg mb-4">
            Not only are photographers expanding their reach, but we also offer a booking system where photographers can show their availability and accept bookings right from the site.
          </p>
          <p className="text-lg mb-4">
            Photographers can use our site as a portfolio showcase and even choose their own unique username for clients to search.
          </p>
          <p className="text-lg mb-4">
            In short, Freel App is the ultimate platform for photographers to:
            ✅ Sign up 🌟 Upload and showcase their work 📅 Facilitate the booking process 🌍 Choose where they take their business
          </p>
          <p className="text-lg mb-4">
            Unlike most freelancing websites, we allow photographers to find clients who can meet them in real life. At Freel, the photographer can travel the world and have reach to a whole new set of clients. We feature available photographers in your immediate area, wherever that may be.
          </p>
          <p className="text-lg mb-4">
            In the future, we have plans to expand our reach to other types of local freelance services. We aim to create a marketplace where you can find local available services in your immediate area while allowing complete freedom to the
            creator.
</p>
<p className="text-lg mb-4">
By embracing the power of technology and innovation, Freel aims to reshape the way freelancers and clients connect, collaborate, and grow together. Our commitment to a user-friendly experience and dedication to continuous improvement will help establish Freel as a leading platform in the freelance industry.
</p>

<Link href="/">
        <h1 className="text-blue-600 hover:underline cursor-pointer">Back to Home</h1>
      </Link>
    </section>
  </div>
</div>
);
}

export default About;