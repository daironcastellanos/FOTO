import React, { useState, useRef, useContext } from "react";
import { app, db } from "../../firebase/firebase";
import { getAuth } from "@firebase/auth";
import { addDoc, collection } from "firebase/firestore";
import Link from "next/link";
import {auth} from '../../firebase/firebase'
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { useRouter } from "next/router";
import { CreateUserInMongo } from "@/call_bakend/working/backend";

const Signup = () => {
    const auth = getAuth(app);
  
    const [fullName, setFullName] = useState("");
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [dob, setDOB] = useState("");
    const [bio, setBio] = useState("");
    const [location, setLocation] = useState({ Coordinates: [0.0, 0.0] });
    const router = useRouter();
  
    const handleSubmit = async (e: { preventDefault: () => void }) => {
      e.preventDefault();
  
      if (email.length && password.length) {
        console.log("Creating user with email and password...");
        const { user } = await createUserWithEmailAndPassword(
          auth,
          email,
          password
        );
  
        const userData = {
          FireID: user.uid,
          FullName: fullName,
          Username: username,
          Email: email,
          Bio: bio,
          Location: location,
          DOB: dob,
          Followers: [],
          Following: [],
          MyPhotos: [],
          SavedPhotos: [],
        };
  
        // Send a request to your Go backend using Axios
        if (await CreateUserInMongo(userData)) {
          router.push("/screens/Login");
        } else {
          console.log("Error creating user in mongo");
        }
      }
    };
  
    const getLocation = () => {
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(
          (position) =>
            setLocation({
              Coordinates: [position.coords.latitude, position.coords.longitude],
            }),
          (error) => console.error("Error getting location:", error)
        );
      } else {
        console.error("Geolocation is not supported by this browser.");
      }
    };
        

       


    return (
        <div className="fixed top-0 left-0 h-screen w-screen flex items-center justify-center bg-gradient-to-r from-blue-200 via-green-200 to-purple-200">
          <Link href="/">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
          <div className="bg-white p-8 rounded-xl shadow-md">
            <h1 className="text-3xl font-medium mb-6 text-purple-600">Sign up for Freel</h1>
    
            <form onSubmit={handleSubmit}>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="fullName">
                            Full Name
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="text" 
                            id="fullName" 
                            value={fullName} 
                            onChange={e => setFullName(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="username">
                            Username
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="text" 
                            id="username" 
                            value={username} 
                            onChange={e => setUsername(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="email">
                            Email
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="email" 
                            id="email" 
                            value={email} 
                            onChange={e => setEmail(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="password">
                            Password
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="password" 
                            id="password" 
                            value={password} 
                            onChange={e => setPassword(e.target.value)} 
                            required 
                        />
                         </div>
                         <div className="mb-4">
                        <label className="block text-gray-700 font-medium mb-2" htmlFor="dob">
                            Date of Birth
                        </label>
                        <input 
                            className="border border-gray-400 p-2 rounded-lg w-full" 
                            type="date" 
                            id="dob" 
                            value={dob} 
                            onChange={e => setDOB(e.target.value)} 
                            required 
                        />
                    </div>
                    <div className="flex justify-center items-center mt-6">
            
            <button className="bg-indigo-500 text-white py-2 px-10 rounded-lg  hover:bg-indigo-600">
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Signup;



