import React from 'react';
import Image from 'next/image';
import { useRouter } from 'next/router';

import UserStatistics from '@/components/userStats';

const Profile = () => {
  const router = useRouter();
  const { id } = router.query;

  // replace the following URL with your profile picture URL
  const profilePictureUrl = 'https://placekitten.com/200/200';

  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <div className="w-36 h-36 relative rounded-full overflow-hidden">
        <Image
          src={profilePictureUrl}
          layout="fill"
          objectFit="cover"
          alt="Profile picture"
        />
      </div>
      <div className="text-center mt-2">
        <h2 className="text-2xl font-bold">Gatico</h2>
        <p className="text-sm text-gray-600">Bio</p>
      </div>
      <UserStatistics posts={10} followers={100} following={100} />
      <div className="grid grid-cols-3 gap-4 mt-4">
        {Array.from({ length: 9 }, (_, i) => (
          <div key={i} className="relative overflow-hidden aspect-w-1 aspect-h-1">
            <Image
              src={`https://placekitten.com/300/300?image=${i}`}
              layout="fill"
              objectFit="cover"
              alt="Posted picture"
            />
          </div>
        ))}
      </div>
    </div>
  );
};

export default Profile;
