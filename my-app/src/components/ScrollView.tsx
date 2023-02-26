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

  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4">
      {images.map((image) => (
        <div key={image.id}>
          <Image src={image.url} alt={image.caption} width={800} height={200} />
          <p>{image.caption}</p>
        </div>
      ))}
    </div>
  );
};

export default ScrollingView;
