package get

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"

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

func Get_Users(w http.ResponseWriter, r *http.Request) {

	collection := mongo.Get_User_Collection()

	var users []User
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.Background())
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

	collection := mongo.Get_User_Collection()

	// Get the user ID from the URL parameter and query the collection
	params := mux.Vars(r)
	var user User
	err := collection.FindOne(context.Background(), bson.M{"_id": params["id"]}).Decode(&user)
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

func Get_Photos(w http.ResponseWriter, r *http.Request) {

	bucket,error := mongo.Get_Photo_Bucket()
	if(error != nil){
		log.Println(error)
	}

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
		return
	}

}