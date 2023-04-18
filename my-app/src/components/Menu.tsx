import React from 'react';
import Link from 'next/link';
import logo from '../../public/images/Freel.png';
import Image from 'next/image';

const Menu: React.FC = () => {
  return (
    <div className="col-span-3 bg-gray-100 p-4 min-h-screen border-r border-gray-300">
      <div className="flex flex-col items-center justify-between h-full">
        <div>
          <div className="w-40 flex justify-start items-center mb-8 ml-4">
            <Image
              className="object-contain"
              src={logo.src}
              alt="Logo"
              width={120}
              height={24}
              layout="responsive"
            />
          </div>
          <Link href="/Home">
            <span className="menu-link text-blue-600 font-semibold text-2xl hover:text-blue-500 text-center flex items-center justify-center space-x-2">
              <i className="fas fa-home"></i>
              <span>Home</span>
            </span>
          </Link>
          <Link href="/Profile">
            <span className="menu-link text-gray-700 font-semibold text-xl hover:text-blue-600 text-center flex items-center justify-center space-x-2">
              <i className="fas fa-user"></i>
              <span>Profile</span>
            </span>
          </Link>
          <Link href="/upload">
            <span className="menu-link text-gray-700 font-semibold text-xl hover:text-blue-600 text-center flex items-center justify-center space-x-2">
              <i className="fas fa-upload"></i>
              <span>Upload</span>
            </span>
          </Link>
          <Link href="/Settings">
            <span className="menu-link text-gray-700 font-semibold text-xl hover:text-blue-600 text-center flex items-center justify-center space-x-2">
              <i className="fas fa-cog"></i>
              <span>Settings</span>
            </span>
          </Link>
          <Link href="/search">
            <span className="menu-link text-gray-700 font-semibold text-xl hover:text-blue-600 text-center flex items-center justify-center space-x-2">
              <i className="fas fa-search"></i>
              <span>Search</span>
            </span>
          </Link>
        </div>
        <div className="mb-4">
        <Link href="/">
          <button className="text-gray-700 font-semibold text-xl hover:text-red-600 text-center flex items-center justify-center space-x-2">
            <i className="fas fa-sign-out-alt"></i>
            <span>Sign Out</span>
          </button>
          </Link>
        </div>
        <style jsx>{`
        .menu-link {
          padding: 5px 0;
          display: block;
        }
      `}</style>
    </div>

    </div>
  );
};

export default Menu;
