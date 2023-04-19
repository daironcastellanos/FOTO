import axios from 'axios';

export async function CreateUserInMongo(userData: any) {
  
  try {
    console.log("Trying to create mongo user with Go");
    const response = await axios.post(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/create/user`,
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

