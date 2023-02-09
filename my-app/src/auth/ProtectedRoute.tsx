import {Navigate, Outlet} from 'react-router-dom';
import {app} from '../firebase/firebase';
import { getAuth, onAuthStateChanged } from 'firebase/auth';
import { useEffect, useState } from 'react';
import AuthLoader from './AuthLoader';

const ProtectedRoute = ({
    redirectPath = '/login',
  }) => {
    let auth = getAuth(app);
    const [user, setUser] = useState(auth.currentUser)

    useEffect(() => {
      auth.onAuthStateChanged(currUser => {
        setUser(currUser);
      });
    }, [])
  
    return user && user.emailVerified ? <Outlet /> : <AuthLoader/>;
};


export default ProtectedRoute;