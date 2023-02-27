package post

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"Freel.com/freel_api/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
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
}

var userCollection = mongo.Get_User_Collection()

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a User struct
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the collection
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the inserted user's ID as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	id := result.InsertedID.(primitive.ObjectID)
	if err := json.NewEncoder(w).Encode(id.Hex()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Create_Account(user User) {

	collection := mongo.Get_User_Collection()

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User inserted successfully")

}

func add_info() {

}

func add_location() {

}

func Create_Fake_Account(w http.ResponseWriter, r *http.Request) {

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

	Create_Account(test_user)
}
