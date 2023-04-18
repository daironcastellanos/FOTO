import React from 'react';
import Image from 'next/image';
import { useRouter } from 'next/router';
import Link from 'next/link';

import UserStatistics from '@/components/userStats';

interface Picture {
  id: string;
  url: string;
}

interface UserProfile {
  id: string;
  name: string;
  bio: string;
  profilePictureUrl: string;
  pictures: Picture[];
}

const Profile: React.FC = () => {
  const router = useRouter();
  const { id } = router.query;

  const [userProfile, setUserProfile] = React.useState<UserProfile>({
    id: '1',
    name: 'Gatico',
    bio: 'Bio',
    profilePictureUrl: 'https://placekitten.com/200/200',
    pictures: [
      {
        id: '1',
        url: 'https://placekitten.com/200/200',
      },
      {
        id: '2',
        url: 'https://placekitten.com/200/200',
      },
      {
        id: '3',
        url: 'https://placekitten.com/300/300?image=3',
      },
      {
        id: '4',
        url: 'https://placekitten.com/300/300?image=4',
      },
      {
        id: '5',
        url: 'https://placekitten.com/300/300?image=5',
      },
      {
        id: '6',
        url: 'https://placekitten.com/300/300?image=6',
      },
      {
        id: '7',
        url: 'https://placekitten.com/300/300?image=7',
      },
      {
        id: '8',
        url: 'https://placekitten.com/300/300?image=8',
      },
      {
        id: '9',
        url: 'https://placekitten.com/300/300?image=9',
      },
    ],
  });

  const handleBackButtonClick = () => {
    router.back();
  };

  const userPictures = userProfile.pictures.filter(
    (picture) => picture.id.startsWith(userProfile.id)
  );


  return (
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <div className="w-36 h-36 relative rounded-full overflow-hidden">
        <Image
          src={userProfile.profilePictureUrl}
          layout="fill"
          objectFit="cover"
          alt="Profile picture"
        />
      </div>
      <div className="absolute top-4 left-4">
      <Link href="/Home">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
      </div>
      <div className="text-center mt-2">
        <h2 className="text-2xl font-bold">{userProfile.name}</h2>
        <p className="text-sm text-gray-600">{userProfile.bio}</p>
      </div>
      <UserStatistics
        posts={userProfile.pictures.length}
        followers={100} // Replace with the actual number of followers
        following={100} // Replace with the actual number of following users
      />
     <div className="grid grid-cols-3 gap-4 mt-4">
  {userProfile.pictures.map((picture) => (
    <div key={picture.id} className="relative overflow-hidden aspect-w-1 aspect-h-1">
      <img
        src={picture.url}
        className="object-cover"
        alt="Posted picture"
      />
    </div>
  ))}
</div>
    </div>
  );
};

export default Profile;
