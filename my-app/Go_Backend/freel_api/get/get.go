package get

import (
	"context"
	"encoding/json"
	"fmt"
	
	"testing"

	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Get a MongoDB client and collection
	client := mongo.GetMongoClient()

	collection := client.Database("freel").Collection("users")
	// Query the collection to get all the users
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Printf("Error finding documents: %v", err)
		http.Error(w, "Error finding documents", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var users []bson.M
	for cursor.Next(context.Background()) {
		var user bson.M
		err := cursor.Decode(&user)
		if err != nil {
			log.Printf("Error decoding document: %v", err)
			http.Error(w, "Error decoding document", http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Printf("Cursor error: %v", err)
		http.Error(w, "Cursor error", http.StatusInternalServerError)
		return
	}

	// Return the users as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter and convert it to a primitive.ObjectID
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get a MongoDB client and collection
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Query the collection for the user with the given ID
	var user bson.M
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		http.Error(w, "Error finding document", http.StatusInternalServerError)
		return
	}

	// Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}

func GetUserById_post(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter and convert it to a primitive.ObjectID
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get a MongoDB client and collection
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Query the collection for the user with the given ID
	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		http.Error(w, "Error finding document", http.StatusInternalServerError)
		return
	}

	// Return the user's posts as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user.Posts)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}



func Get_Photo(w http.ResponseWriter, r *http.Request) {
	// Get the photo ID from the URL parameter and convert it to an ObjectID
	params := mux.Vars(r)
	photoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)

	}

	// Get the MongoDB client and the GridFS bucket for photos
	client := mongo.GetMongoClient()

	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	// Find the file in the bucket and open a download stream for it
	file, err := bucket.OpenDownloadStream(photoID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	defer file.Close()

	// Read the photo data into a byte slice
	data, err := ioutil.ReadAll(file)

	// Write the data to the response writer
	w.Write(data)

}
func TestGetPhoto(t *testing.T) {
	req, err := http.NewRequest("GET", "/random-image", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get_Photo)

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the content type
	expectedContentType := "image/jpg"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}

	// Check if the response body is not empty
	if rr.Body.Len() == 0 {
		t.Error("handler returned empty response body")
	}
}


/*

func GetUserPost(w http.ResponseWriter, r *http.Request, userID string) ([]map[string]interface{}, error) {




	// Get the MongoDB client and the "users" collection
	client := mongo.GetMongoClient()

	collection := client.Database("freel").Collection("users")

	// Find the user document with the specified ID
	filter := bson.M{"_id": userID}
	var user map[string]interface{}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	// Extract the "posts" field from the user document and return it as an array of maps
	posts, ok := user["posts"].([]map[string]interface{})
	if !ok {
		return nil, errors.New("posts field is not an array of maps")
	}

	return posts



}

*/

func GetUserPosts_Help(w http.ResponseWriter, r *http.Request) {
	/*
		params := mux.Vars(r)
		userID, err := primitive.ObjectIDFromHex(params["id"])
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)

		}

		results, error := GetUserPost(w , r, userID);
	*/

}

func Update_Many() {

	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Add the `saved_post` field to all documents in the collection
	filter := bson.M{}
	update := bson.M{"$set": bson.M{"saved_post": []Post{}}}
	result, err := collection.UpdateMany(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated %v documents\n", result.ModifiedCount)

}

func Test_Get_User(t *testing.T) {

	id := "63f5687adcf9b9a96ad516a4"

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

func TestGet_User(t *testing.T) {

	id := "63f5687adcf9b9a96ad516a4"

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
