import React, { useState } from "react";
import { getAuth } from "firebase/auth";
import Image from "next/image";
import { auth } from "@/firebase/firebase";

const ImageDisplay: React.FC<{ src: string; alt: string }> = ({ src, alt }) => {
  return (
    <Image
      className="w-64 h-64 object-cover"
      src={src}
      alt={alt}
      width={256}
      height={256}
    />
  );
};

interface Post {
  title: string;
  body: string;
  tags: string[];
  date: string;
  image: string;
}

interface MongoProfile {
  _id: string;
  name: string;
  bio: string;
  profile_picture: string;
  posts: Post[];
  location: Location;
  saved_post: Post[];
}

const TestApi: React.FC = () => {
  const [id, setId] = useState("");
  const [bioid, setbioId] = useState("");

  const [photoId, setPhotoId] = useState("");
  const [bio, setBio] = useState("");
  const [imageFile, setImageFile] = useState<File | null>(null);
  const [nearbyId, setNearbyId] = useState("");
  const [uploadedFile, setUploadedFile] = useState<File | null>(null);
  const [savedphotoId, setsavedPhotoId] = useState("");
  const [followerId, setFollowerId] = useState("");

  const [randomPhotoUrl, setRandomPhotoUrl] = useState<string | null>(null);
  const [photoUrl, setPhotoUrl] = useState<string | null>(null);

  const handleIdChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setId(e.target.value);
  const handlePhotoIdChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setPhotoId(e.target.value);
  const handleBioInputChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setBio(e.target.value);
  const handleBioChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setbioId(e.target.value);
  const handleNearbyIdChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setNearbyId(e.target.value);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUploadedFile(e.target.files?.[0] || null);
  };

  const handleFollowerIdChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setFollowerId(e.target.value);

  const handlesavedPhotoIdChange = (e: React.ChangeEvent<HTMLInputElement>) =>
    setsavedPhotoId(e.target.value);

  const getUser = async () => {
    const user = auth.currentUser;
    console.log(user?.uid);
    console.log(user?.email);

    return user;
  };

  const getAllUsers = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/users/get", {
        mode: "cors",
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error fetching all users:", error);
    }
  };

  const getUserById = async () => {
    console.log("Trying to get user by ID");

    const user = await getUser();
    const fireID = user?.uid;

    console.log(fireID);

    var data = null;

    try {
      const response = await fetch(
        `http://localhost:8080/api/users/${fireID}/get`
      );
      var data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(`Error fetching user with ID ${fireID}:`, error);
    }

    return data;
  };

  const getNearbyUsers = async () => {
    const user = await getUser();
    const fireID = user?.uid;
    try {
      const response = await fetch(
        `http://localhost:8080/api/nearby_users/${fireID}`
      );
      if (
        response.ok &&
        response.headers.get("Content-Type")?.includes("application/json")
      ) {
        const data = await response.json();
        console.log(data);
      } else {
        console.error(
          "Error fetching nearby users:",
          response.status,
          response.statusText
        );
      }
    } catch (error) {
      console.error(
        `Error fetching nearby users for user with ID ${fireID}:`,
        error
      );
    }
  };  

  const getRandomPhoto = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/random_pic/get");
      const blob = await response.blob();
      const objectUrl = URL.createObjectURL(blob);
      console.log(objectUrl);
      setRandomPhotoUrl(objectUrl);
    } catch (error) {
      console.error("Error fetching random photo:", error);
    }
  };

  const getPhotoById = async () => {
    console.log("Trying to get photo by ID", photoId);
    try {
      const response = await fetch(
        `http://localhost:8080/api/photos/${photoId}`
      );
      const blob = await response.blob();
      const objectUrl = URL.createObjectURL(blob);
      setPhotoUrl(objectUrl);
    } catch (error) {
      console.error(`Error fetching photo with ID ${photoId}:`, error);
    }
  };


  async function getProfilePicture() {
    const user = await getUser();
    const fireID = user?.uid;
    try {
      const response = await fetch(`http://localhost:8080/api/users/${fireID}/getProfilePicture`);
      if (response.ok) {
        const blob = await response.blob();
        const imageUrl = URL.createObjectURL(blob);
        setPhotoUrl(imageUrl);
      } else {
        console.error("Error fetching profile picture:", response.statusText);
        return null;
      }
    } catch (error) {
      console.error("Error fetching profile picture:", error);
      return null;
    }
  }
  


  const uploadPhoto = async () => {
    if (!uploadedFile) {
      console.error("No file selected for upload");
      return;
    } else {
      const user = await getUser();
      const fireID = user?.uid;

      const formData = new FormData();
      formData.append("file", uploadedFile);

      try {
        const response = await fetch(
          `http://localhost:8080/api/upload/${fireID}`,
          {
            method: "POST",
            body: formData,
          }
        );
        if (
          response.ok &&
          response.headers.get("Content-Type")?.includes("application/json")
        ) {
          const data = await response.json();
          console.log(data);
        } else {
          console.error("Error uploading photo:", response.statusText);
        }
      } catch (error) {
        console.error("Error uploading photo:", error);
      }
    }
  };


  const uploadProfilePhoto = async () => {
    if (!uploadedFile) {
      console.error("No file selected for upload");
      return;
    } else {
      const user = await getUser();
      const fireID = user?.uid;
  
      const formData = new FormData();
      formData.append("file", uploadedFile);
  
      try {
        const response = await fetch(
          `http://localhost:8080/api/users/${fireID}/uploadProfilePicture`,
          {
            method: "POST",
            body: formData,
          }
        );
        if (
          response.ok &&
          response.headers.get("Content-Type")?.includes("application/json")
        ) {
          const data = await response.json();
          console.log(data);
        } else {
          console.error("Error uploading photo:", response.statusText);
        }
      } catch (error) {
        console.error("Error uploading photo:", error);
      }
    }
  };
  

  const updateBio = async () => {
    const user = await getUser();
    const fireID = user?.uid;
    console.log("bio", bio);
    try {
      const response = await fetch(
        `http://localhost:8080/api/users/${fireID}/update_bio`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ bio }),
        }
      );

      if (response.ok) {
        console.log("Bio updated successfully");
      } else {
        console.error("Error updating bio:", response.statusText);
      }
    } catch (error) {
      console.error("Error updating bio:", error);
    }
  };

  const search_username = async () => {
    console.log("Trying to get user by username");
    console.log("id", id);

    const username = id;

    try {
      const response = await fetch(
        `http://localhost:8080/api/username/${username}/get`
      );
      if (!response.ok) {
        throw new Error(await response.text());
      }
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(`Error fetching user with username ${username}:`, error);
    }
  };

// Add photo to user's SavedPhotos array
async function addPhotoToSaved() {
  const user = await getUser();
  const fireID = user?.uid;
  const photoId = savedphotoId;
  try {
    const response = await fetch(`http://localhost:8080/api/users/${fireID}/savePhoto`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(photoId)
    });

    if (response.ok) {
      console.log("Photo saved successfully.");
    } else {
      console.log("Error saving photo.");
    }
  } catch (error) {
    console.error("Error in addPhotoToSaved:", error);
  }
}

// Follow another user
async function followUser() {
  const user = await getUser();
  const fireID = user?.uid;
  const targetFireId = followerId;
  try {
    const response = await fetch(`http://localhost:8080/api/users/${fireID}/follow`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(targetFireId)
    });

    if (response.ok) {
      console.log("User followed successfully.");
    } else {
      console.log("Error following user.");
    }
  } catch (error) {
    console.error("Error in followUser:", error);
  }
}

// Remove photo from user's SavedPhotos array
async function removePhotoFromSaved() {
  const user = await getUser();
  const fireID = user?.uid;
  const photoId = savedphotoId;
  try {
    const response = await fetch(`http://localhost:8080/api/users/${fireID}/removePhoto`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(photoId)
    });

    if (response.ok) {
      console.log("Photo removed successfully.");
    } else {
      console.log("Error removing photo.");
    }
  } catch (error) {
    console.error("Error in removePhotoFromSaved:", error);
  }
}

// Unfollow another user
async function unfollowUser() {
  const user = await getUser();
  const fireID = user?.uid;
  const targetFireId = followerId;
  try {
    const response = await fetch(`http://localhost:8080/api/users/${fireID}/unfollow`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(targetFireId)
    });

    if (response.ok) {
      console.log("User unfollowed successfully.");
    } else {
      console.log("Error unfollowing user.");
    }
  } catch (error) {
    console.error("Error in unfollowUser:", error);
  }
}


// Update user's location
async function updateUserLocation() {
  const user = await getUser();
  const fireID = user?.uid;

  if (!navigator.geolocation) {
    console.error("Geolocation is not supported by your browser.");
    return;
  }

  navigator.geolocation.getCurrentPosition(async (position) => {
    const location = {
      lat: position.coords.latitude,
      lng: position.coords.longitude,
      coordinates: [position.coords.latitude, position.coords.longitude],
    };

    try {
      const response = await fetch(`http://localhost:8080/api/users/${fireID}/updateLocation`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(location),
      });

      if (response.ok) {
        console.log("User location updated successfully.");
      } else {
        console.log("Error updating user location.");
      }
    } catch (error) {
      console.error("Error in updateUserLocation:", error);
    }
  });
}



  return (
    <div className="container mx-auto p-4">
      <h1 className="text-3xl font-bold">Test API Routes</h1>
      <p className="text-lg">Click the buttons below to test the API routes.</p>

      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={getUserById}
      >
        Get Current
      </button>

      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={getAllUsers}
      >
        Get All Users
      </button>

      <div className="flex mb-4">
        <label htmlFor="search by username" className="mr-2"></label>
        <input
          type="text"
          id="photoId"
          placeholder="Username"
          onChange={handleIdChange}
          className="border border-gray-400 p-2"
        />
      </div>

      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={search_username}
      >
        Seach User
      </button>

      <div className="flex mb-4">
        <label htmlFor="new bio" className="mr-2">
          Update Bio:
        </label>
        <input
          type="text"
          id="photoId"
          placeholder="Photo ID"
          onChange={handleBioInputChange}
          className="border border-gray-400 p-2"
        />
      </div>

      <button
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
        onClick={updateBio}
      >
        Update Bio
      </button>

      {/* Render buttons for other routes */}
      <div className="mb-8">
        <h2 className="text-xl font-bold">Photo Routes</h2>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
          onClick={getRandomPhoto}
        >
          Get Random Photo
        </button>
        {randomPhotoUrl && (
          <ImageDisplay src={randomPhotoUrl} alt="Random Photo" />
        )}

        <div className="mb-8">
          <div className="flex mb-4">
            <label htmlFor="fileInput" className="mr-2">
              Select a photo:
            </label>
            <input
              type="file"
              id="fileInput"
              onChange={handleFileChange}
              className="border border-gray-400 p-2"
            />
          </div>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
            onClick={uploadPhoto}
          >
            Upload photo
          </button>
        </div>

        <div className="flex mb-4">
          <label htmlFor="photoId" className="mr-2">
            Photo ID:
          </label>
          <input
            type="text"
            id="photoId"
            placeholder="Photo ID"
            onChange={handlePhotoIdChange}
            className="border border-gray-400 p-2"
          />
        </div>

        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
          onClick={getPhotoById}
        >
          Get Photo
        </button>

        {photoUrl && <ImageDisplay src={photoUrl} alt="Photo By ID" />}

        {/* Add inputs for other routes if required */}


        <div className="mb-8">
          <div className="flex mb-4">
            <label htmlFor="fileInput" className="mr-2">
              Chose profile picture 
            </label>
            <input
              type="file"
              id="fileInput"
              onChange={handleFileChange}
              className="border border-gray-400 p-2"
            />
          </div>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
            onClick={uploadProfilePhoto}
          >
            Upload Profile Picture
          </button>
        </div>

        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
          onClick={getProfilePicture}
        >
          Get Profile Photo
        </button>

  

        <div className="mb-8">
          <h2 className="text-xl font-bold">Location Routes</h2>
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
            onClick={getNearbyUsers}
          >
            Get Nearby Users
          </button>
        </div>

              
<button
  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
  onClick={updateUserLocation}
>
  Update Users Location
</button>



        <div className="flex mb-4">
          <label htmlFor="photoId" className="mr-2">
            Photo ID to add to saved
          </label>
          <input
            type="text"
            id="photoId"
            placeholder="Saved Photo ID"
            onChange={handlesavedPhotoIdChange}
            className="border border-gray-400 p-2"
          />
        </div>
<button
  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
  onClick={addPhotoToSaved}
>
  Save Photo
</button>
<button
  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
  onClick={removePhotoFromSaved}
>
  Un Save Photo
</button>



<div className="flex mb-4">
          <label htmlFor="photoId" className="mr-2">
            Follower ID:
          </label>
          <input
            type="text"
            id="photoId"
            placeholder="Follower ID"
            onChange={handleFollowerIdChange}
            className="border border-gray-400 p-2"
          />
        </div>
<button
  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
  onClick={followUser}
>
  Follow User
</button>

<button
  className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
  onClick={unfollowUser}
>
  UnFollow User
</button>



      </div>
    </div>
  );
};
export default TestApi;
