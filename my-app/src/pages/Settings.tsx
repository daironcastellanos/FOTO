import React, { useState } from 'react';
import Link from 'next/link';

interface SettingsData {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
  notifications: boolean;
}

const Settings = () => {
  const [settings, setSettings] = useState<SettingsData>({
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
    notifications: true,
  });

  const handleUsernameChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSettings((prevSettings) => ({
      ...prevSettings,
      username: event.target.value,
    }));
  };

  const handleEmailChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSettings((prevSettings) => ({
      ...prevSettings,
      email: event.target.value,
    }));
  };

  const handlePasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSettings((prevSettings) => ({
      ...prevSettings,
      password: event.target.value,
    }));
  };

  const handleConfirmPasswordChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSettings((prevSettings) => ({
      ...prevSettings,
      confirmPassword: event.target.value,
    }));
  };

  const handleNotificationChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    setSettings((prevSettings) => ({
      ...prevSettings,
      notifications: event.target.checked,
    }));
  };

  const handlePasswordReset = () => {
    console.log('Password reset');
  };

  const handleSaveChanges = () => {
    console.log('Settings saved:', settings);
  };

  const handleLogOut = () => {
    console.log('Logged out');
  };

  const handleSubmit = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    handleSaveChanges();
  };

  return (
    <div className="min-h-screen bg-gray-100">
      <Link href="/Home">
            <h1 className="absolute top-3 left-3 bg-green-500 text-white py-2 px-4 rounded-lg hover:bg-green-600">
              Back
            </h1>
          </Link>
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="max-w-md mx-auto bg-white shadow-sm rounded-md overflow-hidden">
          <div className="px-4 py-5 sm:p-6">
            <h1 className="text-3xl font-bold mb-6">Settings</h1>
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label htmlFor="username" className="block text-gray-700 font-bold mb-2">
                  Username
                </label>
                <input
                  type="text"
                  id="username"
                  value={settings.username}
                  onChange={handleUsernameChange}
                  className="border-2 border-gray-300 p-2 w-full rounded-md"
                />
              </div>
              <div className="mb-4">
                <label htmlFor="email" className="block text-gray-700 font-bold mb-2">
                  Email
                </label>
                <input
                  type="email"
                  id="email"
                  value={settings.email}
                  onChange={handleEmailChange}
                  className="border-2 border-gray-300 p-2 w-full rounded-md"
                />
              </div>
              <div className="mb-4">
                <label htmlFor="password" className="block text-gray-700 font-bold mb-2">
                  New Password
                </label>
                <input
                  type="password"
                  id="password"
                  value={settings.password}
                  onChange={handlePasswordChange}
                  className="border-2 border-gray-300 p-2 w-full rounded-md"
                />
              </div>
              <div className="mb-4">
            <label htmlFor="confirmPassword" className="block text-gray-700 font-bold mb-2">
              Confirm Password
            </label>
            <input
              type="password"
              id="confirmPassword"
              value={settings.confirmPassword}
              onChange={handleConfirmPasswordChange}
              className="border-2 border-gray-300 p-2 w-full rounded-md"
            />
          </div>
          <div className="mb-4">
            <div className="flex items-center">
              <input
                type="checkbox"
                id="notifications"
                checked={settings.notifications}
                onChange={handleNotificationChange}
                className="mr-2"
              />
              <label htmlFor="notifications" className="text-gray-700 font-bold">
                Receive notifications
              </label>
            </div>
          </div>
          <div className="mb-4 flex justify-between items-center">
            <button
              type="button"
              onClick={handlePasswordReset}
              className="text-sm text-gray-500 hover:text-gray-700"
            >
              Reset password
            </button>
            <button type="submit" className="px-4 py-2 bg-blue-500 text-white rounded-md">
              Save Changes
            </button>
          </div>
        </form>
      </div>
      <div className="px-4 py-3 bg-gray-50 text-right sm:px-6">
        <button type="button" onClick={handleLogOut} className="text-sm text-gray-500 hover:text-gray-700">
          Log out
        </button>
      </div>
    </div>
  </div>
</div>
);
};

export default Settings;