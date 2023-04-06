import React, { useState } from 'react';
import styles from './TestApi.module.css';

const ImageDisplay: React.FC<{ src: string; alt: string }> = ({ src, alt }) => {
    return <img className={styles.imageDisplay} src={src} alt={alt} />;
};

interface Post {
    title: string;
    body: string;
    tags: string[];
    date: string;
    image: string;
}

interface MongoProfile {
    _id: string;
    name: string;
    bio: string;
    profilepicture: string;
    posts: Post[];
    location: Location;
    saved_post: Post[];
}

const TestApi: React.FC = () => {
    const [id, setId] = useState('');
    const [bioid, setbioId] = useState('');
    const [photoId, setPhotoId] = useState('');
    const [bio, setBio] = useState('');
    const [imageFile, setImageFile] = useState<File | null>(null);
    const [nearbyId, setNearbyId] = useState('');
    const [uploadedFile, setUploadedFile] = useState<File | null>(null);

    const [randomPhotoUrl, setRandomPhotoUrl] = useState<string | null>(null);
    const [photoUrl, setPhotoUrl] = useState<string | null>(null);

    const handleIdChange = (e: React.ChangeEvent<HTMLInputElement>) => setId(e.target.value);
    const handlePhotoIdChange = (e: React.ChangeEvent<HTMLInputElement>) => setPhotoId(e.target.value);
    const handlebioIdChange = (e: React.ChangeEvent<HTMLInputElement>) => setBio(e.target.value);
    const handleBioChange = (e: React.ChangeEvent<HTMLInputElement>) => setbioId(e.target.value);
    const handleNearbyIdChange = (e: React.ChangeEvent<HTMLInputElement>) => setNearbyId(e.target.value);

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setUploadedFile(e.target.files?.[0] || null);
      };
      
    const getAllUsers = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/users/get');
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error('Error fetching all users:', error);
        }
    };

    const getUserById = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/users/${id}/get`);
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error(`Error fetching user with ID ${id}:`, error);
        }
    };

    const getNearbyUsers = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/nearby_users/${nearbyId}`);
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error(`Error fetching nearby users for user with ID ${nearbyId}:`, error);
        }
    };

    const getRandomPhoto = async () => {
        try {
            const response = await fetch('http://localhost:8080/api/random_pic/get');
            const blob = await response.blob();
            const objectUrl = URL.createObjectURL(blob);
            setRandomPhotoUrl(objectUrl);
        } catch (error) {
            console.error('Error fetching random photo:', error);
        }
    };

    const getPhotoById = async () => {
        try {
            const response = await fetch(`http://localhost:8080/api/photos/${photoId}`);
            const blob = await response.blob();
            const objectUrl = URL.createObjectURL(blob);
            setPhotoUrl(objectUrl);
        } catch (error) {
            console.error(`Error fetching photo with ID ${id}:`, error);
        }
    };

    const updateBio = async () => {
        try {
            const response = await fetch(`/api/users/${bioid}/${bio}`, {
                method: 'POST',
            });

            if (response.ok) {
                console.log('Bio updated successfully');
            } else {
                console.error('Error updating bio:', response.statusText);
            }
        } catch (error) {
            console.error('Error updating bio:', error);
        }
    };

    const uploadPhoto = async () => {
        if (!uploadedFile) {
            console.error('No file selected for upload');
            return;
        }

        const formData = new FormData();
        formData.append('file', uploadedFile);

        try {
            const response = await fetch('http://localhost:8080/api/upload', {
                method: 'POST',
                body: formData,
            });
            const data = await response.json();
            console.log(data);
        } catch (error) {
            console.error('Error uploading photo:', error);
        }
    };





    return (
        <div className={styles.container}>
            <h1>Test API Routes</h1>
            <p>Click the buttons below to test the API routes.</p>
            <p></p>
            <p></p>

            <button className={styles.apiButton} onClick={getAllUsers}>
                Get All Users
            </button>

            <div className={styles.section}>
                <h2 className={styles.sectionTitle}>User Routes</h2>
                <div className={styles.inputGroup}>
                    <label htmlFor="userId">User ID:</label>
                    <input type="text" id="userId" placeholder="User ID" onChange={handleIdChange} />
                </div>
            </div>
            <p></p>

            <button className={styles.apiButton} onClick={getUserById}>
                Get User By ID
            </button>

            <div className={styles.section}>
                <h2 className={styles.sectionTitle}>Location Routes</h2>
                <div className={styles.inputGroup}>
                    <label htmlFor="nearbyId">User ID:</label>
                    <input type="text" id="nearbyId" placeholder="User ID" onChange={handleNearbyIdChange} />
                </div>

                <button className={styles.apiButton} onClick={getNearbyUsers}>
                    Get Nearby Users
                </button>
            </div>
            <p></p>
            <p></p>

            {/* Render buttons for other routes */}
            <div className={styles.section}>
                <h2 className={styles.sectionTitle}>Photo Routes</h2>
                <button className={styles.apiButton} onClick={getRandomPhoto}>
                    Get Random Photo
                </button>
                {randomPhotoUrl && <ImageDisplay src={randomPhotoUrl} alt="Random Photo" />}
                <div className={styles.inputGroup}>
                    <label htmlFor="photoId">Photo ID:</label>
                    <input type="text" id="photoId" placeholder="Photo ID" onChange={handlePhotoIdChange} />
                </div>
                <button className={styles.apiButton} onClick={getPhotoById}>
                    Get Photo By ID
                </button>

                {photoUrl && <ImageDisplay src={photoUrl} alt="Photo By ID" />}
                <p></p>
                <p></p>

                {/* Add inputs for other routes if required */}

                <div className={styles.section}>
                    <h2 className={styles.sectionTitle}>Update Routes working progress...</h2>
                    <label htmlFor="userId">User ID:</label>
                    <input type="text" id="userId" placeholder="User ID" onChange={handlebioIdChange} />
                </div>
                <div className={styles.inputGroup}>
                    <label htmlFor="bio">Bio:</label>
                    <input type="text" id="bio" placeholder="Bio" onChange={handleBioChange} />
                </div>
                <button className={styles.apiButton} onClick={updateBio}>
                    Update Bio
                </button>
            </div>
            <p></p>
            <p></p>

            <div className={styles.section}>
                <h2 className={styles.sectionTitle}>Uploading a photo</h2>
                <div className={styles.inputGroup}>
                    <label htmlFor="fileInput">Select a photo:</label>
                    <input type="file" id="fileInput" onChange={handleFileChange} />
                </div>
                <button className={styles.apiButton} onClick={uploadPhoto}>
                    Upload photo
                </button>
            </div>

        </div>
    );  
};

export default TestApi;