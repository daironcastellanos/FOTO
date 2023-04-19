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

	"go.mongodb.org/mongo-driver/bson"

	"Freel.com/freel_api/get"

	"github.com/joho/godotenv"

	"github.com/rs/cors"

	
	"mime/multipart"
)

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
	ProfilePicture string `bson:"ProfilePicture" json:"ProfilePicture"`
}

type Picture struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Filename string             `bson:"filename"`
	Data     []byte             `bson:"data"`
}

type ResponseMessage struct {
	Message string `json:"message"`
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


func AddProfilePictureFieldToAllUsers() {
    // Get a MongoDB client and collection
    client := GetMongoClient()
    collection := client.Database("freel").Collection("users")

    // Define the update to add the ProfilePicture field with an empty string
    update := bson.M{
        "$set": bson.M{
            "ProfilePicture": "",
        },
    }

    // Update all documents in the collection
    result, err := collection.UpdateMany(context.Background(), bson.M{}, update)
    if err != nil {
        log.Printf("Error updating documents: %v", err)
        return
    }

    log.Printf("Updated %d documents", result.ModifiedCount)
}


func uploadProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading profile picture")
	
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

	err = updateUserProfilePicture(user.ID, photoID.Hex())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Profile picture uploaded successfully")
}



func updateUserProfilePicture(userID primitive.ObjectID, photoID string) error {
	fmt.Println("Updating user's profile picture")
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"ProfilePicture": photoID}}

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

func GetProfilePicture(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fireID := params["fireID"]

	fmt.Println("Getting profile picture for user with fireID: ", fireID)

	// Get a MongoDB client and collection
	client := GetMongoClient()
	userCollection := client.Database("freel").Collection("users")

	// Find the user with the specified FireID
	var user User
	err := userCollection.FindOne(context.Background(), bson.M{"FireID": fireID}).Decode(&user)
	if err != nil {
		log.Printf("Error finding user: %v", err)
		http.Error(w, "Error finding user", http.StatusInternalServerError)
		return
	}

	// Get the user's profile picture ID
	photoID := user.ProfilePicture

	// Convert the photoID string to an ObjectID
    objectId, err := primitive.ObjectIDFromHex(photoID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

	fmt.Println("Profile picture ID: ", objectId)

	// Find the picture with the specified ID
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

func AddFollower(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fireID := params["fireID"]
	followerID := params["followerID"]

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireID}
	update := bson.M{"$push": bson.M{"Followers": followerID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Follower added successfully")
}

func RemoveFollower(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fireID := params["fireID"]
	followerID := params["followerID"]

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireID}
	update := bson.M{"$pull": bson.M{"Followers": followerID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Follower removed successfully")
}

func AddFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fireID := params["fireID"]
	followingID := params["followingID"]

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireID}
	update := bson.M{"$push": bson.M{"Following": followingID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Following added successfully")
}

func RemoveFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fireID := params["fireID"]
	followingID := params["followingID"]

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireID}
	update := bson.M{"$pull": bson.M{"Following": followingID}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Following removed successfully")
}

func addPhotoToSaved(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fireId := vars["fireId"]

	decoder := json.NewDecoder(r.Body)
    var photoId string
    err := decoder.Decode(&photoId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseMessage{"Error saving photo."})
		return
	}

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireId}
	update := bson.M{"$push": bson.M{"SavedPhotos": photoId}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Photo with ID '%s' saved successfully for user with FireID '%s'.\n", photoId, fireId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseMessage{"Photo saved successfully."})
}

func followUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fireId := vars["fireId"]


	decoder := json.NewDecoder(r.Body)
    var targetFireId string
    err := decoder.Decode(&targetFireId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseMessage{"Error following user."})
		return
	}

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireId}
	update := bson.M{"$push": bson.M{"Following": targetFireId}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filter = bson.M{"FireID": targetFireId}
	update = bson.M{"$push": bson.M{"Followers": fireId}}
	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User with FireID '%s' followed user with FireID '%s'.\n", fireId, targetFireId)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ResponseMessage{"User followed successfully."})
}

func removePhotoFromSaved(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fireId := vars["fireId"]

	decoder := json.NewDecoder(r.Body)
    var photoId string
    err := decoder.Decode(&photoId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseMessage{"Error saving photo."})
		return
	}

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireId}
    update := bson.M{"$pull": bson.M{"SavedPhotos": photoId}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Photo with ID '%s' removed successfully for user with FireID '%s'.\n", photoId, fireId)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(ResponseMessage{"Photo removed successfully."})
}

func unfollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	fireId := vars["fireId"]


	decoder := json.NewDecoder(r.Body)
    var targetFireId string
    err := decoder.Decode(&targetFireId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseMessage{"Error following user."})
		return
	}

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	filter := bson.M{"FireID": fireId}
    update := bson.M{"$pull": bson.M{"Following": targetFireId}}
    _, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	filter = bson.M{"FireID": targetFireId}
    update = bson.M{"$pull": bson.M{"Followers": fireId}}
    _, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User with FireID '%s' unfollowed user with FireID '%s'.\n", fireId, targetFireId)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(ResponseMessage{"User unfollowed successfully."})
}

func updateUserLocation(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    vars := mux.Vars(r)
    fireId := vars["fireId"]

    var location Location
    err := json.NewDecoder(r.Body).Decode(&location)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(ResponseMessage{"Error updating user location."})
        return
    }

    client := GetMongoClient()
    collection := client.Database("freel").Collection("users")

    filter := bson.M{"FireID": fireId}
    update := bson.M{"$set": bson.M{"Location": location}}

    _, err = collection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    fmt.Printf("User with FireID '%s' location updated successfully.\n", fireId)
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(ResponseMessage{"User location updated successfully."})
}


func Get_Nearby_users(w http.ResponseWriter, r *http.Request) {
	fireID := mux.Vars(r)["fireID"]

	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")

	// Create the 2dsphere index on the Location field
	indexModel := mongo.IndexModel{
		Keys: bson.M{
			"Location": "2dsphere",
		},
	}
	_, err := collection.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		log.Println(err)
		return
	}

	var user User
	err = collection.FindOne(context.Background(), bson.M{"FireID": fireID}).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}

	filter := bson.M{
		"Location": bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": user.Location.Coordinates,
				},
				"$maxDistance": 10000, // 10 km radius
			},
		},
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return
	}
	defer cursor.Close(context.Background())

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

func Freel_Api() {

	r := mux.NewRouter()

	// Photo endpoints
	r.HandleFunc("/api/photos/{photoId}", Get_Photo).Methods("GET")
	r.HandleFunc("/api/upload/{fireID}", uploadPhotoHandler).Methods("POST")
	r.HandleFunc("/api/random_pic/get", get.GetRandomImage).Methods("GET")

	// User endpoints
	r.HandleFunc("/api/create/user", CreateUser).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/get", get.GetUserByFireID).Methods("GET")
	r.HandleFunc("/api/username/{username}/get", get.GetUserByUsername).Methods("GET")
	r.HandleFunc("/api/users/get", get.GetAllUsers).Methods("GET")

	r.HandleFunc("/api/users/{fireID}/update_bio", Update_Bio).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/uploadProfilePicture", uploadProfilePictureHandler).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/getProfilePicture", GetProfilePicture).Methods("GET")
	r.HandleFunc("/api/users/{fireId}/updateLocation", updateUserLocation).Methods("POST")
	r.HandleFunc("/api/users/{fireId}/savePhoto", addPhotoToSaved).Methods("POST")
	r.HandleFunc("/api/users/{fireId}/removePhoto", removePhotoFromSaved).Methods("POST")
	r.HandleFunc("/api/users/{fireId}/follow", followUser).Methods("POST")
	r.HandleFunc("/api/users/{fireId}/unfollow", unfollowUser).Methods("POST")

	// Follower and Following endpoints
	r.HandleFunc("/api/users/{fireID}/addFollower/{followerID}", AddFollower).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/removeFollower/{followerID}", RemoveFollower).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/addFollowing/{followingID}", AddFollowing).Methods("POST")
	r.HandleFunc("/api/users/{fireID}/removeFollowing/{followingID}", RemoveFollowing).Methods("POST")

	// Nearby users endpoint
	r.HandleFunc("/api/nearby_users/{fireID}", Get_Nearby_users).Methods("GET")
	
	// Start the server
	log.Println("Starting server on :8080")

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":8080",handler))
}


func main() {
	Freel_Api()
}
