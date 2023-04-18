import React from 'react';
import Link from 'next/link';

type FooterProps = {
  isLoggedIn: boolean;
};

const Footer: React.FC<FooterProps> = ({ isLoggedIn }) => {
  const currentYear = new Date().getFullYear();
    return (
      <footer className="bg-gray-800 text-white py-12">
        <div className="container mx-auto px-3 grid grid-cols-1 md:grid-cols-3 gap-8">
          <div>
            <h3 className="font-bold text-xl mb-4">Freel</h3>
            <p>
              Freel is a platform that connects freelancers with clients, making it easy to find the right talent for your projects. We're dedicated to providing a seamless experience for both freelancers and clients to collaborate and grow.
            </p>
          </div>
          <div className="justify-center mx-auto">
            <h3 className="font-bold text-xl mb-4">Quick Links</h3>
            <ul>
              <li className="mb-2">
                <Link href="/">
                  <h3 className="hover:text-blue-400">Home</h3>
                </Link>
              </li>
              <li className="mb-2">
                <Link href="/features">
                  <h3 className="hover:text-blue-400">Features</h3>
                </Link>
              </li>
              <li className="mb-2">
                <Link href="/about">
                  <h3 className="hover:text-blue-400">About</h3>
                </Link>
              </li>
              <li className="mb-2">
                <Link href="/follow">
                  <h3 className="hover:text-blue-400">Follow</h3>
                </Link>
              </li>
            </ul>
          </div>
          <div>
            <h3 className="font-bold text-xl mb-4">Quick Contact</h3>
            <form>
              <input
                type="text"
                className="w-full bg-white text-black px-3 py-2 mb-4 rounded"
                placeholder="Name"
              />
              <input
                type="email"
                className="w-full bg-white text-black px-3 py-2 mb-4 rounded"
                placeholder="Email address"
              />
              <button className="bg-blue-600 hover:bg-blue-700 text-white font-bold px-4 py-2 rounded">
                Submit
              </button>
            </form>
          </div>
        </div>
        <div className="mt-8 border-t-2 border-gray-700">
          <div className="container mx-auto px-4 py-4 flex flex-wrap justify-between items-center">
            <p className="text-sm">© {new Date().getFullYear()} Freel. All rights reserved.</p>
            <div className="flex space-x-4">
              <Link href="/privacy">
                <h3 className="text-sm hover:text-blue-400">Privacy Policy</h3>
              </Link>
              <Link href="/terms">
                <h3 className="text-sm hover:text-blue-400">Terms of Service</h3>
                </Link>
      </div>
    </div>
  </div>
</footer>
);
};

export default Footer;
  