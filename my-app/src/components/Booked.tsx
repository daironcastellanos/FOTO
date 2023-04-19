
import React, { useState } from 'react';

import BookingModal from './BookingModal';


const ButtonComponent = () => {
  const [showModal, setShowModal] = useState(false);

  const handleClick = () => {
    setShowModal(!showModal);
  };

  return (
    <>
      <button 
      className="py-2 px-4 bg-gradient-to-br from-purple-700 to-pink-600 rounded-lg shadow-md text-white hover:shadow-lg transform hover:-translate-y-1 transition-all duration-200"
      onClick={handleClick}>Book a Photographer</button>
      
      {showModal && <BookingModal onClose={handleClick} />}
    </>
  );
};

export default ButtonComponent;
