import { useEffect, useState } from "react";
import jwt_decode from "jwt-decode";
import Image from "next/image";

declare var google: any;

interface User {
  name?: string;
  picture?: string;
}

function GoogleSignIn() {
  const [user, setUser] = useState<User>({});

function handleCallbackResponse(response: { credential: string; }) {
    console.log("Encoded JWT ID token: " + response.credential);
    var userObject: User = jwt_decode(response.credential);
    console.log(userObject);
    setUser(userObject);
    const signInDiv = document.getElementById("signInDiv");
    if (signInDiv) {
        signInDiv.hidden = true;
    }
}

function handleSignOut() {
    setUser({});
    const signInDiv = document.getElementById("signInDiv");
    if (signInDiv) {
        signInDiv.hidden = false;
    }
}

  useEffect(() => {
    console.log("before IF");
    if (typeof google !== "undefined") {
      console.log("test");
      google.accounts.id.initialize({
        client_id:
          "212725659815-f59aick9qoksihrb9qhvhk9evtr2qks7.apps.googleusercontent.com",
        callback: handleCallbackResponse,
      });

      google.accounts.id.renderButton(
        document.getElementById("signInDiv"),
        { theme: "outline", size: "large" }
      );

      google.accounts.id.prompt();
    }
  }, []);

  return (
    <div className="App">
      <div id="signInDiv"></div>
      {Object.keys(user).length !== 0 && (
        <button onClick={(e) => handleSignOut()}>Sign Out</button>
      )}

      {user && (
        <div>
          {user.picture && (
            <Image src={user.picture} alt="user-image" width={48} height={48} />
          )}
          <h3>{user.name}</h3>
        </div>
      )}
    </div>
  );
}

export default GoogleSignIn;
