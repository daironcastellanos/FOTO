import React, { useEffect, useState } from 'react';
import FadeLoader from 'react-spinners/FadeLoader';

interface Props {
  wait: number;
}

const LoadAuth: React.FC<Props> = ({ wait }) => {
  const [showPrompt, setShowPrompt] = useState(false);

  const authTimeout = () => {
    setTimeout(() => {
      setShowPrompt(true);
    }, wait);
  };

  useEffect(() => {
    authTimeout();
  }, []);

  return showPrompt ? (
    <div>Try logging in?</div>
  ) : (
    <FadeLoader color="black" loading={true} />
  );
};

export default LoadAuth;