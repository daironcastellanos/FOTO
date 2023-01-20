import Head from 'next/head'
import { Inter } from '@next/font/google'
import Login from '../components/Login';


const inter = Inter({ subsets: ['latin'] })

export default function Home() {
  return (
    <div className="">
      <Head>
        <title>Freel</title>
      </Head>

      
      

      {/* Login*/}
      <Login/>
      
      {/* Feed*/}
      {/* Modal*/}
    </div>
  );
}