import Head from 'next/head'
import { Inter } from '@next/font/google'
import Login from './screens/Login';
import SignUpForm from './screens/SignUpForm';
import Testr from '@/pages/screens/test';
import Sidebar from '@/components/Sidebar';



const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <div className="">
      <Head>
        <title>Freel</title>
      </Head>

      {/* Login*/}
      
      <Login />
      
      {/* Feed*/}
      {/* Modal*/}
    </div>
  );
}