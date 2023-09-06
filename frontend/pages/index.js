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
        <div className='container-fluid navbar-container'>
        <nav className={`navbar navbar-expand-lg navbar-light ${styles.navbar}`}>
        <div className="container-fluid">
          <a className="navbar-brand" href="/">Your App</a>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav ml-auto">

              { !tokenExists ? (
                <>
                  <li className="nav-item">
                    <button onClick={() => routeTo('/login')} className={`btn btn-primary ${styles.button}`}>Login</button>
                  </li>
                  <li className="nav-item">
                    <button onClick={() => routeTo('/register')} className={`btn btn-primary ${styles.button}`}>Register</button>
                  </li>
                </>
              ) : (
                <li className="nav-item">
                  <button onClick={() => routeTo('/profile')} className={`btn btn-primary ${styles.button}`}>View Profile</button>
                </li>
              )}
            </ul>
          </div>
          <button className="navbar-brand float-right" onClick={logout}>Logout</button>
        </div>
      </nav>
        </div>
        <div className={styles.logoContainer}>
          <Image
            src="/next.svg"
            alt="Next.js Logo"
            width={180}
            height={37}
            priority
          />
        </div>

        <div className={styles.buttonContainer}>
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
      </main>
    </>
  )
}
