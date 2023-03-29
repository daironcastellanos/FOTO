import React, { useState } from 'react';
import Link from 'next/link';

interface SettingsData {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
  notifications: boolean;
}

const SettingsPage = () => {
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
    <div className="flex flex-col items-center justify-center min-h-screen py-2">
      <div className="flex w-full justify-between items-center mb-4">
        <div>
          <Link href="/screens/HomePage">
            Back
          </Link>
        </div>
        <Link href="/">
        <button
          onClick={handleLogOut}
          className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded"
        >
          Log Out
        </button>
        </Link>
      </div>
      <h1 className="text-2xl font-bold mb-4">Settings</h1>
      <form onSubmit={handleSubmit} className="w-full max-w-md">
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
          <label htmlFor="confirm-password" className="block text-gray-700 font-bold mb-2">
            Confirm Password
          </label>
          <input
            type="password"
            id="confirm-password"
            value={settings.confirmPassword}
            onChange={handleConfirmPasswordChange}
            className="border-2 border-gray-300 p-2 w-full rounded-md"
          />
        </div>
        <div className="mb-4">
          <label htmlFor="notifications" className="block text-gray-700 font-bold mb-2">
            Notifications
          </label>
          <input
            type="checkbox"
            id="notifications"
            checked={settings.notifications}
            onChange={handleNotificationChange}
          />
        </div>
        <button
          type="submit"
          className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
        >
          Save Changes
        </button>
        <button
          type="button"
          onClick={handlePasswordReset}
          className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded ml-4"
        >
          Reset Password
        </button>
      </form>
    </div>
  );
};
export default SettingsPage;

