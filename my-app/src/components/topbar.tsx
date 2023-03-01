import React from 'react';
import Link from "next/link";

const TopBar = () => {
  return (
    <div className="bg-gray-800 h-16 flex justify-between items-center px-4">
      <div>
        <h1 className="text-white text-2xl font-bold">FREEL</h1>
      </div>
      <div>
      <Link href="/screens/Profile">
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-5 rounded mr-4">Profile</button>
        </Link>
       
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-5 rounded mr-4">Upload</button>
        <button className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-5 rounded mr-4">Settings</button>
      </div>
    </div>
  );
}

export default TopBar;
