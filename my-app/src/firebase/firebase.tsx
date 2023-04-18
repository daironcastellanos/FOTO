import { initializeApp } from "firebase/app";
import {getAuth} from '@firebase/auth'
import { getFirestore } from "@firebase/firestore";
import { useState, useEffect, useContext, createContext } from 'react'
import { GoogleAuthProvider } from "firebase/auth";

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyAn5rjr-x7S6q5qs-g6IkyF9dHQgTNscYs",
  authDomain: "freel-ee39b.firebaseapp.com",
  projectId: "freel-ee39b",
  storageBucket: "freel-ee39b.appspot.com",
  messagingSenderId: "828543545656",
  appId: "1:828543545656:web:c692555e6b8d49c30d11cd",
  measurementId: "G-6KVJXG35XJ"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const auth = getAuth(app);
const db = getFirestore(app);
const provider = new GoogleAuthProvider();

export {app,auth,db,provider};