import React, { useEffect, useState } from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { app } from '../config/firebaseConfig';
import { User, getAuth, onAuthStateChanged } from 'firebase/auth';
import LoadAuth from './LoadAuth';

interface Props {
  redirectPath?: string;
  children?: React.ReactNode;
}

const ProtectedRoute: React.FC<Props> = ({ redirectPath = '/login', children }) => {
  const auth = getAuth(app);
  const [user, setUser] = useState<User | null>(auth.currentUser);

  useEffect(() => {
    const unsubscribe = auth.onAuthStateChanged((currUser) => {
      setUser(currUser);
    });
    return () => {
      unsubscribe();
    };
  }, []);

  return user && user.emailVerified ? <Outlet /> : <LoadAuth wait={3000} />;
};

export default ProtectedRoute;
