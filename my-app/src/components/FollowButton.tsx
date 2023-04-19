import React, { useState } from 'react';

interface FollowButtonProps {
  isFollowing: boolean;
  onFollow: () => void;
  onUnfollow: () => void;
}

const FollowButton: React.FC<FollowButtonProps> = ({ isFollowing, onFollow, onUnfollow }) => {
  const [following, setFollowing] = useState(isFollowing);

  const handleButtonClick = () => {
    if (following) {
      onUnfollow();
    } else {
      onFollow();
    }
    setFollowing(!following);
  };

  return (
    <div className="flex items-center justify-center">
      <button
        onClick={handleButtonClick}
        className={`${
          following ? 'bg-red-500' : 'bg-indigo-500'
        } py-2 px-4 text-white font-semibold rounded-md focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 hover:opacity-90`}
      >
        {following ? 'Unfollow' : 'Follow'}
      </button>
    </div>
  );
};

export default FollowButton;
