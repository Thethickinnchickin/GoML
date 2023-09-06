import { parseCookies, destroyCookie } from 'nookies';

// Function to check if a token exists in cookies
export function checkTokenExists() {
  const cookies = parseCookies();
  const token = cookies.token; // Replace 'token' with your actual cookie name

  return !!token; // Returns true if token exists, false otherwise
}


export function removeCookie () {
  try {
      destroyCookie(null, 'token');
    } catch (error) {
      console.error('Error setting cookie:', error);
    }
    
}

