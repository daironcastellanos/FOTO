//Im using this page as an test for how the homepage will look.

import React, { useState } from 'react';
import Head from 'next/head';
import ScrollingView from '@/components/ScrollView';
import TopBar from '@/components/topbar';

function HomePage() {
  return (
    <div>
      <Head>
     
        <title>My App</title>
      </Head>
     
      <TopBar/>

      <div className="flex justify-center ">
      <ScrollingView/>
      </div>
    </div>
  );
}

export default HomePage;
