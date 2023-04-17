package get_test

import (
	"context"


	"Freel.com/freel_api/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"

	"testing"
)

type Like struct {
	Username string `bson:"username,omitempty" json:"username"`
	Date     string `bson:"date,omitempty" json:"date"`
}

type Comment struct {
	Username string `bson:"username,omitempty" json:"username"`
	Date     string `bson:"date,omitempty" json:"date"`
	Comment  string `bson:"comment,omitempty" json:"comment"`
}

type Post struct {
	gorm.Model
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []string  `json:"tags"`
	Date     string    `json:"date"`
	Image    string    `json:"image"`
	Likes    []Like    `bson:"likes,omitempty" json:"likes"`
	Comments []Comment `bson:"comments,omitempty" json:"comments"`
}

type Location struct {
	Type        string    `bson:"type,omitempty" json:"type"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string             `bson:"name,omitempty" json:"name"`
	Bio            string             `bson:"bio,omitempty" json:"bio"`
	ProfilePicture string             `bson:"profilepicture,omitempty" json:"profilepicture"`
	Posts          []Post             `bson:"posts,omitempty" json:"posts"`
	Location       Location           `bson:"location,omitempty" json:"location"`
	SavedPosts     []Post             `bson:"saved_post,omitempty" json:"saved_post"`
}

func TestGet_User(t *testing.T) {

	id := "63f5687adcf9b9a96ad516a4";
	

	objectID, err := primitive.ObjectIDFromHex(id)
  
	// Query the collection for the user with the given ID
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")
	

	var user User
    err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
    if err != nil {
        return 
    }

	// Check if the user ID matches the expected value
	expectedID, _ := primitive.ObjectIDFromHex(id)
	if user.ID != expectedID {
		t.Errorf("User ID mismatch: expected %s, got %s", expectedID, user.ID)
	}
}



