import React from 'react';
import Link from 'next/link';

const PrivacyPolicy = () => {
  return (
    <div className="min-h-screen bg-blue-100 py-6 flex flex-col justify-center sm:py-12">
     <div>
        <Link href="/">
          <button className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
            Back
          </button>
        </Link>
        </div>
     <div className="relative py-3 sm:max-w-xl sm:mx-auto">
        <div className="absolute inset-0 bg-gradient-to-r from-blue-200 to-purple-200 shadow-lg transform -skew-y-6 sm:skew-y-0 sm:-rotate-6 sm:rounded-3xl"></div>
        <div className="relative px-4 py-10 bg-white shadow-lg sm:rounded-3xl sm:p-20">
          <h1 className="text-4xl font-bold mb-4 text-purple-600">Privacy Policy</h1>
          <p className="text-lg mb-6">
            Your privacy is important to us. It is Freel's policy to respect your privacy regarding any information we may collect from you across our website, https://freel.example.com, and other sites we own and operate.
          </p>
          <h2 className="text-2xl font-semibold mb-4">Information we collect</h2>
          {/* Add your information collection policy */}
          <h2 className="text-2xl font-semibold mb-4">Log data</h2>
          {/* Add your log data policy */}
          <h2 className="text-2xl font-semibold mb-4">Cookies</h2>
          {/* Add your cookies policy */}
          <h2 className="text-2xl font-semibold mb-4">Third-party services</h2>
          {/* Add your third-party services policy */}
          <h2 className="text-2xl font-semibold mb-4">Security</h2>
          {/* Add your security policy */}
          <h2 className="text-2xl font-semibold mb-4">Links to other sites</h2>
          {/* Add your external links policy */}
          <h2 className="text-2xl font-semibold mb-4">Children's privacy</h2>
          {/* Add your children's privacy policy */}
          <h2 className="text-2xl font-semibold mb-4">Changes to this policy</h2>
          {/* Add your policy changes statement */}
          <h2 className="text-2xl font-semibold mb-4">Contact us</h2>
          <p className="text-lg">
            For any questions or concerns regarding our privacy policy, please contact us at support@freel.example.com.
          </p>
        </div>
      </div>
    </div>
  );
};

export default PrivacyPolicy;
