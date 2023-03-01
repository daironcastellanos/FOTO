package delete

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


func Delete_User(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter and convert it to a primitive.ObjectID
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Delete the user with the given ID from the collection
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")
	result, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Println(err)
		return
	}

	// Return the number of deleted documents as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result.DeletedCount); err != nil {
		log.Println(err)
		return
	}
}
	



func Delete_Pic(w http.ResponseWriter, r *http.Request) {
	// Get the photo ID from the URL parameter and convert it to a primitive.ObjectID
	params := mux.Vars(r)
	photoID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Open a GridFS bucket named "photos"
	bucket, err := mongo.Get_Photo_Bucket()
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to get bucket", http.StatusInternalServerError)
		return
	}

	// Delete the photo with the specified ObjectID
	err = bucket.Delete(photoID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Failed to delete photo", http.StatusInternalServerError)
		return
	}

	fmt.Printf("Photo with ID %s deleted successfully\n", photoID.Hex())
	w.WriteHeader(http.StatusOK)
}




func Delete_User_Location(){




}