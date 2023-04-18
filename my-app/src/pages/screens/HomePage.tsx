//Im using this page as an test for how the homepage will look.

import React, { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import Link from "next/link";
import Head from 'next/head';


import ScrollingView from '@/components/ScrollView';
import TopBar from '@/components/topbar';


function HomePage() {
  return (
    <div>
      <Head>
      <TopBar/>
        <title>My App</title>
      </Head>
     

      <div className="flex justify-center ">
      <ScrollingView/>
      </div>

      

       
      
    </div>
  );
}

export default HomePage;
