import { BrowserRouter as Router, Route } from 'react-router-dom';
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import ProtectedRoute from "../Auth/ProtectedRoute";
import { app } from "../config/firebaseConfig";
import { getAuth, onAuthStateChanged } from "firebase/auth";


const App = ({ Component, pageProps }: AppProps) => (
  
  <Component {...pageProps} />
);

export default App;