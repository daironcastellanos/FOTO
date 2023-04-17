import { useEffect, useState } from 'react';
import Image from 'next/image';
import { getAuth, GoogleAuthProvider, signInWithPopup } from 'firebase/auth';
import { auth } from '../../firebase/firebase';

function GoogleSignIn() {
  const [user, setUser] = useState(null);

  const handleSignIn = async () => {
    const provider = new GoogleAuthProvider();
    try {
      const result = await signInWithPopup(auth, provider);
      console.log(result);
    } catch (error) {
      console.error(error);
    }
  };

  const handleSignOut = () => {
    auth.signOut().then(() => {
      setUser(null);
    }).catch((error) => {
      console.error(error);
    });
  };

  return (
    <div className="App">
      {user === null ? (
        <button onClick={handleSignIn}>Sign in with Google</button>
      ) : (
        <button onClick={handleSignOut}>Sign Out</button>
      )}

      {user && (
        <div>
         
        </div>
      )}
    </div>
  );
}

export default GoogleSignIn;
