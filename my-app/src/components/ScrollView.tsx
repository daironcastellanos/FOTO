import React, { useState, useEffect } from 'react';
import Image from 'next/image';

interface ImageData {
  id: string;
  url: string;
  caption: string;
}

const base64ToBlobUrl = async (base64Data: string, contentType = 'image/jpeg') => {

  const byteCharacters = atob(base64Data);
  const byteArrays = [];

  for (let offset = 0; offset < byteCharacters.length; offset += 512) {
    const slice = byteCharacters.slice(offset, offset + 512);
    const byteNumbers = new Array(slice.length);

    for (let i = 0; i < slice.length; i++) {
      byteNumbers[i] = slice.charCodeAt(i);
    }

    const byteArray = new Uint8Array(byteNumbers);
    byteArrays.push(byteArray);
  }

  const blob = new Blob(byteArrays, { type: contentType });
  const blobUrl = URL.createObjectURL(blob);
  return blobUrl;
}

const getAllUsers = async () => {
  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users/get`, {
      mode: "cors",
    });
    const data = await response.json();
    return data;
  } catch (error) {
    console.error("Error fetching all users:", error);
  }
};

const ScrollingView = () => {
  const [images, setImages] = useState<ImageData[]>([]);

  useEffect(() => {
    const fetchImages = async () => {
      const allUsers = await getAllUsers();
      const allPictures: ImageData[] = [];

      for (const user of allUsers) {
        const photos = user?.MyPhotos || [];

        for (const photoId of photos) {
          try {
            const response = await fetch(
              `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/photos/${photoId}/feed`
            );
            const jsonResponse = await response.json();

            const blobUrl = await base64ToBlobUrl(jsonResponse.data);

            allPictures.push({
              id: user.Username,
              url: blobUrl,
              caption: jsonResponse.caption || 'No caption',
            });
          } catch (error) {
            console.error(`Error fetching photo with ID ${photoId}:`, error);
          }
        }        
      }
      setImages(allPictures);
    };
    fetchImages();
  }, []);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      {images.map((image) => (
        <div key={image.id} className="w-full max-w-2xl mb-4">
          <Image src={image.url} alt={image.caption} width={800} height={600} />
          <div className="bg-gray-200 p-4">
            <p className="font-bold">{`${image.id}`}</p>
            <p>{image.caption}</p>
          </div>
        </div>
      ))}
      </div>
  )
};


export default ScrollingView

      