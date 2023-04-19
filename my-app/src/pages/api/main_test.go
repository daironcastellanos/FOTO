package main_test

import (

	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"


	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Location struct {
	Lat         float64   `bson:"lat" json:"lat"`
	Lng         float64   `bson:"lng" json:"lng"`
	Coordinates []float64 `bson:"coordinates" json:"coordinates"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FireID         string             `bson:"FireID" json:"FireID"`
	FullName       string             `bson:"FullName" json:"FullName"`
	Username       string             `bson:"Username" json:"Username"`
	Email          string             `bson:"Email" json:"Email"`
	Bio            string             `bson:"Bio" json:"Bio"`
	Location       Location           `bson:"Location" json:"Location"`
	DOB            string             `bson:"DOB" json:"DOB"`
	Followers      []string           `bson:"Followers" json:"Followers"`
	Following      []string           `bson:"Following" json:"Following"`
	MyPhotos       []string           `bson:"MyPhotos" json:"MyPhotos"`
	SavedPhotos    []string           `bson:"SavedPhotos" json:"SavedPhotos"`
	ProfilePicture string             `bson:"ProfilePicture" json:"ProfilePicture"`
}

func GetMongoClient() *mongo.Client {
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
	client := GetMongoClient()
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


func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getting")

	// Get the Username from the URL parameter
	params := mux.Vars(r)
	username := params["username"]

	fmt.Println("getting user", username)

	// Get a MongoDB client and collection
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Query the collection for the user with the given Username
	var user bson.M
	err := collection.FindOne(context.Background(), bson.M{"Username": username}).Decode(&user)
	if err != nil {
		log.Printf("Error finding document: %v", err)
		http.Error(w, fmt.Sprintf(`{"error": "Error finding document: %v"}`, err), http.StatusInternalServerError)
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

func TestGetUserByUsername(t *testing.T) {
	// Save the original GetMongoClient and restore it after the test

	// Create a test user with unique values
	_id,err := primitive.ObjectIDFromHex("643dbd5c00167a9b3050eee7")

	testUser := User{
		ID:             _id,
		FireID:         "1bLnAlQE7IQEnXTkwza2f4Xxn6S2",
		FullName:       "Eric deQuevedo",
		Username:       "ericdequu",
		Email:          "ericdequuevedo@gmail.com",
		Bio:            "",
		Location:       Location{Lat: 29.6156734, Lng: -82.3659168, Coordinates: []float64{29.6156734, -82.3659168}},
		DOB:            "2023-04-20",
		Followers:      []string{},
		Following:      []string{"QbKWQD6DBWQD6Pp5VXABRXZWVng1"},
		MyPhotos:       []string{"643dbd9b00167a9b3050eeee"},
		SavedPhotos:    []string{},
		ProfilePicture: "643f369f1da746fbb904ce5a",
	}

	// Create a new request with the test user's username
	req, err := http.NewRequest("GET", "/api/username/"+testUser.Username,nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the HTTP response
	w := httptest.NewRecorder()

	// Call the GetUserByUsername function with the request
	r := mux.NewRouter()
	r.HandleFunc("/api/username/{username}", GetUserByUsername)
	r.ServeHTTP(w, req)

	// Check the HTTP response status code is 200
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a User
	var returnedUser User
	if err := json.NewDecoder(w.Body).Decode(&returnedUser); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	// Check the returned user matches the test user
	if returnedUser.Username != testUser.Username || returnedUser.FullName != testUser.FullName {
		t.Errorf("handler returned wrong user: got %+v want %+v", returnedUser, testUser)
	}
}


func GetUserByFireID(w http.ResponseWriter, r *http.Request) {
	// Get the FireID from the URL parameter
	params := mux.Vars(r)
	fireID := params["fireID"]

	// Get a MongoDB client and collection
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Query the collection for the user with the given FireID
	var user bson.M
	err := collection.FindOne(context.Background(), bson.M{"FireID": fireID}).Decode(&user)
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


func TestGetUserByFireID(t *testing.T) {
	// Create a test user with unique values
	_id, err := primitive.ObjectIDFromHex("643dbd5c00167a9b3050eee7")
	if err != nil {
		t.Fatalf("Error parsing ObjectID: %v", err)
	}

	testUser := User{
		ID:             _id,
		FireID:         "1bLnAlQE7IQEnXTkwza2f4Xxn6S2",
		FullName:       "Eric deQuevedo",
		Username:       "ericdequu",
		Email:          "ericdequuevedo@gmail.com",
		Bio:            "",
		Location:       Location{Lat: 29.6156734, Lng: -82.3659168, Coordinates: []float64{29.6156734, -82.3659168}},
		DOB:            "2023-04-20",
		Followers:      []string{},
		Following:      []string{"QbKWQD6DBWQD6Pp5VXABRXZWVng1"},
		MyPhotos:       []string{"643dbd9b00167a9b3050eeee"},
		SavedPhotos:    []string{},
		ProfilePicture: "643f369f1da746fbb904ce5a",
	}

	// Create a new request with the test user's FireID
	req, err := http.NewRequest("GET", "/api/fireid/"+testUser.FireID, nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	// Create a ResponseRecorder to record the HTTP response
	w := httptest.NewRecorder()

	// Call the GetUserByFireID function with the request
	r := mux.NewRouter()
	r.HandleFunc("/api/fireid/{fireID}", GetUserByFireID)
	r.ServeHTTP(w, req)

	// Check the HTTP response status code is 200
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the response body into a User
	var returnedUser User
	if err := json.NewDecoder(w.Body).Decode(&returnedUser); err != nil {
		t.Errorf("handler returned invalid JSON: %v", err)
	}

	// Check the returned user matches the test user
	if returnedUser.FireID != testUser.FireID || returnedUser.FullName != testUser.FullName {
		t.Errorf("handler returned wrong user: got %+v want %+v", returnedUser, testUser)
	}
}
