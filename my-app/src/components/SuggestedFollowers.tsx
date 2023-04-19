import React from 'react';
import TopBarSearch from './topbarsearch';

const SuggestedFollowers: React.FC = () => {

  

  const photographers = [
    { id: 1, name: 'John Doe', username: 'john_doe' },
    { id: 2, name: 'Jane Smith', username: 'jane_smith' },
    { id: 3, name: 'Charlie Brown', username: 'charlie_brown' },
    { id: 4, name: 'Oliver Twist', username: 'oliver_twist' },
    { id: 5, name: 'Emma Woodhouse', username: 'emma_woodhouse' },
  ];

  return (
    <div className="col-span-3 bg-gray-100 p-4 min-h-screen border-l border-gray-300">
      <h3 className="text-gray-800 font-bold mb-4">Suggested Followers</h3>
      <div className="space-y-4">
      <TopBarSearch></TopBarSearch>
        {photographers.map((photographer) => (
          <div key={photographer.id} className="flex items-center space-x-3">
            <img
              src={`https://source.unsplash.com/random/50x50?sig=${photographer.id}`}
              alt={photographer.name}
              className="w-12 h-12 rounded-full"
            />
            <div>
              <p className="text-gray-700 font-semibold">{photographer.name}</p>
              <p className="text-gray-500">@{photographer.username}</p>
            </div>
            <button className="ml-auto bg-blue-600 text-white px-4 py-1 rounded font-semibold">
              Follow
            </button>
           
          </div>
        ))}
      </div>
      
    </div>
  );
};

export default SuggestedFollowers;
