import React, { useState } from 'react';
import Link from 'next/link';

interface UploadPageProps {
  onSubmit: (imageUrl: string) => void;
}

const UploadPage: React.FC<UploadPageProps> = ({ onSubmit }) => {
  const [imageUrl, setImageUrl] = useState('');

  const handleFileInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onload = () => {
        const result = reader.result;
        if (typeof result === 'string') {
          setImageUrl(result);
        }
      };
      reader.readAsDataURL(file);
    }
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    onSubmit(imageUrl);
  };

  return (
    <div className="max-w-xl mx-auto p-4">
         <Link href="/Home">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
      <h1 className="text-3xl font-bold mb-4">Upload an Image</h1>
      <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md p-4">
        <div className="mb-4">
          <label htmlFor="file-input" className="block font-medium mb-1">
            Choose a file:
          </label>
          <input type="file" id="file-input" onChange={handleFileInputChange} className="border rounded-lg p-2" />
        </div>
        {imageUrl && (
          <div className="mb-4">
            <img src={imageUrl} alt="Preview" className="max-w-full h-auto" />
          </div>
        )}
        <div className="text-center">
          <button
            type="submit"
            disabled={!imageUrl}
            className={`bg-blue-500 text-white rounded-lg font-medium py-2 px-4 ${
              !imageUrl ? 'opacity-50 cursor-not-allowed' : ''
            }`}
          >
            Upload
          </button>
        </div>
      </form>
    </div>
  );
};

export default UploadPage;
