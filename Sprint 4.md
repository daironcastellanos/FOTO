# Sprint 4 FREEL


# Front-End

# SPRINT 4 BACK-END

## Backend API Documentation - Sprint 4 Updated
### Introduction
In Sprint 4, we've made significant improvements to our backend API by introducing new user endpoints and additional user-related features. These updates allow for actions such as creating a new user, retrieving users by FireID or username, updating user information, and managing follower and following lists. We've also implemented a nearby users endpoint to discover users based on location. To ensure a smooth experience, we've provided documentation on utilizing the GET, POST, and DELETE methods and included helpful commands for package installation, running unit tests, and executing the backend application. With a strong focus on robust testing, our team has effectively verified the correct functionality of key endpoints, bolstering the quality and reliability of our backend API.


### How to use it
Some helpful commands when you start using the Go_Backend directory are:
```
go mod tidy
go install
```
to make sure all the packages are installed.
Also you can run the unit tests with  
```
go test
```
Finally you can run the backend application with
```
go run main.go
```
To do testing of the backend application
```
go test -v
```
on the directory of my-app/src/pages/api

#### What's New
Created User endpoints:
* POST /api/create/user for creating a new user
* GET /api/users/{fireID}/get for getting a user by their FireID
* GET /api/username/{username}/get for getting a user by their username
* GET /api/users/get for getting all users
Implemented additional user-related features:

* POST /api/users/{fireID}/update_bio for updating a user's bio
* POST /api/users/{fireID}/uploadProfilePicture for uploading a profile picture
* GET /api/users/{fireID}/getProfilePicture for retrieving a profile picture
* POST /api/users/{fireId}/updateLocation for updating a user's location
* POST /api/users/{fireId}/savePhoto for adding a photo to saved photos
* POST /api/users/{fireId}/removePhoto for removing a photo from saved photos
* POST /api/users/{fireId}/follow for following a user
* POST /api/users/{fireId}/unfollow for unfollowing a user

Created Follower and Following endpoints:
* POST /api/users/{fireID}/addFollower/{followerID} for adding a follower
* POST /api/users/{fireID}/removeFollower/{followerID} for removing a follower
* POST /api/users/{fireID}/addFollowing/{followingID} for adding to the following list
* POST /api/users/{fireID}/removeFollowing/{followingID} for removing from the following list

Implemented a Nearby users endpoint:
* GET /api/nearby_users/{fireID} for getting nearby users based on location

#### Unit Tests
The TestGetUserByFireID function is a unit test for the GetUserByFireID endpoint in the back end. This test verifies that the endpoint correctly retrieves a user by their FireID and returns the expected user information.

The test follows these steps:
* Creates a test user with unique values and a known FireID.
* Constructs a new HTTP GET request with the test user's FireID.
* Creates a ResponseRecorder to record the HTTP response.
* Calls the GetUserByFireID function with the constructed request.
* Checks if the HTTP response status code is 200 (OK).
* Decodes the response body into a User struct.
* Verifies that the returned user matches the test user, specifically comparing the FireID and FullName fields.
* This test helps ensure that the GetUserByFireID function is working as intended and returns accurate user data based on the provided FireID.


The TestGetUserByUsername function is a unit test for the GetUserByUsername endpoint in the back end. This test ensures that the endpoint accurately retrieves a user by their username and returns the expected user information.

The test follows these steps:
* Creates a test user with unique values and a known username.
* Constructs a new HTTP GET request with the test user's username.
* Creates a ResponseRecorder to record the HTTP response.
* Calls the GetUserByUsername function with the constructed request.
* Checks if the HTTP response status code is 200 (OK).
* Decodes the response body into a User struct.
* Verifies that the returned user matches the test user, specifically comparing the Username and FullName fields.

### Getting Started
To effectively utilize the GET, POST, and DELETE methods in the backend API, we defined API endpoints corresponding to desired actions, such as creating, retrieving, or deleting resources. We implemented HTTP handlers for each endpoint, using the appropriate method to perform the corresponding action. To ensure the handlers worked correctly, we wrote test cases using Go's built-in testing package, verifying that they returned expected results. This approach was applied to all three methods, resulting in a robust and reliable backend API.

### Hosting API Services
We used a Docker Image to upload the Go Backend service to google cloud Run so we didnt have to always have the backend running to call the Front End and then connected the frontend thats hosted on vercel to the backend on gooogle cloud run server.

### Conclusion
Sprint 4 has brought significant enhancements to the backend API by introducing new user endpoints, additional user-related features, and a nearby users endpoint. These updates streamline actions such as creating, retrieving, and updating users, as well as managing follower and following lists. The provided documentation and commands facilitate package installation, running unit tests, and executing the backend application. Through rigorous testing and verification, the quality and reliability of the backend API have been effectively bolstered, resulting in a robust and user-friendly experience.