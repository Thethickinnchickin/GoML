import styles from '@/styles/Home.module.css'
import { useRouter } from 'next/router'
import { checkTokenExists } from '@/helpers/auth'
import { useState, useEffect } from 'react'


const Navbar = () => {
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
        <nav className={`navbar navbar-expand-lg navbar-light ${styles.navbar}`}>
        <div className="container-fluid">
          <a className="navbar-brand" href="/">
            <img src="logo.svg" alt="Logo" width="100" height="50"/>
          </a>
          <button
            className="navbar-toggler"
            type="button"
            data-bs-toggle="collapse"
            data-bs-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >

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
          { tokenExists ? (
            <button className="navbar-brand float-right" onClick={logout}>Logout</button>
          ): (
            <></>
          )}
        </div>
      </nav>
    )
}


export default Navbar;