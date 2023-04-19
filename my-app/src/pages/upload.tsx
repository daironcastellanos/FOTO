import React, { useState } from "react";
import Link from "next/link";
import Image from "next/image";
import { auth } from "@/firebase/firebase";

const getUser = async () => {
  const user = auth.currentUser;
  console.log(user?.uid);
  console.log(user?.email);

  return user;
};

interface UploadPageProps {
  onSubmit: (imageUrl: string) => void;
}

const UploadPage: React.FC<UploadPageProps> = ({ onSubmit }) => {
  const [uploadedFile, setUploadedFile] = useState<File | null>(null);
  const [imageUrl, setImageUrl] = useState("");

  const handleFileInputChange = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    const file = event.target.files?.[0];
    if (file) {
      setUploadedFile(file);
      const reader = new FileReader();
      reader.onload = () => {
        const result = reader.result;
        if (typeof result === "string") {
          setImageUrl(result);
        }
      };
      reader.readAsDataURL(file);
    }
  };

  const handlePhotoUpload = async (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
    if (uploadedFile) {
      await uploadPhoto();
    }
  };

  const handleProfilePhotoUpload = async (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    event.preventDefault();
    if (uploadedFile) {
      await uploadProfilePhoto();
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
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/upload/${fireID}`,
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
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users/${fireID}/uploadProfilePicture`,
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

  return (
    <div className="max-w-xl mx-auto p-4">
      <Link href="/Home">
        <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
          Back
        </h1>
      </Link>
      <h1 className="text-3xl font-bold mb-4">Upload an Image</h1>
      <form className="bg-white rounded-lg shadow-md p-4">
        <div className="mb-4">
          <label htmlFor="file-input" className="block font-medium mb-1">
            Choose a file:
          </label>
          <input
            type="file"
            id="file-input"
            onChange={handleFileInputChange}
            className="border rounded-lg p-2"
          />
        </div>
        {imageUrl && (
          <div className="mb-4">
            <Image
              src={imageUrl}
              alt="Preview"
              className="max-w-full h-auto"
              width={640}
              height={480}
              layout="responsive"
              objectFit="contain"
            />
          </div>
        )}
        <div className="text-center mb-4">
          <button
            onClick={handlePhotoUpload}
            disabled={!imageUrl}
            className={imageUrl ? "btn" : "btn-disabled"}
          >
            Upload
          </button>
        </div>
        <div className="text-center">
          <button
            onClick={handleProfilePhotoUpload}
            disabled={!imageUrl}
            className={imageUrl ? "btn" : "btn-disabled"}
          >
            Upload Profile Picture
          </button>
        </div>
      </form>
    </div>
  );
};

export default UploadPage;
