import React from 'react';
import Image from 'next/image';

function Sidebar() {
  return (
    <div>
      <Image className="h-10 w-10" src="/images/freellogo.png" alt="Freel Logo" width={40} height={40} />
    </div>
  );
}

export default Sidebar;
