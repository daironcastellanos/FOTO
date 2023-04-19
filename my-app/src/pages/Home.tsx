import React from 'react';
import Link from 'next/link';
import PhotoFeed from '@/components/Photofeed';
import ScrollingView from '@/components/ScrollView';
import Menu from '@/components/Menu';
import SuggestedFollowers from '@/components/SuggestedFollowers';
import Header from '@/components/Header';
import TopBarSearch from '@/components/topbarsearch';


const Home: React.FC = () => {
  return (
    <div className="relative grid grid-cols-12 gap-4">
      {/* Left column: Menu */}
     

      <Menu/>

      {/* Middle column: Feed */}
      <div className="col-span-6 col-start-4 bg-white p-4 max-h-screen">
        <h3 className="text-gray-800 font-bold mb-4 text-center">Photography Feed</h3>
      <div className="col-span-6 col-start-4 bg-white p-4 max-h-screen overflow-y-auto">
        <ScrollingView/>
      </div>
      </div>

      {/* Right column: Suggested followers */}
      <SuggestedFollowers/>
    </div>
  );
};

export default Home;
