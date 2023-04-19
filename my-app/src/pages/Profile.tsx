import React, {useState, useEffect} from 'react';
import Image from 'next/image';
import { useRouter } from 'next/router';
import Link from 'next/link';
import { auth } from "@/firebase/firebase";
import UserStatistics from '@/components/userStats';

const ImageDisplay: React.FC<{ src: string; alt: string }> = ({ src, alt }) => {
  return (
    <Image
      className="w-64 h-64 object-cover"
      src={src}
      alt={alt}
      width={256}
      height={256}
    />
  );
};


const getUser = async () => {
  const user = await auth.currentUser;
  console.log(user?.uid);
  console.log(user?.email);

  return user;
};

const getUid = async () => {
  const user = await getUser();
  const fireID = await user?.uid;

  return fireID;}

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
  const [photoUrl, setPhotoUrl] = useState<string | string>("");

const getAllUsers = async () => {
    try {
      const response = await fetch("http://localhost:8080/api/users/get", {
        mode: "cors",
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error fetching all users:", error);
    }
};

  
const getProfilePicture = async () => {
  const fireID = await getUid();
  console.log("firebase id",fireID)
  try {
    const response = await fetch(`http://localhost:8080/api/users/${fireID}/getProfilePicture`);
    
      const blob = await response.blob();
      const imageUrl = URL.createObjectURL(blob);
      setPhotoUrl(imageUrl);
  
      console.error("Error fetching profile picture:", response.statusText);
      return null;
    
  } catch (error) {

    console.error("Error fetching profile picture:", error);
  }
}

  useEffect(() => {
    console.log("useEffect called")
    getAllUsers();
    getProfilePicture();
  }, []);
  

  const [userProfile, setUserProfile] = React.useState<UserProfile>({
    id: '1',
    name: 'Gatico',
    bio: 'Bio',
    profilePictureUrl: photoUrl,
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
        <ImageDisplay src={photoUrl} alt="Profile picture"
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
        src={photoUrl}
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
