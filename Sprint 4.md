# Sprint 4 FREEL


# Front-End



# SPRINT 4 BACK-END

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




#### What's New

Here are some additional functions that have been added to the backend API:

Get_Photo
This function retrieves a photo from the MongoDB database using the specified photo ID. The function reads the photo data into a byte slice and writes it to the HTTP response writer. We are continue working on the unit testing of this function.

Get_Random_Picture
This function retrieves a random picture from the MongoDB database. The function uses the math/rand package to generate a random number, and then uses the number to select a picture from the database.

Get_User_Post_by_ID
This function retrieves the posts created by a user with the specified user ID. The function uses the MongoDB Aggregate function to perform a join between the users and posts collections, and then returns the posts created by the user.

GetUserById_post 
This function returns an array of unique photo IDs. from the user object 

To use these functions, you will need to define additional API endpoints and implement HTTP handlers that correspond to the functions. You will also need to write test cases for the handlers to verify that they are working correctly.

#### New Issues to work on the next sprint
* Post_Pic
The Post_Pic issue involves implementing functionality to allow users to upload pictures to the application. This involves creating an API endpoint that allows users to submit a POST request with a picture, and implementing an HTTP handler function that will save the picture to a database (such as MongoDB) and return a response to the client. This issue may also involve adding functionality to resize images and validate file types.

* Update_Bio
The Update_Bio issue involves implementing functionality to allow users to update their profile bio. This involves creating an API endpoint that allows users to submit a PUT request with a new bio, and implementing an HTTP handler function that will update the user's bio in a database (such as MongoDB) and return a response to the client.

* Update_Profile_on_uploads
The Update_Profile_on_uploads issue involves implementing functionality to automatically update a user's profile picture when they upload a new picture. This involves modifying the Post_Pic implementation to also update the user's profile picture in the database, and implementing an HTTP handler function that will return the updated profile picture to the client.

#### Conclusion
By implementing the GET, POST, and DELETE methods in your backend API, you can provide powerful functionality that allows clients to retrieve, create, and delete information from the server. To get started with implementing these methods, you will need to define your API endpoints, implement your HTTP handlers