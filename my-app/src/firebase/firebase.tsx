import { initializeApp } from "firebase/app";
import {getAuth} from '@firebase/auth'
import { getFirestore } from "@firebase/firestore";
import { GoogleAuthProvider } from "firebase/auth";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: "AIzaSyDWpNxp0sgOG2mNp13ya5TKLjqQ9hWgPSk",
  authDomain: "freel-376723.firebaseapp.com",
  databaseURL: "https://freel-376723-default-rtdb.firebaseio.com",
  projectId: "freel-376723",
  storageBucket: "freel-376723.appspot.com",
  messagingSenderId: "212725659815",
  appId: "1:212725659815:web:ff6ca6c2590c36f5ab8382"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
const db = getFirestore(app);
const provider = new GoogleAuthProvider();

export {app,auth,db,provider};