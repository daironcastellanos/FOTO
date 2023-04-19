import React from 'react';
import Link from 'next/link';

const TermsOfService = () => {
  return (
    <div className="min-h-screen bg-blue-100 py-6 flex flex-col justify-center sm:py-12">
      <div><Link href="/">
          <button className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
            Back
          </button>
        </Link></div>
      <div className="relative py-3 sm:max-w-xl sm:mx-auto">
        <div className="absolute inset-0 bg-gradient-to-r from-blue-200 to-purple-200 shadow-lg transform -skew-y-6 sm:skew-y-0 sm:-rotate-6 sm:rounded-3xl"></div>
        <div className="relative px-4 py-10 bg-white shadow-lg sm:rounded-3xl sm:p-20">
          <h1 className="text-4xl font-bold mb-4 text-purple-600">Terms of Service</h1>
          <p className="text-lg mb-6">
            By using the Freel website, you are agreeing to be bound by the following terms and conditions (Terms of Service).
          </p>
          {/* Add your Terms of Service content */}
          <h2 className="text-2xl font-semibold mb-4">Account Terms</h2>
          {/* Add your account terms */}
          <h2 className="text-2xl font-semibold mb-4">Payment and Refunds</h2>
          {/* Add your payment and refund policy */}
          <h2 className="text-2xl font-semibold mb-4">Cancellation and Termination</h2>
          {/* Add your cancellation and termination policy */}
          <h2 className="text-2xl font-semibold mb-4">Modifications to the Service</h2>
          {/* Add your service modification policy */}
          <h2 className="text-2xl font-semibold mb-4">Copyright and Content Ownership</h2>
          {/* Add your copyright and content ownership policy */}
          <h2 className="text-2xl font-semibold mb-4">General Conditions</h2>
          {/* Add your general conditions */}
          <h2 className="text-2xl font-semibold mb-4">Disclaimer of Warranties</h2>
          {/* Add your disclaimer of warranties */}
          <h2 className="text-2xl font-semibold mb-4">Limitation of Liability</h2>
          {/* Add your limitation of liability policy */}
          <h2 className="text-2xl font-semibold mb-4">Governing Law</h2>
          {/* Add your governing law policy */}
          <h2 className="text-2xl font-semibold mb-4">Changes to the Terms of Service</h2>
          {/* Add your changes to Terms of Service statement */}
          <h2 className="text-2xl font-semibold mb-4">Contact Us</h2>
          <p className="text-lg">
            For any questions or concerns regarding our Terms of Service, please contact us at support@freel.example.com.
          </p>
        </div>
      </div>
    </div>
  );
};

export default TermsOfService;
