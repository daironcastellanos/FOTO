import React, { useState } from 'react';
import Link from 'next/link';

interface SearchResult {
  title: string;
  description: string;
  url: string;
}

const Search = () => {
  const [query, setQuery] = useState('');
  const [results, setResults] = useState<SearchResult[]>([]);

  const handleSearch = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    // TODO: Implement search logic here
    const data = await fetch(`https://api.example.com/search?q=${query}`);
    const results = await data.json();
    setResults(results);
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <Link href="/Home">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
        <div className="py-6">
          <form onSubmit={handleSearch} className="flex">
            <input
              type="text"
              value={query}
              onChange={(event) => setQuery(event.target.value)}
              placeholder="Search"
              className="flex-1 border-gray-300 focus:border-indigo-500 focus:ring-indigo-500 rounded-md py-2 px-4 block w-full appearance-none leading-5"
            />
            <button
              type="submit"
              className="ml-4 bg-indigo-500 hover:bg-indigo-700 text-white font-bold py-2 px-4 rounded"
            >
              Search
            </button>
          </form>
        </div>
        {results.length > 0 && (
          <div className="mt-6">
            <h1 className="text-3xl font-bold text-gray-900 mb-6">Search Results</h1>
            {results.map((result) => (
              <div key={result.url} className="bg-white shadow-sm rounded-md mb-6">
                <a href={result.url} target="_blank" rel="noopener noreferrer">
                  <div className="p-6">
                    <h2 className="text-xl font-bold text-gray-900">{result.title}</h2>
                    <p className="text-gray-700 mt-2">{result.description}</p>
                    <p className="text-blue-600 mt-2">{result.url}</p>
                  </div>
                </a>
              </div>
            ))}
          </div>
        )}
      </div>
    </div>
  );
};

export default Search;
