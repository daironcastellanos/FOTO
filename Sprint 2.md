# SPRINT2 FRONT-END

Features:

$ We added a sleek and intuitive top bar that includes profile information, settings, and an upload button. Additionally, we included our application's       logo for branding purposes.

$ To provide a better user experience, we created multiple pages for our application, including a homepage and a profile page.

$ To showcase dynamic content, we implemented a scrolling view component that fetches live images from a server. The images are displayed in a responsive     grid layout, making it easy for users to view and interact with them.

Tests:

$ To ensure that our scrolling view component was fetching images correctly, we implemented Cypress tests that simulate user behavior and validate that the images are being displayed properly.

$ We also created Cypress tests for our navigation buttons to ensure that users are being directed to the correct pages when clicking them.

$ Tested Back button , tested navegation button , tested images fetch
Tested navigation and image fetching functionality, as well as the back button, for our web application.


Documentation:
Using Cypress 
Simply run command 
npx cypress open

# SPRINT2 BACK-END
Some of out accomplishments for Sprint 2 is we setup mongo db Atlases that has a users collection and a photo bucket that stores the photos in chucks to optimize efficency, These elements are accessed through the mongoDB client interface, which is called in individual functions.

The front end will use these functions to access the mongo Database by the Rest API we implemented using gorilla mux. We made functions and end points that are commented out in the video because we have not written unit test but these endpoints will be used to get photos from the bucket for users feed and we plan on making the function to post pictures accessible from the Gorilla mux router.

## Backend API Documentation - GET, POST, and DELETE Methods
### Introduction
This documentation explains how to use the GET, POST, and DELETE methods to implement a backend API that includes the actions of retrieving, creating, and deleting information. These methods are commonly used in HTTP-based APIs to communicate between the client and the server.


### How to use it
Some helpful commands when you start using the Go_Backend directory are:
'''
go mod tidy
go install
'''
to make sure all the packages are installed.
Also you can run the unit tests with  
'''
go test
'''
Finally you can run the backend application with
'''
go run main.go
'''

### Getting Started
To use the GET, POST, and DELETE methods in the backend API, we did the following:

Defined the API endpoints that correspond to the actions that you want to perform. For example, you might create an endpoint /users that returns a list of all users in the system.

Implemented the HTTP handlers that correspond to the endpoints that we  defined. These handlers should use the appropriate HTTP method (such as GET, POST, or DELETE) to perform the action that corresponds to the endpoint.

Wrote test cases for handlers using a testing framework such as Go's built-in testing package or a third-party package like goconvey or testify. On this case, we used the testing package from GO, and tested all the three methods. These tests verify that the handlers are working correctly and that they return the expected results.

### GET Method
To implement the GET method in the backend API, we did the following:

Defined an API endpoint that corresponds to the action that we want to perform. For example, we created an endpoint /users that returns a list of all users in the system.

Implemented an HTTP handler function that corresponds to the endpoint that you defined. This handler should use the GET method to retrieve the requested information from the server and return it to the client.

Wrote test cases for your handler function that verify that it is working correctly and that it returns the expected results. For example, you might write a test case that sends a GET request to the /users endpoint and verifies that the response contains a list of users.

### POST Method
To implement the POST method in the backend API, we did the following:

Defined an API endpoint that corresponds to the action that you want to perform. For example, we created an endpoint /users/new that allows clients to create new user accounts.

Implemented an HTTP handler function that corresponds to the endpoint that you defined. This handler should use the POST method to create a new resource on the server based on the data sent in the request body.

Wrote test cases for the handler function that verify that it is working correctly and that it creates the expected resources on the server. For example, we wrote a test case that sends a POST request to the /users/new endpoint with a JSON payload representing a new user account, and then verifies that the user account was created on the server.

### DELETE Method
To implement the DELETE method in your backend API, we did the following:

Defineed an API endpoint that corresponds to the action that you want to perform. For example, you might create an endpoint /users/delete that allows clients to delete user accounts.

Implemented an HTTP handler function that corresponds to the endpoint that you defined. This handler should use the DELETE method to delete the specified resource on the server.

Wrote test cases for your handler function that verify that it is working correctly and that it deletes the expected resources on the server. For example, you might write a test case that sends a DELETE request to the /users/delete endpoint with the ID of a user account to be deleted, and then verifies that the user account was deleted on the server.

#### Conclusion
By implementing the GET, POST, and DELETE methods in your backend API, you can provide powerful functionality that allows clients to retrieve, create, and delete information from the server. To get started with implementing these methods, you will need to define your API endpoints, implement your HTTP handlers

