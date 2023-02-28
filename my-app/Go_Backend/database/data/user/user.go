package user

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

type User struct {
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilepicture"`
	Posts          []Post `json:"posts"`
}

type Post struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
}

func Insert_User(user User, URI_ string) error {
	// Set up a MongoDB client and connect to the database
	

	// Insert the user profile document into the "users" collection
	collection := client.Database("freel").Collection("users")
	_, err = collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	fmt.Println("User inserted successfully")
	return nil
}

func Create_Test_User(URI_ string) {
	/* Check if already a user */

	fmt.Println("Creating sample data to insert into mongo since no data was given")
	test_user := User{
		Name:           "John Doe",
		Bio:            "I'm a software engineer and hobbyist photographer.",
		ProfilePicture: "https://example.com/profile.jpg",
		Posts: []Post{
			{
				Title: "My First Post",
				Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
				Tags:  []string{"programming", "photography"},
				Date:  "2022-01-01T12:00:00Z",
				Image: "https://example.com/post1.jpg",
			},
			{
				Title: "My Second Post",
				Body:  "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
				Tags:  []string{"travel", "food"},
				Date:  "2022-01-05T12:00:00Z",
				Image: "https://example.com/post2.jpg",
			},
		},
	}

	Insert_User(test_user, URI_)
}

func Update_User_Bio(newBio string, URI_ string, userID string) error {
	// Set up a MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI(URI_)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Find the user with the specified ObjectID and update their bio
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	collection := client.Database("freel").Collection("users")
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"bio": newBio}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("User with ID %s updated successfully\n", userID)
	return nil
}
