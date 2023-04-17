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

    try {
      const response = await fetch(
        `http://localhost:8080/api/users/${fireID}/get`
      );
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(`Error fetching user with ID ${fireID}:`, error);
    }
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
      <div className="mb-8">
        <h2 className="text-xl font-bold">Location Routes</h2>
        <button
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded m-4"
          onClick={getNearbyUsers}
        >
          Get Nearby Users
        </button>
      </div>

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
          <h2 className="text-xl font-bold">Uploading a photo</h2>
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

        <div className="mb-8">
          <h2 className="text-xl font-bold">Update Bio</h2>
          <div className="flex mb-4">
            <label htmlFor="bioInput" className="mr-2">
              New Bio:
            </label>
            <input
              type="text"
              id="bioInput"
              placeholder="New bio"
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
        </div>
      </div>
    </div>
  );
};
export default TestApi;
