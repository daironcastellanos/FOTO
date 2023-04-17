package main

import (
	"fmt"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"io/ioutil"


	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"

	"Freel.com/freel_api/get"

	"github.com/joho/godotenv"

	"github.com/rs/cors"

	
	"mime/multipart"
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

type Picture struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Filename string             `bson:"filename"`
	Data     []byte             `bson:"data"`
}



func GetMongoClient() *mongo.Client {
	fmt.Println("Getting mongo client")
	
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	mongo_uri := os.Getenv("MONGODB_URI")

	// Set up a connection to MongoDB
	clientOptions := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}



	return client
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Creating User")

    var User User

    if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    client := GetMongoClient()
    userCollection := client.Database("freel").Collection("users")
    res, err := userCollection.InsertOne(context.Background(), User)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(res.InsertedID.(primitive.ObjectID).Hex()); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func Update_Bio(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters
	params := mux.Vars(r)
	FireID := params["fireID"]
  
	// Get the bio string from the request body
	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
	  http.Error(w, "Invalid request body", http.StatusBadRequest)
	  return
	}
	Bio := requestBody["bio"]
  
	fmt.Println("Updating user with fireID: ", FireID)
	fmt.Println("Updating user with bio: ", Bio)
  
	// Update the user's bio in the database
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")
	filter := bson.M{"FireID": FireID}
	
	// Find and print the user document before the update
	var foundUser bson.M
	err = collection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
	  fmt.Println("Error finding user before update:", err)
	} else {
	  fmt.Println("User document before update:", foundUser)
	}
  
	update := bson.M{"$set": bson.M{"Bio": Bio}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
	  http.Error(w, err.Error(), http.StatusInternalServerError)
	  return
	}
  
	// Find and print the user document after the update
	err = collection.FindOne(context.Background(), filter).Decode(&foundUser)
	if err != nil {
	  fmt.Println("Error finding user after update:", err)
	} else {
	  fmt.Println("User document after update:", foundUser)
	}
  
	fmt.Printf("User with fireID %s updated successfully. Modified count: %d\n", FireID, updateResult.ModifiedCount)
	fmt.Fprintf(w, "User with fireID %s updated successfully\n", FireID)
  }

func uploadPhotoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading photo")
	
	vars := mux.Vars(r)
	fireID := vars["fireID"]

	user, err := findUserByFireID(fireID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	photoID, err := uploadPhoto(file, header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = addUserPhoto(user.ID, photoID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Photo uploaded successfully")
}

func findUserByFireID(fireID string) (*User, error) {
	fmt.Println("Finding user by fireID: ", fireID)
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	var user User
	filter := bson.M{"FireID": fireID}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	return &user, err
}

func uploadPhoto(file multipart.File, filename string) (primitive.ObjectID, error) {
	fmt.Println("Uploading photo with filename: ", filename)
	client := GetMongoClient()
	collection := client.Database("freel").Collection("pictures")

	// Read the file data
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Create a new Picture struct
	picture := Picture{
		Filename: filename,
		Data:     fileData,
	}

	// Insert the picture into the "pictures" collection
	result, err := collection.InsertOne(context.Background(), picture)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Get the inserted picture's ObjectID
	photoID := result.InsertedID.(primitive.ObjectID)

	return photoID, nil
}

func addUserPhoto(userID, photoID primitive.ObjectID) error {
	fmt.Println("Adding photo to user's MyPhotos array")
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"MyPhotos": photoID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	return err
}

func Get_Photo(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Getting photo")
    params := mux.Vars(r)
    photoID := params["photoId"]

    client := GetMongoClient()

    // Convert the photoID string to an ObjectID
    objectId, err := primitive.ObjectIDFromHex(photoID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Println("ObjectID: ", objectId)

    // Find the picture in the "pictures" collection
    var picture Picture
    err = client.Database("freel").Collection("pictures").FindOne(context.Background(), bson.M{"_id": objectId}).Decode(&picture)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    fmt.Println("Picture: ", picture)

    // Set the content type header
    w.Header().Set("Content-Type", "image/jpeg")

    // Write the picture data to the response
    _, err = w.Write(picture.Data)
    if err != nil {
        fmt.Println("Error writing picture data to response:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func Freel_Api() {
	
	r := mux.NewRouter()
	r.HandleFunc("/api/photos/{photoId}", Get_Photo).Methods("GET")

	r.HandleFunc("/api/upload/{fireID}", uploadPhotoHandler).Methods("POST")
	r.HandleFunc("/api/create/user", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/get", get.GetUserByFireID).Methods("GET")
	r.HandleFunc("/api/users/{username}/get", get.GetUserByUsername).Methods("GET")
	r.HandleFunc("/api/users/get", get.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/random_pic/get", get.GetRandomImage).Methods("GET")
	r.HandleFunc("/api/users/{fireID}/update_bio", Update_Bio).Methods("Post")

	/* Serves application */
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("../.next")))

	/* These work and have been added to test page */
	
	//r.HandleFunc("/api/users/{id}/post/get", get.GetUserById_post).Methods("GET")
	
	/* Gets nearby users */
	r.HandleFunc("/api/nearby_users/{fireID}", get.Get_Nearby_users).Methods("GET")
	
	// Start the server
	log.Println("Starting server on :8080")

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080",handler))
}

func main() {
	Freel_Api()
}
