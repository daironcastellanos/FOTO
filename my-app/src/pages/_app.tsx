import { BrowserRouter as Router, Route } from 'react-router-dom';
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import Script from 'next/script';


const App = ({ Component, pageProps }: AppProps) => (
  
  /* Global google*/
  <>
    <Component {...pageProps} /><Script
      src="https://accounts.google.com/gsi/client"
      async
      defer /></>
);

export default App;