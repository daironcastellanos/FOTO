//Im using this page as an test for how  the homepage will look.

import React, { useState } from 'react';
import { Navigate, useNavigate } from 'react-router-dom';
import Link from "next/link";
import Head from 'next/head';


const HomePage = () => {
  return (
    <div>
      <Head>
        <title>My App</title>
      </Head>
      <h1>Welcome to My App!</h1>
      <p>Please login to continue.</p>
      
    </div>
  );
};

export default HomePage;


