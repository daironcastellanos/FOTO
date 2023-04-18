package put

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
    Lat         float64   `bson:"lat" json:"lat"`
    Lng         float64   `bson:"lng" json:"lng"`
    Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FireID      string             `bson:"FireID" json:"FireID"`
	FullName    string             `bson:"FullName" json:"FullName"`
	Username    string             `bson:"Username" json:"Username"`
	Email       string             `bson:"Email" json:"Email"`
	Bio         string             `bson:"Bio" json:"Bio"`
	Location    Location           `bson:"Location" json:"Location"`
	DOB         string             `bson:"DOB" json:"DOB"`
	Followers   []string           `bson:"Followers" json:"Followers"`
	Following   []string           `bson:"Following" json:"Following"`
	MyPhotos    []string           `bson:"MyPhotos" json:"MyPhotos"`
	SavedPhotos []string           `bson:"SavedPhotos" json:"SavedPhotos"`
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
	params := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get the image data from the request body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Upload photo to GridFS
	photoID, err := Upload_Photo(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the MongoDB client and the users collection
	client := mongo.GetMongoClient()
	userCollection := client.Database("freel").Collection("users")

	// Find the user and update their posts array with the photo ID
	filter := bson.M{"_id": userID}
	update := bson.M{"$addToSet": bson.M{"posted_pics": photoID}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Post added successfully")
}


