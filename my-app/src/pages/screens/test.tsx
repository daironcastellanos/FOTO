import React from 'react'
import Link from 'next/link'

function Testr() {
  return (
    <header>
        <div>
            <Link href="./screens/Login">
                <img src="./images/Logo.png" alt="" width="249" height="107"/>
            </Link>
            
            <h1>Signed in</h1>
        </div>
        <div></div>
    </header>
  )
}

export default Testr

