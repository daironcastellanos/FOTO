package location

import (
	"context"
	"encoding/json"

	"log"
	"math"
	"net/http"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name           string             `json:"name"`
	Bio            string             `json:"bio"`
	ProfilePicture string             `json:"profilepicture"`
	Posts          []Post             `json:"posts"`
	Location       Location           `json:"location"`
}

type Post struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
}

type Location struct {
	Type        string    `json:"type" bson:"type"`
	Coordinates []float64 `json:"coordinates" bson:"coordinates"`
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
	// Extract the user ID from the URL path
	userID := mux.Vars(r)["id"]

	// Set up a MongoDB client and connect to the database
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Find the user with the specified ID
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		log.Println(err)
		return
	}
	var user User
	err = collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
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
				"$maxDistance": 10000,
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