import Head from 'next/head'
import { Inter } from '@next/font/google'
import Login from './screens/Login';
import SignUpForm from './screens/SignUpForm';
import Testr from 'src/pages/screens/test';
import Sidebar from '@/components/Sidebar';

export async function getServerSideProps() {
  const {status} = await fetch("http://localhost:8000/status").then(x => x.json());
  //const {username} = await fetch("http://localhost:8000/username").then(x => x.json());
  return {
    props: {
      status: status,
      //username: username,
    }
  }
}

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