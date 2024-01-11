import { useEffect, useState } from 'react';
import { parseCookies } from 'nookies';
import { removeCookie} from '@/helpers/auth'
import { useRouter } from 'next/router';

function ProfilePage() {
  const [user, setUser] = useState(null);
  const router = useRouter();



  useEffect(() => {
    // Check for the token cookie
    const cookies = parseCookies();
    const token = cookies.token;

    if (!token) {
      // Token is missing, redirect to the login page
      router.push('/login'); // Replace with your login page URL
      return;
    }

   fetchUserProfile(token)
      .then((userData) => {
        console.log(userData)
        setUser(userData);
      })
      .catch((error) => {
        console.error('Error fetching user profile:', error);
        // Handle errors, e.g., display an error message or redirect
      });

  }, [router]);

  const handleLogout = () => {
    removeCookie();
    router.push('/')
  }

  async function fetchUserProfile(token) {
    try {
      // Define the URL for your user profile API endpoint
      const apiUrl = 'http://localhost:8080/profile'; // Replace with your actual API URL
  
      // Make an authenticated request to the API using the token
      const response = await fetch(apiUrl, {
        method: 'GET',
        headers: {
          'Authorization': `Bearer ${token}`, // Add the token to the headers
        },
      });
  
      if (!response.ok) {
        throw new Error(`Error fetching user profile: ${response.statusText}`);
      }
  
      // Parse the JSON response and return the user data
      const userData = await response.json();
      return userData;
    } catch (error) {
      // Handle errors, e.g., display an error message or throw an error
      throw new Error(`Error fetching user profile: ${error.message}`);
    }
  }

  if (!user) {
    return <div>Loading...</div>;
  }

  return (
    <div className='container'>
      <div className='row'>
        <div className='col'>
          <h1>User Profile</h1>
          <p>Welcome! {user.username}</p>
          {/* Display user profile data here */}
          <button onClick={() => handleLogout()}>Logout</button>
        </div>
      </div>
    </div>
  );
}

export default ProfilePage;
