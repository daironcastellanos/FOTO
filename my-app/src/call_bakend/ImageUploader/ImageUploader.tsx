// src/components/ImageUploader.tsx
import React, { ChangeEvent, useState } from 'react';

interface ImageUploaderProps {
  onUploadSuccess: (response: Response) => void;
  onUploadError: (error: Error) => void;
}

const ImageUploader: React.FC<ImageUploaderProps> = ({ onUploadSuccess, onUploadError }) => {
  const [file, setFile] = useState<File | null>(null);

  const handleFileChange = (event: ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files.length > 0) {
      setFile(event.target.files[0]);
    }
  };

  const handleSubmit = async () => {
    if (file) {
      try {
        const formData = new FormData();
        formData.append('file', file);

        const response = await fetch('http://localhost:8080/upload', {
          method: 'POST',
          body: formData,
        });

        if (!response.ok) {
          throw new Error('Failed to upload image');
        }

        onUploadSuccess(await response.json());
      } catch (error) {
        onUploadError(error);
      }
    }
  };

  return (
    <div>
      <input type="file" onChange={handleFileChange} />
      <button onClick={handleSubmit}>Upload</button>
    </div>
  );
};

export default ImageUploader;
