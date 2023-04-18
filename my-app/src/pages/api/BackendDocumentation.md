# Freel API Routes

This document outlines the routes for the Freel API, which provides functionality for managing photos, user accounts, and more.

## Endpoints

### Get Photo
- URL: `/api/photos/{photoId}`
- Method: `GET`
- Description: Retrieve a specific photo by its ID.

### Upload Photo
- URL: `/api/upload/{fireID}`
- Method: `POST`
- Description: Upload a photo associated with a user's Firebase ID.

### Create User
- URL: `/api/create/user`
- Method: `POST`
- Description: Create a new user account.

### Get User by Firebase ID
- URL: `/api/users/{fireID}/get`
- Method: `GET`
- Description: Retrieve a user's information by their Firebase ID.

### Get User by Username
- URL: `/api/username/{username}/get`
- Method: `GET`
- Description: Retrieve a user's information by their username.

### Get All Users
- URL: `/api/users/get`
- Method: `GET`
- Description: Retrieve a list of all users.

### Get Random Image
- URL: `/api/random_pic/get`
- Method: `GET`
- Description: Retrieve a random image from the collection.

### Update User Bio
- URL: `/api/users/{fireID}/update_bio`
- Method: `POST`
- Description: Update a user's bio using their Firebase ID.

### Upload Profile Picture
- URL: `/api/users/{fireID}/uploadProfilePicture`
- Method: `POST`
- Description: Upload a profile picture for a user using their Firebase ID.

### Get Profile Picture
- URL: `/api/users/{fireID}/getProfilePicture`
- Method: `GET`
- Description: Retrieve a user's profile picture by their Firebase ID.

### Get Nearby Users
- URL: `/api/nearby_users/{fireID}`
- Method: `GET`
- Description: Retrieve a list of nearby users based on a user's Firebase ID.

## Server

The server listens on port 8080 and uses CORS middleware to handle cross-origin requests.
