import React, { useEffect, useState } from 'react'
import FadeLoader from 'react-spinners/FadeLoader';

function AuthLoader() {

  let [showPrompt, setShowPrompt] = useState(false);

  let authTimeout = () => {
    setTimeout(function() {
        setShowPrompt(true)
    });
  }
  useEffect(() => {
    authTimeout();
  }, [])
   return (
    showPrompt ? (<div> Try logging in?</div>) :
    <FadeLoader
        color={'black'}
        loading={true}
    />
  )
}

export default AuthLoader