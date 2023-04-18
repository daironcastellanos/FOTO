# functions in TestApi.js

## 1. ImageDisplay
A functional component that receives `src` and `alt` as props and renders an image using Next.js `Image` component.

## 2. TestApi
A functional component that provides an interface to test various API routes. It contains several state variables and functions to interact with the API.

### Functions:

#### 2.1. getUser
Returns the currently authenticated Firebase user.

#### 2.2. getAllUsers
Fetches all users from the API and logs the data to the console.

#### 2.3. getUserById
Fetches a user by their Firebase ID and logs the data to the console.

#### 2.4. getRandomPhoto
Fetches a random photo from the API, creates an object URL from the received blob, and updates the `randomPhotoUrl` state.

#### 2.5. getPhotoById
Fetches a photo by its ID, creates an object URL from the received blob, and updates the `photoUrl` state.

#### 2.6. getProfilePicture
Fetches the authenticated user's profile picture, creates an object URL from the received blob, and updates the `photoUrl` state.

#### 2.7. uploadPhoto
Uploads a photo to the server using the API.

#### 2.8. uploadProfilePhoto
Uploads a profile picture for the authenticated user using the API.

#### 2.9. updateBio
Updates the bio of the authenticated user using the API.

#### 2.10. search_username
Fetches a user by their username and logs the data to the console.


## Working on nearby users

