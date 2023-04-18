package get

import (
	"context"
	"encoding/json"
	"testing"
	"log"
	"net/http"

	"math"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"


	"time"
	"math/rand"
	"go.mongodb.org/mongo-driver/mongo/options"

	"fmt"
	
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

type Image struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Data     []byte             `bson:"data"`
	Filename string             `bson:"filename"`
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

func GetUserByFireID(w http.ResponseWriter, r *http.Request) {
	// Get the FireID from the URL parameter
	params := mux.Vars(r)
	fireID := params["fireID"]

	// Get a MongoDB client and collection
	client := mongo.GetMongoClient()
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


func GetRandomImage(w http.ResponseWriter, r *http.Request) {
	client := mongo.GetMongoClient()
	
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	collection := client.Database("freel").Collection("test_images")

	// Get the count of documents in the collection
	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Generate a random index
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Int63n(count)

	// Get the random image
	var randomImage Image
	opts := options.FindOne().SetSkip(randomIndex)
	err = collection.FindOne(ctx, bson.M{}, opts).Decode(&randomImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the image data to the browser
	w.Header().Set("Content-Type", "image/webp")
	w.Write(randomImage.Data)
}

func distance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // earth radius in km
	dLat := deg2rad(lat2 - lat1)
	dLon := deg2rad(lon2 - lon1)
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func Get_Nearby_users(w http.ResponseWriter, r *http.Request) {
	// Extract the user FireID from the URL path
	fireID := mux.Vars(r)["fireID"]

	// Set up a MongoDB client and connect to the database
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Find the user with the specified FireID
	var user User
	err := collection.FindOne(context.Background(), bson.M{"FireID": fireID}).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}

	

	// Find all users within a 10 km radius of the specified location
	filter := bson.M{
		"location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": user.Location.Coordinates,
				},
				"$maxDistance": 10000000,
			},
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return
	}
	defer cursor.Close(context.Background())

	// Print the results
	var nearbyUsers []User
	if err := cursor.All(context.Background(), &nearbyUsers); err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(nearbyUsers); err != nil {
		log.Println(err)
		return
	}
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


func GetUserByUsername(w http.ResponseWriter, r *http.Request) {

	fmt.Println("getting")

	// Get the Username from the URL parameter
	params := mux.Vars(r)
	username := params["username"]

	fmt.Println("getting user", username)

	// Get a MongoDB client and collection
	client := mongo.GetMongoClient()
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