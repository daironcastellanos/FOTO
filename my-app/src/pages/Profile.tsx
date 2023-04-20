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


const getUserById = async () => {
  console.log("Trying to get user by ID");

  const user = await getUser();
  const fireID = user?.uid;

  console.log(fireID);

  var data = null;

  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users/${fireID}/get`
    );
    var data = await response.json();
    console.log(data);
  } catch (error) {
    console.error(`Error fetching user with ID ${fireID}:`, error);
  }

  return data;
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
  Following: string[];
  Followers: string[];
  profilePictureUrl: string;
  pictures: Picture[];
}

interface User {
  id: string;
  Bio: string;
  DOB: string;
  Email: string;
  FireID: string;
  Followers: string[];
  Following: string[];
  FullName: string;
  Location: {
    coordinates: number[];
    lat: number;
    lng: number;
  };
  MyPhotos: string[];
  ProfilePicture: string;
  SavedPhotos: string[];
  Username: string;
  _id: string;
}

const Profile: React.FC = () => {

  const router = useRouter();
  const { id } = router.query;
  const [usersPhotos, setUsersPhotos] = useState<Picture[]>([]);
  const [photoUrl, setPhotoUrl] = useState<string | string>("");
  const [user, setUser] = useState<User | null>(null);

  const getAllUsers = async () => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users/get`, {
        mode: "cors",
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error("Error fetching all users:", error);
    }
  };

  async function getProfilePicture() {
    const user = await getUser();
    const fireID = user?.uid;
    try {
      const response = await fetch(
        `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/users/${fireID}/getProfilePicture`
      );
      if (response.ok) {
        const blob = await response.blob();
        const imageUrl = URL.createObjectURL(blob);
        setPhotoUrl(imageUrl);
      } else {
        console.error("Error fetching profile picture:", response.statusText);
        return null;
      }
    } catch (error) {
      console.error("Error fetching profile picture:", error);
      return null;
    }
  }


useEffect(() => {
  console.log("useEffect called");

  const fetchData = async () => {
    await getAllUsers();
    await getProfilePicture();
    
    const user = await getUserById();
    setUser(user);
  };

  fetchData();
}, []);

  const [userProfile, setUserProfile] = React.useState<UserProfile>({
    id: '1',
    name: user?.FullName || 'Name',
    bio: user?.Bio || 'Bio',
    Following: user?.Following || [],
    Followers: user?.Followers || [],
    profilePictureUrl: '',
    pictures: [
      {
        id: '1',
        url: '',
      },
    ],
  });

  useEffect(() => {
    getUsrPhotoArray();
    console.log("updated user: ", user);
    console.log("updated user profile picture: ", user?.ProfilePicture);
    var newProfile = userProfile;
    newProfile.name = user?.FullName || 'Name';
    newProfile.bio = user?.Bio || 'Bio';
    setUserProfile(newProfile)
  }, [user]);

  const getUsrPhotoArray = async () => {
    console.log("Trying to get user photo array");
    const userdata = await getUserById();
    const usrObj = userdata;
    
    const photos = usrObj?.MyPhotos || [];
    const newPictures: Picture[] = [];
  
    for (const photoId of photos) {
      try {
        const response = await fetch(
          `${process.env.NEXT_PUBLIC_BACKEND_URL}/api/photos/${photoId}`
        );
        const blob = await response.blob();
        const objectUrl = URL.createObjectURL(blob);
  
        newPictures.push({
          id: photoId,
          url: objectUrl
        });
      } catch (error) {
        console.error(`Error fetching photo with ID ${photoId}:`, error);
      }
    }
  
    setUserProfile(prevProfile => ({
      ...prevProfile,
      pictures: newPictures
    }));
  }

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
        followers={userProfile.Followers.length} // Replace with the actual number of followers
        following={userProfile.Following.length} // Replace with the actual number of following users
      />
     <div className="grid grid-cols-3 gap-4 mt-4">
  {userProfile.pictures.map((picture) => (
    <div key={picture.id} className="relative overflow-hidden aspect-w-1 aspect-h-1">
      <Image
        src={picture.url}
        className="object-cover"
        alt="Posted picture"
        width={500}
        height={500}
      />
    </div>
  ))}
</div>
    </div>
  );
};

export default Profile;
