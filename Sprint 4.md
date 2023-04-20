# Sprint 4 FREEL

# Front-End

#### What's New

Sprint 4 has been an incredibly intense and transformative period for our team, with numerous improvements and new features added to the platform. The front-end team dedicated their efforts to a complete overhaul of the design, creating an immersive and enhanced user experience. Here's an in-depth look at the developments made during this sprint:

- Landing Page Transformation: We kicked off Sprint 4 by developing an entirely new landing page. Before users log in, they are directed to this modern and engaging homepage. The landing page offers an interactive experience, allowing users to explore our features, understand the software's benefits, and learn about our team's background and expertise. Additionally, users can easily sign up or log in from the landing page itself.

- Streamlined Sign-up and Login Process: We took a fresh approach to the sign-up and login process, focusing on simplicity and visual appeal. Our goal was to make it easy and enjoyable for users to access our platform, minimizing any potential friction points.

- Redesigned Feed Experience: Once users successfully log in, they are greeted with a beautifully designed feed that has been restructured for optimal user engagement. We divided the feed page into three sections:

  - Left Section: This section houses the main navigation menu, providing users with a seamless experience as they explore the platform. It includes buttons for home, profile, upload, settings, and search. Additionally, a logout button is conveniently located at the bottom left for easy access.
  - Middle Section: The central part of the feed, also known as the "infinite feed," showcases all the photos uploaded by the people the user follows. Each picture features the uploader's username and bio, fostering a sense of community and connection among users.
  - Right Section: This section not only offers search functionality but also includes a list of suggested followers, enabling users to discover top users in their area and expand their network.

- Revamped Profile Experience: The profile page underwent significant enhancements, resulting in a more dynamic and intuitive user experience:

  - UI Design: The profile page now boasts a visually stunning and user-friendly interface, enabling users to easily browse and view all pictures posted by a particular user.
  - Follow and Book Features: Users can now follow other users or book photographers directly from their profiles. As our primary target audience is photographers, this feature facilitates potential clients to book them based on their showcased portfolios.

- Responsive Design: The platform has been optimized to work seamlessly across various devices and screen sizes, ensuring a consistent and enjoyable experience for users, regardless of the device they use.

- Performance Enhancements: During this sprint, we also focused on improving the platform's performance, implementing various optimizations to ensure a smooth and responsive user experience. This includes faster loading times, reduced latency, and efficient resource usage.

In conclusion, Sprint 4 marked a significant turning point in our platform's development, with the front-end team working tirelessly to create a visually impressive and user-friendly interface. The introduction of new features, coupled with an enhanced design and improved performance, has set the stage for continued growth and expansion of our platform.

### How to run Freel:

1. Prerequisites: Make sure Node.js and npm (Node Package Manager) is installed. This can be checked by running the following commands in the terminal:

```
    node -v
    npm -v
```

If Node.js and npm are not installed, they can download and install them from the official website: https://nodejs.org/

2. Clone the repository: Clone the application's repository to your local machine using a Git client or by running the following command in the terminal:

```
git clone https://github.com/daironcastellanos/Freel.git
```

3. Navigate to the project folder: Navigate to the my-app folder within the cloned repository using the terminal:

```
cd path/to/Freel/my-app
```

Replace path/to with the appropriate path on the your machine.

4. Install dependencies: Next, you should run the following command to install all the required dependencies for the project:

```
npm install
```

5. Start the development server: Once the dependencies are installed, you need to start the development server by running:

```
npm run dev
```

6. Access the application: Finally, you can open their preferred web browser and navigate to the application at the following address:

```
http://localhost:3000
```

Now, you should see the application running in their browser, and they can start exploring its features.

### Front-end Testing

🧪Test the Upload button: This test checks whether the "Upload" button is functioning correctly by navigating to the Upload page when clicked.

🧪Test the Choose File button: This test verifies that the "Choose File" button on the Upload page is operational and allows users to select a file.

🧪Test the Back button on Upload page: This test ensures that the "Back" button on the Upload page is functioning correctly, taking the user back to the Home page when clicked.

🧪Test the Search bar functionality: This test checks if the search bar on the top of the page is working correctly by allowing users to type in the search field and remove the text using the provided arrow button.

🧪Test the Book Photographer button: This test confirms that the "Book Photographer" button is working correctly, navigating to the appropriate page when clicked.

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

# BACK-END

#### What's New

Created User endpoints:

- POST /api/create/user for creating a new user
- GET /api/users/{fireID}/get for getting a user by their FireID
- GET /api/username/{username}/get for getting a user by their username
- GET /api/users/get for getting all users
  Implemented additional user-related features:

- POST /api/users/{fireID}/update_bio for updating a user's bio
- POST /api/users/{fireID}/uploadProfilePicture for uploading a profile picture
- GET /api/users/{fireID}/getProfilePicture for retrieving a profile picture
- POST /api/users/{fireId}/updateLocation for updating a user's location
- POST /api/users/{fireId}/savePhoto for adding a photo to saved photos
- POST /api/users/{fireId}/removePhoto for removing a photo from saved photos
- POST /api/users/{fireId}/follow for following a user
- POST /api/users/{fireId}/unfollow for unfollowing a user

Created Follower and Following endpoints:

- POST /api/users/{fireID}/addFollower/{followerID} for adding a follower
- POST /api/users/{fireID}/removeFollower/{followerID} for removing a follower
- POST /api/users/{fireID}/addFollowing/{followingID} for adding to the following list
- POST /api/users/{fireID}/removeFollowing/{followingID} for removing from the following list

Implemented a Nearby users endpoint:

- GET /api/nearby_users/{fireID} for getting nearby users based on location

#### Unit Tests

The TestGetUserByFireID function is a unit test for the GetUserByFireID endpoint in the back end. This test verifies that the endpoint correctly retrieves a user by their FireID and returns the expected user information.

The test follows these steps:

- Creates a test user with unique values and a known FireID.
- Constructs a new HTTP GET request with the test user's FireID.
- Creates a ResponseRecorder to record the HTTP response.
- Calls the GetUserByFireID function with the constructed request.
- Checks if the HTTP response status code is 200 (OK).
- Decodes the response body into a User struct.
- Verifies that the returned user matches the test user, specifically comparing the FireID and FullName fields.
- This test helps ensure that the GetUserByFireID function is working as intended and returns accurate user data based on the provided FireID.

The TestGetUserByUsername function is a unit test for the GetUserByUsername endpoint in the back end. This test ensures that the endpoint accurately retrieves a user by their username and returns the expected user information.

The test follows these steps:

- Creates a test user with unique values and a known username.
- Constructs a new HTTP GET request with the test user's username.
- Creates a ResponseRecorder to record the HTTP response.
- Calls the GetUserByUsername function with the constructed request.
- Checks if the HTTP response status code is 200 (OK).
- Decodes the response body into a User struct.
- Verifies that the returned user matches the test user, specifically comparing the Username and FullName fields.

### Getting Started

To effectively utilize the GET, POST, and DELETE methods in the backend API, we defined API endpoints corresponding to desired actions, such as creating, retrieving, or deleting resources. We implemented HTTP handlers for each endpoint, using the appropriate method to perform the corresponding action. To ensure the handlers worked correctly, we wrote test cases using Go's built-in testing package, verifying that they returned expected results. This approach was applied to all three methods, resulting in a robust and reliable backend API.

### Hosting API Services

We used a Docker Image to upload the Go Backend service to google cloud Run so we didnt have to always have the backend running to call the Front End and then connected the frontend thats hosted on vercel to the backend on gooogle cloud run server.

### Conclusion

Sprint 4 has brought significant enhancements to the backend API by introducing new user endpoints, additional user-related features, and a nearby users endpoint. These updates streamline actions such as creating, retrieving, and updating users, as well as managing follower and following lists. The provided documentation and commands facilitate package installation, running unit tests, and executing the backend application. Through rigorous testing and verification, the quality and reliability of the backend API have been effectively bolstered, resulting in a robust and user-friendly experience.
