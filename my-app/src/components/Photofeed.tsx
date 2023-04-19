import React from 'react';
import Image from 'next/image';

const PhotoFeed = () => {
  return (
    <div className="grid grid-cols-2 gap-4">
      {/* Loop through your photos and render them in this grid */}
      <div className="bg-white p-4 rounded-lg shadow">
        <Image src="https://via.placeholder.com/200" alt="Placeholder" className="w-full h-48 object-cover mb-2 rounded" />
        <div className="space-y-1">
          <h3 className="font-semibold">Photographer Name</h3>
          <p className="text-sm text-gray-600">Location</p>
        </div>
      </div>
      {/* Add more photo elements like the one above */}
    </div>
  );
};

export default PhotoFeed;
