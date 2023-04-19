import React, { useState } from "react";

const TopBarSearch = () => {
  const [search, setSearch] = useState("");

  const handleSearchChange = (e) => {
    setSearch(e.target.value);
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    // Perform your search logic here
    console.log("Searching for:", search);
  };

  const handleClearSearch = () => {
    setSearch("");
  };

  return (
    <div className="p-4 flex justify-center">
      <form onSubmit={handleSearchSubmit} className="w-full max-w-md flex">
        <input
          type="text"
          value={search}
          onChange={handleSearchChange}
          placeholder="Search"
          className="w-full px-3 py-2 border-0 border-b border-gray-300 rounded-none bg-transparent focus:outline-none focus:border-gray-500"
        />
        {search && (
          <button
            type="button"
            onClick={handleClearSearch}
            className="text-gray-800 font-semibold py-2 px-3 focus:outline-none"
          >
            X
          </button>
        )}
      </form>
    </div>
  );
};

export default TopBarSearch;
