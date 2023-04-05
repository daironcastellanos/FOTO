import React, { useEffect } from 'react';
import Image from 'next/image';
import { useRouter } from 'next/router';

import UserStatistics from '@/components/userStats';

interface Picture {
  id: string;
  url: string;
}

interface Post {
  title: string;
  body: string;
  tags: string[];
  date: string;
  image: string;
}



interface MongoProfile {
  _id: string;
  name: string;
  bio: string;
  profilepicture: string;
  posts: Post[];
  location: Location;
  saved_post: Post[];
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
 
    id: '63f565f8df6db2c34aed8997',
    name: 'Gatico',
    bio: 'Bio',
    profilePictureUrl: 'https://placekitten.com/200/200',
    pictures: [
      {
        id: '63f565f8df6db2c34aed8997',
        url: 'https://placekitten.com/200/200',
      },
      {
        id: '2',
        url: 'https://placekitten.com/200/200',
      },
      {
        id: '3',
        url: 'https://placekitten.com/200/200',
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

  const [mongoProfile, setMongoProfile] = React.useState<MongoProfile | null>(null);

  React.useEffect(() => {
    //console.log("before")
    if (userProfile.id) {
      //console.log('after')
      fetch(`http://localhost:8080/api/users/${userProfile.id}/get`)
        .then((response) => response.json())
        .then((data) => setMongoProfile(data))
        .catch((error) => console.error('Error fetching user data:', error));
        
    }
  }, []);
  
  useEffect(() => {
    //console.log("mongo profile update ")
    console.log(mongoProfile)

  }, [mongoProfile]);


  const handleBackButtonClick = () => {
    router.back();
  };

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
        <button onClick={handleBackButtonClick}>
          <svg className="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M15 19l-7-7 7-7" />
          </svg>
        </button>
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
            <Image
              src={picture.url}
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
