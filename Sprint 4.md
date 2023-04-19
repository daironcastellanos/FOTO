# Sprint 4 FREEL


# Front-End



# SPRINT 4 BACK-END
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

Configured CORS settings and started the server on port 8080.
## Backend API Documentation - GET, POST, and DELETE Methods
### Introduction



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



#### New Issues to work on the next sprint
* Post_Pic
The Post_Pic issue involves implementing functionality to allow users to upload pictures to the application. This involves creating an API endpoint that allows users to submit a POST request with a picture, and implementing an HTTP handler function that will save the picture to a database (such as MongoDB) and return a response to the client. This issue may also involve adding functionality to resize images and validate file types.

* Update_Bio
The Update_Bio issue involves implementing functionality to allow users to update their profile bio. This involves creating an API endpoint that allows users to submit a PUT request with a new bio, and implementing an HTTP handler function that will update the user's bio in a database (such as MongoDB) and return a response to the client.

* Update_Profile_on_uploads
The Update_Profile_on_uploads issue involves implementing functionality to automatically update a user's profile picture when they upload a new picture. This involves modifying the Post_Pic implementation to also update the user's profile picture in the database, and implementing an HTTP handler function that will return the updated profile picture to the client.

#### Conclusion
By implementing the GET, POST, and DELETE methods in your backend API, you can provide powerful functionality that allows clients to retrieve, create, and delete information from the server. To get started with implementing these methods, you will need to define your API endpoints, implement your HTTP handlers