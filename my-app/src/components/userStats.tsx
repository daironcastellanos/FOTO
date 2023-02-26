import React from 'react';

interface UserStatisticsProps {
  posts: number;
  followers: number;
  following: number;
}

const UserStatistics: React.FC<UserStatisticsProps> = ({ posts, followers, following }) => {
  return (
    <div>
      <p>Posts: {posts}</p>
      <p>Followers: {followers}</p>
      <p>Following: {following}</p>
    </div>
  );
};

export default UserStatistics;
