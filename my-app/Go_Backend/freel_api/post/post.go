package post

import (

	"context"
	"encoding/json"
	"fmt"

	"net/http"


	"Freel.com/freel_api/mongo"
	
	
	"go.mongodb.org/mongo-driver/bson/primitive"
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
    SavedPosts     []Post             `bson:"saved_post,omitempty" json:"saved_post"`
}


func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a User struct
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := mongo.GetMongoClient()

	userCollection := client.Database("freel").Collection("users")

	// Insert the user into the collection
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the inserted user's ID as JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	id := result.InsertedID.(primitive.ObjectID)
	if err := json.NewEncoder(w).Encode(id.Hex()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Create_Account(user User) {

	client := mongo.GetMongoClient()

	collection := client.Database("freel").Collection("users")

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User inserted successfully")

}

func add_info() {
	//client := mongo.GetMongoClient()

}



func add_location() {
	//client := mongo.GetMongoClient()

}


func post_picture(){
	//client := mongo.GetMongoClient()



}

func Create_Fake_Account(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Creating sample data to insert into mongo since no data was given")

    post1 := Post{
        Title: "My First Post",
        Body:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
        Tags:  []string{"programming", "photography"},
        Date:  "2022-01-01T12:00:00Z",
        Image: "https://example.com/post1.jpg",
    }

    post2 := Post{
        Title: "My Second Post",
        Body:  "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        Tags:  []string{"travel", "food"},
        Date:  "2022-01-05T12:00:00Z",
        Image: "https://example.com/post2.jpg",
    }

    test_user := User{
        Name:           "John Doe",
        Bio:            "I'm a software engineer and hobbyist photographer.",
        ProfilePicture: "https://example.com/profile.jpg",
        Posts: []Post{
            post1,
            post2,
        },
        Location: Location{
            Type:        "Point",
            Coordinates: []float64{-122.4194, 37.7749},
        },
        SavedPosts: []Post{
            post1,
        },
    }

    Create_Account(test_user)
}


/*
func Upload_Photo(w http.ResponseWriter, r *http.Request) {
	// Get the photo file from the request body
	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get the MongoDB client and the GridFS bucket for photos
	client := mongo.GetMongoClient()
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Upload the file to the bucket
	uploadStream, err := bucket.OpenUploadStream(header.Filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer uploadStream.Close()

	_, err = io.Copy(uploadStream, file)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Get the ObjectID of the uploaded photo
	photoID := uploadStream.FileID()

	// Get the user ID from the URL parameter and convert it to an ObjectID
	params := mux.Vars(r)
	userID, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Get the MongoDB client and the users collection
	userCollection := client.Database("freel").Collection("users")

	// Find the user and update their profile picture
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"profilepicture": photoID.Hex()}}
	_, err = userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Return a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Photo uploaded successfully")
}

*/

/*

func Add_Post(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a Post object
	var post Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate an ID for the post
	post.ID = primitive.NewObjectID()

	// Add the post to the "posts" collection in MongoDB
	client := mongo.GetMongoClient()
	collection := client.Database("freel").Collection("posts")
	_, err = collection.InsertOne(context.Background(), post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Add the post ID to the user's "posts" array in MongoDB
	userID := filepath.Base(r.URL.Path)
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userCollection := client.Database("freel").Collection("users")
	_, err = userCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": objectID},
		bson.M{"$push": bson.M{"posts": post.ID, "saved_posts": post}},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the image to the "photos" GridFS bucket in MongoDB
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileID := primitive.NewObjectID()
	_, err = bucket.UploadFromStreamWithID(
		fileID,
		post.Image,
		bytes.NewBuffer([]byte{}),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the post object in the "posts" collection with the file ID
	_, err = collection.UpdateOne(
		context.Background(),
		bson.M{"_id": post.ID},
		bson.M{"$set": bson.M{"image": fileID}},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the post as JSON
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
*/