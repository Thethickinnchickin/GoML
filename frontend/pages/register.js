// pages/register.js
import { useState } from 'react';
import '../styles/Register.module.css';
import { useRouter } from 'next/router';
import  Cookies from 'js-cookie'

export default function Register() {
  const [formData, setFormData] = useState({
    username: '',
    password: '',
    confirmPassword: '',
  });

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
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (formData.password !== formData.confirmPassword) {
      alert('Passwords do not match');
      return;
    }

    const response = await fetch('http://localhost:8080/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: formData.username,
        password: formData.password,
      }),
    });

    if (response.ok) {
        const data = await response.json();
        console.log(data)
        addCookie(data.token)
        router.push('/profile');
      
    } else {
      alert('Registration failed');
    }
  };

  return (
    <div className='form container'>
      <div className='col-4'>
      <h1>Register</h1>
      <form  onSubmit={handleSubmit}>
        <div>
          <label>Username:</label>
          <input
            type="text"
            name="username"
            value={formData.username}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Password:</label>
          <input
            type="password"
            name="password"
            value={formData.password}
            onChange={handleChange}
            required
          />
        </div>
        <div>
          <label>Confirm Password:</label>
          <input
            type="password"
            name="confirmPassword"
            value={formData.confirmPassword}
            onChange={handleChange}
            required
          />
        </div>
        <div className='mt-4'>
          <button type="submit">Register</button>
        </div>
      </form>
      </div>

    </div>
  );
}
