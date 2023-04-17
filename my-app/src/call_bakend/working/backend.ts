import axios from 'axios';

export async function CreateUserInMongo(userData: any) {
  
  try {
    console.log("Trying to create mongo user with Go");
    const response = await axios.post(
      `http://localhost:8080/api/create/user`,
      JSON.stringify(userData), // Convert userData to JSON string
      {
        headers: {
          "Content-Type": "application/json",
        },
      }
    );
    console.log("User created successfully:", response);
    return true;
  } catch (error) {
    console.error("Error creating user:", error);
    return false;
  }
}