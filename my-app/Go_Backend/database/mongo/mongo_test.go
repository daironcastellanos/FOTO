package mongo_test

import (
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"your-project/mongo" // replace with your package import path
)

const mongoURI = "mongodb://localhost:27017"

func TestInitiateMongoClient(t *testing.T) {
	client := mongo.InitiateMongoClient(mongoURI)
	defer client.Disconnect(context.Background())

	require.NotNil(t, client, "mongo client should not be nil")
}

func TestFindUserByID(t *testing.T) {
	// connect to the database and insert a test user
	client := mongo.InitiateMongoClient(mongoURI)
	defer client.Disconnect(context.Background())

	userID := primitive.NewObjectID().Hex()
	testUser := &mongo.User{
		Name: "Test User",
		Bio:  "I am a test user",
		Posts: []mongo.Post{
			{
				Title: "Test Post",
				Body:  "This is a test post",
				Tags:  []string{"test", "mongo"},
				Date:  "2022-01-01",
				Image: "test.jpg",
			},
		},
	}
	_, err := client.Database("freel").Collection("users").InsertOne(context.Background(), testUser)
	require.NoError(t, err, "should not get error when inserting test user")

	// call FindUserByID and check the returned user matches the test user
	foundUser, err := mongo.Find_User_By_ID(userID, mongoURI)
	require.NoError(t, err, "should not get error when finding test user")
	assert.Equal(t, testUser.Name, foundUser.Name, "user name should match")
	assert.Equal(t, testUser.Bio, foundUser.Bio, "user bio should match")
	assert.Equal(t, testUser.ProfilePicture, foundUser.ProfilePicture, "user profile picture should match")
	assert.Equal(t, testUser.Posts[0].Title, foundUser.Posts[0].Title, "user post title should match")
	assert.Equal(t, testUser.Posts[0].Body, foundUser.Posts[0].Body, "user post body should match")
	assert.Equal(t, testUser.Posts[0].Tags, foundUser.Posts[0].Tags, "user post tags should match")
	assert.Equal(t, testUser.Posts[0].Date, foundUser.Posts[0].Date, "user post date should match")
	assert.Equal(t, testUser.Posts[0].Image, foundUser.Posts[0].Image, "user post image should match")

	// clean up test user
	result, err := client.Database("freel").Collection("users").DeleteOne(context.Background(), bson.M{"_id": userID})
	require.NoError(t, err, "should not get error when deleting test user")
	fmt.Printf("Deleted %d document(s)\n", result.DeletedCount)
}

