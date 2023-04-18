import Link from 'next/link';
import Image from 'next/image';
import React from 'react';
import logo from '../../public/images/Freel.png';

type HeaderProps = {
  isLoggedIn: boolean;
};

const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
  e.preventDefault();
  try {
    if (sessionStorage.getItem('loggedIn') != null) {
      sessionStorage.removeItem('loggedIn');
      window.location.reload();
    } else {
      console.log('user has not signed in.');
    }
  } catch (error) {
    console.error('Error:', error);
  }
};

const Header = ({ isLoggedIn }: HeaderProps) => {
  return (
    <div>
      {isLoggedIn ? <SignoutHeader /> : <SigninHeader />}
    </div>
  );
};

const SigninHeader = () => {
  return (
    <header className="bg-gradient-to-r from-blue-500 via-purple-500 to-pink-500 sticky top-0 z-10">
      <nav className="mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex items-center justify-between h-20">
          <div className="flex items-center">
            <Link href="/">
              <div className="w-20 cursor-pointer">
                <Image
                  className="object-contain"
                  src={logo.src}
                  alt="Logo"
                  width={150}
                  height={40}
                  layout="responsive"
                />
              </div>
            </Link>
          </div>
          <div className="hidden md:block">
            <div className="ml-10 flex items-baseline space-x-4">
              <Link href="/">
                <span className="text-white text-xl font-semibold hover:text-blue-300 cursor-pointer">Home</span>
              </Link>
              <Link href="/features">
                <span className="text-white text-xl font-semibold hover:text-blue-300 cursor-pointer">Features</span>
              </Link>
              <Link href="/about">
                <span className="text-white text-xl font-semibold hover:text-blue-300 cursor-pointer">About</span>
              </Link>
              <Link href="/follow">
                <span className="text-white text-xl font-semibold hover:text-blue-300 cursor-pointer">Follow</span>
              </Link>
            </div>
          </div>
          <div className="md:flex items-center justify-end md:flex-1 lg:w-0">
            <Link href="/screens/Login">
              <span className="whitespace-nowrap text-xl font-semibold text-white hover:text-blue-300 cursor-pointer">
                Sign in
              </span>
            </Link>
            <Link href="/screens/SignUpForm">
              <span className="whitespace-nowrap ml-4 inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-xl font-semibold text-white bg-yellow-400 hover:bg-yellow-500 cursor-pointer">
                Get started
              </span>
            </Link>
          </div>
        </div>
      </nav>
    </header>
  );
};


const SignoutHeader = () => {
  return (
    <header className="bg-white py-4 px-4 sm:px-6 lg:px-8 shadow-md">
      <nav className="relative flex items-center justify-between">
        <Link href="/">
          <a>
            <Image src={logo.src} alt="Logo" width={200} height={50} layout="intrinsic" />
          </a>
        </Link>
        <div className="hidden md:flex space-x-10">
          <Link href="/">
            <a className="text-gray-700 font-medium hover:text-blue-500">Home</a>
          </Link>
          <Link href="/features">
            <a className="text-gray-700 font-medium hover:text-blue-500">Features</a>
          </Link>
          <Link href="/about">
            <a className="text-gray-700 font-medium hover:text-blue-500">
            About</a>
</Link>
<Link href="/follow">
<a className="text-gray-700 font-medium hover:text-blue-500">Follow</a>
</Link>
</div>
<div className="md:flex items-center justify-end md:flex-1 lg:w-0">
<form onSubmit={handleSubmit}>
<button
           className="whitespace-nowrap inline-flex items-center justify-center px-4 py-2 border border-transparent rounded-md shadow-sm text-base font-medium text-white bg-blue-600 hover:bg-blue-700"
           type="submit"
         >
Sign out
</button>
</form>
</div>
</nav>
</header>
);
};

export default Header;