package location

import (
	"context"
	"fmt"
	"log"
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Freel.com/freel_api/mongo"
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

/* Also returns self */
func All_User_In_10km(userID string) {
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

}

func Add_Location() {

	// Set up a MongoDB client and connect to the database
	client := mongo.GetMongoClient()

	// Update all documents in the "users" collection to include the "location" field
	collection := client.Database("freel").Collection("users")
	update := bson.M{
		"$set": bson.M{
			"location": bson.M{
				"type": "Point",
				"coordinates": []float64{
					-122.431297, // replace with the desired longitude
					37.773972,   // replace with the desired latitude
				},
			},
		},
	}
	result, err := collection.UpdateMany(context.Background(), bson.M{}, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated %v documents\n", result.ModifiedCount)

}

/* Pass in a location value too later */
func Update_Location(User_ID string) {

	client := mongo.GetMongoClient()

	// Find the user with the given ID
	collection := client.Database("freel").Collection("users")
	id, err := primitive.ObjectIDFromHex(User_ID) // replace with the desired user ID
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": id}
	var user User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		panic(err)
	}

	// This will be replaced with typescript code that gets users location from the browser

	user.Location = Location{
		Type: "Point",
		Coordinates: []float64{
			-100.431297, // replace with the desired longitude
			37.773972,   // replace with the desired latitude
		},
	}
	update := bson.M{
		"$set": bson.M{
			"location": user.Location,
		},
	}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Updated %v documents\n", result.ModifiedCount)
}

func Distance_Between_Two_Users(User_ID1 string, User_ID2 string) {
	// Set up a MongoDB client and connect to the database

	client := mongo.GetMongoClient()

	// Find the users with the given IDs
	collection := client.Database("freel").Collection("users")
	id1, err := primitive.ObjectIDFromHex(User_ID1) // replace with the first user's ID
	if err != nil {
		panic(err)
	}
	id2, err := primitive.ObjectIDFromHex(User_ID2) // replace with the second user's ID
	if err != nil {
		panic(err)
	}
	filter := bson.M{"_id": bson.M{"$in": []primitive.ObjectID{id1, id2}}}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(context.Background())

	// Extract the users' location coordinates
	var users []User
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	if len(users) != 2 {
		panic("Expected to find exactly 2 users")
	}

	// Calculate the distance between the users in miles
	lat1 := users[0].Location.Coordinates[1]
	lon1 := users[0].Location.Coordinates[0]
	lat2 := users[1].Location.Coordinates[1]
	lon2 := users[1].Location.Coordinates[0]
	distance := distanceInMiles(lat1, lon1, lat2, lon2)
	fmt.Printf("Distance between %s and %s: %.2f miles\n", users[0].Name, users[1].Name, distance)
}

/* Distance helper functions */

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func distanceInMiles(lat1, lon1, lat2, lon2 float64) float64 {
	// convert decimal degrees to radians
	radLat1 := deg2rad(lat1)
	radLon1 := deg2rad(lon1)
	radLat2 := deg2rad(lat2)
	radLon2 := deg2rad(lon2)

	radLat1 = radLat1 * math.Pi / 180
	radLon1 = radLon1 * math.Pi / 180
	radLat2 = radLat2 * math.Pi / 180
	radLon2 = radLon2 * math.Pi / 180

	// Haversine formula
	dLat := radLat2 - radLat1
	dLon := radLon2 - radLon1
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(radLat1)*math.Cos(radLat2)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	earthRadius := 3958.8 // miles
	distance := earthRadius * c

	return distance
}
