import fetch from 'node-fetch'; // Import the fetch library

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

export default fetchUserProfile;
