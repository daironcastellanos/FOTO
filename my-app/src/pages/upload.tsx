import React, { useState, useEffect } from 'react';
import Link from 'next/link';
import { auth } from '@/firebase/firebase';
import { useRouter } from 'next/router';

const UploadPage: React.FC = () => {
  const [uploadedFile, setUploadedFile] = useState<File | null>(null);
  const [uploadSuccess, setUploadSuccess] = useState(false);
  const router = useRouter();

  const getUser = async () => {
    const user = auth.currentUser;
    console.log(user?.uid);
    console.log(user?.email);

    return user;
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUploadedFile(e.target.files?.[0] || null);
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
          setUploadSuccess(true);
        } else {
          console.error("Error uploading photo:", response.statusText);
        }
      } catch (error) {
        console.error("Error uploading photo:", error);
      }
    }
  };

  useEffect(() => {
    if (uploadSuccess) {
      setTimeout(() => {
        router.push('/Home');
      }, 2000);
    }
  }, [uploadSuccess, router]);

  return (
    <div className="max-w-xl mx-auto p-4">
      <Link href="/Home">
        <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
          Back
        </h1>
      </Link>
      <h1 className="text-3xl font-bold mb-4">Upload an Image</h1>
      <div className="bg-white rounded-lg shadow-md p-4">
        <div className="mb-4">
          <label htmlFor="fileInput" className="block font-medium mb-1">
            Choose a file:
          </label>
          <input
            type="file"
            id="fileInput"
            onChange={handleFileChange}
            className="border rounded-lg p-2"
          />
        </div>
        {uploadSuccess && (
          <div className="mb-4 text-green-500 font-semibold">
            Image uploaded successfully!
          </div>
        )}
        <div className="text-center">
          <button
            className="bg-blue-500 text-white rounded-lg font-medium py-2 px-4"
            onClick={uploadPhoto}
          >
            Upload
          </button>
        </div>
      </div>
    </div>
  );
};

export default UploadPage;
