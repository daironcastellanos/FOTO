import { BrowserRouter as Router, Route } from 'react-router-dom';
import '@/styles/globals.css'
import type { AppProps } from 'next/app'


const App = ({ Component, pageProps }: AppProps) => (
  
  <Component {...pageProps} />
);

export default App;