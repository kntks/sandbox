import type { NextPage } from 'next'
import { useSession, signIn, signOut } from "next-auth/react"

import { Button } from '@nextui-org/react';


const Home: NextPage = () => {
  const { data: session, status } = useSession()

  console.log("session data", session)
  if(session) {
    return (
      <>
        Signed in as {session.user?.name} <br />
        <Button onClick={() => signOut()}>Log Out</Button>
      </>
    )
  }
  return (
    <>
      <div>Example App</div>
      <Button onClick={() => signIn()}>Log In</Button>
    </>
  )
}

export default Home
