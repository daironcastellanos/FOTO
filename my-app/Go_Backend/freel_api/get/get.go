package get

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"io/ioutil"
	"log"
	"net/http"

	"bytes"

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
    Title    string     `json:"title"`
    Body     string     `json:"body"`
    Tags     []string   `json:"tags"`
    Date     string     `json:"date"`
    Image    string     `json:"image"`
    Likes    []Like     `bson:"likes,omitempty" json:"likes"`
    Comments []Comment  `bson:"comments,omitempty" json:"comments"`
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

func Get_Users(w http.ResponseWriter, r *http.Request) {
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Get the total number of documents in the collection
	total, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Total number of documents: %d", total)

	// Query the collection to get all the users
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	var users []User
	for cursor.Next(context.Background()) {
		var user User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	// Return the users as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Println(err)
		return
	}
}

func Get_User(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter and convert it to a primitive.ObjectID
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
		return
	}

	// Query the collection for the user with the given ID
	collection := mongo.Get_User_Collection()
	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}

	// Return the user as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println(err)
		return
	}
}

func Get_Photos(w http.ResponseWriter, r *http.Request)([][]byte) {

	client := mongo.GetMongoClient()

	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	

	// Get all the files in the bucket
	filter := bson.M{}
	cursor, err := bucket.Find(filter)

	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())

	// Loop through the files and read their data
	var photos [][]byte
	for cursor.Next(context.Background()) {
		// Get the file information
		fileInfo := &gridfs.File{}
		err := cursor.Decode(fileInfo)
		if err != nil {
			log.Fatal(err)
		}

		// Open a download stream for the file
		downloadStream, err := bucket.OpenDownloadStream(fileInfo.ID)
		if err != nil {
			log.Fatal(err)
		}
		defer downloadStream.Close()

		// Read the file data into a byte slice
		data :=
			make([]byte, fileInfo.Length)
		_, err = downloadStream.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		// Append the photo data to the slice of byte slices
		photos = append(photos, data)
	}

	// Return the photos as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(photos); err != nil {
		log.Println(err)
		
	}

	return photos

}

func Get_Photo(w http.ResponseWriter, r *http.Request) ([]byte, error) {
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

	// Set the content type header based on the file's metadata
	
	// Read the photo data into a byte slice
	data, err := ioutil.ReadAll(file)
	

	// Print the length of the data to confirm that we have read the file
	fmt.Printf("Read %d bytes\n", len(data))

	return data, err
}



func Serve_Pics(w http.ResponseWriter, r *http.Request) {
	// Load the MONGO_URI from the .env file

	// Get the image data (e.g. from a file or database)
	data := Get_Photos(w, r)
	
	// Concatenate the byte slices into a single byte slice
	var buf bytes.Buffer
	for _, b := range data {
		buf.Write(b)
	}
	imgData := buf.Bytes()

	// Set the Content-Type header to indicate that this is an image
	w.Header().Set("Content-Type", "image/jpeg")

	// Write the image data to the response
	if _, err := w.Write(imgData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
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

func GetUserPosts_Help(w http.ResponseWriter, r *http.Request){
	/*
	params := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		
	}

	results, error := GetUserPost(w , r, userID);
	*/

}




func Update_Many(){

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