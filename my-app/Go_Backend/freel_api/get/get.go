package get

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type Post struct {
	gorm.Model
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
}


type User struct {
	gorm.Model
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilepicture"`
	Posts          []Post `json:"posts"`
}


func Get_Users(w http.ResponseWriter, r *http.Request){

	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

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

    // Print the users to the console
    fmt.Println(users)

	//return users
}





func Get_User(w http.ResponseWriter, r *http.Request){

	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	 // Get the user ID from the URL parameter and query the collection
	 params := mux.Vars(r)
	 var user User
	 err := collection.FindOne(context.Background(), bson.M{"_id": params["id"]}).Decode(&user)
	 if err != nil {
		 log.Println(err)
		 return
	 }
 
	 // Return the user as JSON
	 if err := json.NewEncoder(w).Encode(user); err != nil {
		 log.Println(err)
		 return
	 }

}



func Get_Photos(w http.ResponseWriter, r *http.Request){

	client := mongo.GetMongoClient()
	

	// Open a GridFS bucket named "photos"
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		return nil, err
	}

	// Get all the files in the bucket
	filter := bson.M{}
	cursor, err := bucket.Find(filter)
	if err != nil {
		return nil, err
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
		data := make([]byte, fileInfo.Length)
		_, err = downloadStream.Read(data)
		if err != nil {
			log.Fatal(err)
		}

		// Append the photo data to the slice of byte slices
		photos = append(photos, data)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	//return photos, nil
}








func Get_User_Location(w http.ResponseWriter, r *http.Request){


}



