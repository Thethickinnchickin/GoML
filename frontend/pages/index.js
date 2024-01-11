import Head from 'next/head'
import Image from 'next/image'
import styles from '@/styles/Home.module.css'
import { useRouter } from 'next/router'
import { checkTokenExists } from '@/helpers/auth'
import { useState, useEffect } from 'react'


export default function Home() {
  const [tokenExists, setTokenExists] = useState(false);


  useEffect(() => {
    // Check if a token exists on the client side
    setTokenExists(checkTokenExists());
  }, []);
  
  const router = useRouter();

  const routeTo = (route) => {
    router.push(route);
  }

  const logout = () => {
    
  }

  return (
    <>
      <Head>
        <title>Home Page</title>
        <meta name="description" content="Home Page for Your App" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <main className={styles.main}>


        <div className={styles.logoContainer}>
        { !tokenExists ? (
            <>
              <h1>Want to hear a joke?</h1>
            </>
          ) : (
            <>
              <h1>You Suck</h1>
            </>
          )}
            
        </div>

        <div className={styles.buttonContainer} >
          { !tokenExists ? (
            <>
              <button onClick={() => routeTo('/login')} className={styles.button}>Login</button>
              <button onClick={() => routeTo('/register')} className={styles.button}>Register</button>
            </>
          ) : (
            <>
              <button onClick={() => routeTo('/profile')} className={styles.button}>View Profile</button>
            </>
          )}


          
        </div>
        <div className='container'>
        <div className="coffee-container">
    <div className="coffee-header">
      <div className="coffee-header__buttons coffee-header__button-one"></div>
      <div className="coffee-header__buttons coffee-header__button-two"></div>
      <div className="coffee-header__display"></div>
      <div className="coffee-header__details"></div>
    </div>
    <div className="coffee-medium">
      <div className="coffe-medium__exit"></div>
      <div className="coffee-medium__arm"></div>
      <div className="coffee-medium__liquid"></div>
      <div className="coffee-medium__smoke coffee-medium__smoke-one"></div>
      <div className="coffee-medium__smoke coffee-medium__smoke-two"></div>
      <div className="coffee-medium__smoke coffee-medium__smoke-three"></div>
      <div className="coffee-medium__smoke coffee-medium__smoke-for"></div>
      <div className="coffee-medium__cup"></div>
    </div>
    <div className="coffee-footer"></div>
        </div>
        </div>


      </main>
    </>
  )
}
