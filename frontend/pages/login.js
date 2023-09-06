
import '../styles/Register.module.css';
import { useState } from 'react';
import Cookies from 'js-cookie';
import { useRouter } from 'next/router';

function LoginPage() {
  const [credentials, setCredentials] = useState({ username: '', password: '' });
  const [token, setToken] = useState('');
  const router = useRouter();

  const addCookie = (token) => {
    try {
        Cookies.set('token', token, {
            expires: 7, // Set the cookie expiration in days
          });
      } catch (error) {
        console.error('Error setting cookie:', error);
      }
      
 }

 

  const handleChange = (e) => {
    const { name, value } = e.target;
    setCredentials({ ...credentials, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(credentials),
      });

      if (response.ok) {
        const data = await response.json();

        addCookie(data.token)
        router.push('/profile');
      } else {
        // Handle authentication error here
        console.error('Authentication failed');
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div>
      <h1>Login</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input type="text" name="username" value={credentials.username} onChange={handleChange} />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" name="password" value={credentials.password} onChange={handleChange} />
        </div>
        <div>
          <button type="submit">Login</button>
        </div>
      </form>

    </div>
  );
}

export default LoginPage;
