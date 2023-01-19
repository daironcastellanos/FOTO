import React, { useState } from "react";
import '../styles/Login.css';



interface Props {
}

export default function Login({}: Props) {
  // State to store the email and password input values
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  return (
    <form>
      <label>
        Email:
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
      </label>
      <br />
      <label>
        Password:
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />
      </label>
      <br />
      <button type="submit">Login</button>
    </form>
  );
}