import React from 'react';
import ButtonComponent from './Booked';
import FollowButton from './FollowButton';

interface UserStatisticsProps {
  posts: number;
  followers: number;
  following: number;
}

const UserStatistics: React.FC<UserStatisticsProps> = ({ posts, followers, following }) => {
  return (

    
    <div className="flex mt-4 space-x-4">
      <div className="text-center">
        <h3 className="text-2xl font-bold text-indigo-500">{posts}</h3>
        <p className="text-sm text-gray-600 tracking-wider uppercase">Posts</p>
      </div>
      <FollowButton></FollowButton>
      <ButtonComponent></ButtonComponent>
      <div className="text-center">
      
        <h3 className="text-2xl font-bold text-indigo-500">{followers}</h3>
        <p className="text-sm text-gray-600 tracking-wider uppercase">Followers</p>
      </div>
      <div className="text-center">
        <h3 className="text-2xl font-bold text-indigo-500">{following}</h3>
        <p className="text-sm text-gray-600 tracking-wider uppercase">Following</p>
        
      </div>
    </div>
   
  );
};

export default UserStatistics;
