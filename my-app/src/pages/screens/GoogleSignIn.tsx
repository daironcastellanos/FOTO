
import {useEffect, useState } from 'react';
import jwt_decode from "jwt-decode";

declare var google: any;
function GoogleSignIn() {
  const[ user, setUser ] = useState({});

  function handleCallbackResponse(response) {
    console.log("Encoded JWT ID token: " + response.credential);
    var userObject = jwt_decode(response.credential);
    console.log(userObject);
    setUser(userObject);
    document.getElementById("signInDiv").hidden = true;
  }

  function handleSignOut(event){
    setUser({});
    document.getElementById("signInDiv").hidden = false;
  }

  useEffect(() => {
    console.log("before IF");
    if (typeof google !== 'undefined') {
      console.log("test")
      google.accounts.id.initialize({
        client_id:"212725659815-f59aick9qoksihrb9qhvhk9evtr2qks7.apps.googleusercontent.com",
        callback: handleCallbackResponse
      });

      google.accounts.id.renderButton(
        document.getElementById("signInDiv"),
        { theme: "outline", size:"large"}
      );

      google.accounts.id.prompt();
    }
  }, []);

  // 

  return (
    <div className="App">
      <div id="signInDiv"></div>
      { Object.keys(user).length != 0 &&
        <button onClick={ (e) => handleSignOut(e)}>Sign Out</button>
      }
      
      { user &&
        <div>
          <img src={user.picture}></img>
          <h3>{user.name}</h3>
          </div>
      }
    </div>
  );
}

export default GoogleSignIn;
