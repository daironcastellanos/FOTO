import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import PhotoFeed from '@/components/Photofeed';
import ScrollingView from '@/components/ScrollView';
import Menu from '@/components/Menu';
import SuggestedFollowers from '@/components/SuggestedFollowers';
import Header from '@/components/Header';
import { GiHamburgerMenu } from 'react-icons/gi';

const Home: React.FC = () => {
  const [isClient, setIsClient] = useState(false);
  const [showSuggestedFollowers, setShowSuggestedFollowers] = useState(false);

  useEffect(() => {
    setIsClient(true);
  }, []);

  const handleHamburgerClick = () => {
    setShowSuggestedFollowers(!showSuggestedFollowers);
  };

  return (
    <div className="relative grid grid-cols-12 gap-4">
      {/* Left column: Menu */}
      <Menu />

      {/* Middle column: Feed */}
      <div
        className={`${
          showSuggestedFollowers || typeof window === 'undefined' ? 'col-span-6' : 'col-span-9'
        } ${
          showSuggestedFollowers || typeof window === 'undefined' ? 'col-start-4' : 'col-start-3'
        } bg-white p-4 max-h-screen`}
      >
        <h3 className="text-gray-800 font-bold mb-4 text-center">Photography Feed</h3>
        <div className="bg-white p-4 max-h-screen overflow-y-auto">
          <ScrollingView />
        </div>
      </div>

      {/* Right column: Suggested followers */}
      {isClient && (
        <div
          className={`${
            showSuggestedFollowers || typeof window === 'undefined'
              ? 'block'
              : 'hidden'
          } md:block col-span-3 bg-gray-100 p-4 min-h-screen border-l border-gray-300`}
        >
          {typeof window !== 'undefined' && (
            <button
              className="hamburger p-1 mb-4 bg-blue-600 text-white rounded flex items-center justify-center"
              onClick={handleHamburgerClick}
            >
              {showSuggestedFollowers ? (
                <span className="mr-2">Close</span>
              ) : (
                <GiHamburgerMenu className="text-2xl" />
              )}
            </button>
          )}
          <SuggestedFollowers />
        </div>
      )}
    </div>
  );
};

export default Home;
