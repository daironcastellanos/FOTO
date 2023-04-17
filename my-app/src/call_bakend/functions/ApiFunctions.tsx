const baseUrl: string = 'http://localhost:8080/api/';

const ApiFunctions = () => {
  /**
   * Fetches all users from the API
   */
  const getAllUsers = async () => {
    try {
      const response = await fetch(`${baseUrl}users/get`);
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error('Error fetching all users:', error);
    }
  };

  /**
   * Fetches user data by ID from the API
   * @param id - The user ID to fetch
   */
  const getUserById = async (id: string) => {
    try {
      const response = await fetch(`${baseUrl}users/${id}/get`);
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(`Error fetching user with ID ${id}:`, error);
    }
  };

  /**
   * Fetches nearby users by ID from the API
   * @param nearbyId - The ID of the user to fetch nearby users for
   */
  const getNearbyUsers = async (nearbyId: string) => {
    try {
      const response = await fetch(`${baseUrl}nearby_users/${nearbyId}`);
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(`Error fetching nearby users for user with ID ${nearbyId}:`, error);
    }
  };

  /**
   * Fetches a random photo from the API
   * @returns The URL of the photo as a string
   */
  const getRandomPhoto = async () => {
    try {
      const response = await fetch(`${baseUrl}random_pic/get`);
      const blob = await response.blob();
      const objectUrl = URL.createObjectURL(blob);
      return objectUrl;
    } catch (error) {
      console.error('Error fetching random photo:', error);
    }
  };

  /**
   * Fetches a photo by ID from the API
   * @param photoId - The ID of the photo to fetch
   * @returns The URL of the photo as a string
   */
  const getPhotoById = async (photoId: string) => {
    try {
      const response = await fetch(`${baseUrl}photos/${photoId}`);
      const blob = await response.blob();
      const objectUrl = URL.createObjectURL(blob);
      return objectUrl;
    } catch (error) {
      console.error(`Error fetching photo with ID ${photoId}:`, error);
    }
  };

  /**
   * Updates a user's bio in the API
   * @param bioId - The ID of the user to update the bio for
   * @param bio - The new bio to set for the user
   */
const updateBio = async (bioId: string, bio: string) => {
    try {
        const response = await fetch(`http://localhost:8080/api/users/${bioId}/${bio}`, {
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

const uploadPhoto = async (file: File) => {
    if (!file) {
        console.error('No file selected for upload');
        return;
    }

    const formData = new FormData();
    formData.append('file', file);

    try {
        const response = await fetch('http://localhost:8080/api/upload', {
            method: 'POST',
            body: formData,
        });
        const data = await response.json();
        console.log(data);
    }

    catch (error) {
        console.error('Error uploading file:', error);
    }
};

}
export default ApiFunctions;

