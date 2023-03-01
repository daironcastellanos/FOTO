import React, { useState, useEffect } from 'react';
import axios from 'axios';
import Image from 'next/image';

interface ImageData {
  id: number;
  url: string;
  caption: string;
}

const ScrollingView = () => {
  const [images, setImages] = useState<ImageData[]>([]);

  useEffect(() => {
    const fetchImages = async () => {
      const response = await axios.get('https://jsonplaceholder.typicode.com/photos');
      setImages(response.data);
    };

    fetchImages();
  }, []);

  const getRandomId = () => Math.floor(Math.random() * 1000);

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      {images.map((image) => (
        <div key={image.id} className="w-full max-w-2xl mb-4">
          <Image src={`https://picsum.photos/id/${getRandomId()}/800/600`} alt={image.caption} width={800} height={600} />
          <div className="bg-gray-200 p-4">
            <p className="font-bold">{`Username: ${image.id}`}</p>
            <p>{image.caption}</p>
            <p>{'This is the caption for the picture'}</p>
          </div>
        </div>
      ))}
    </div>
  );
};

export default ScrollingView;
