package put

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"
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



func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")
	

	// Update the user with the given ID in the collection
	update := bson.M{
		"$set": bson.M{
			"name":           user.Name,
			"bio":            user.Bio,
			"profilepicture": user.ProfilePicture,
			"posts":          user.Posts,
		},
	}
	if _, err := collection.UpdateOne(context.Background(), bson.M{"_id": id}, update); err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Encode the updated user as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

type UserUpdate struct {
	Bio string `json:"bio"`
}




func Update_Bio(w http.ResponseWriter, r *http.Request) {
	// Get the user ID and bio string from the URL parameters
	params := mux.Vars(r)
	objectID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	bio := params["bio"]

	// Update the user's bio in the database
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"bio": bio}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User with ID %s updated successfully\n", params["id"])
}

func Upload_Photo(data []byte) (primitive.ObjectID, error) {
	// Get the MongoDB client and the GridFS bucket for photos
	client := mongo.GetMongoClient()
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("images"),
	)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	// Upload the image data to the bucket
	uploadStream, err := bucket.OpenUploadStream("photo")
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(data)
	if err != nil {
		log.Println(err)
		return primitive.NilObjectID, err
	}

	// Get the ObjectID of the uploaded photo
	photoID, ok := uploadStream.FileID.(primitive.ObjectID)
	if !ok {
		log.Println("Error asserting FileID to primitive.ObjectID")
		return primitive.NilObjectID, fmt.Errorf("Error asserting FileID to primitive.ObjectID")
	}

	// Return the photo ID
	return photoID, nil
}

func Post_Pic(w http.ResponseWriter, r *http.Request){

	// Get the user ID from the URL parameter

	// Upload photo to GridFS

	// Get the photo ID from the GridFS upload

	// if user has no postedd_pics array, create one
	// Create a new array in users called posteddd_pics
	// else add the photo ID to the array

	// Update the user's posts array with the ID of the photo
}
 



