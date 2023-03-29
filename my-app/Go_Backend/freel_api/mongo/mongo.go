package mongo

import (
	"context"
	//"encoding/json"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
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
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	Tags     []string  `json:"tags"`
	Date     string    `json:"date"`
	Image    string    `json:"image"`
	Likes    []Like    `bson:"likes,omitempty" json:"likes"`
	Comments []Comment `bson:"comments,omitempty" json:"comments"`
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

func GetMongoClient_() (*mongo.Client, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("Error loading .env file: %v", err)
	}
	mongo_uri := os.Getenv("MONGODB_URI")

	// Set up a connection to MongoDB
	clientOptions := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("Error pinging MongoDB: %v", err)
	}

	return client, nil
}

func Get_Freel_DataBase() *mongo.Database {

	mongo_client := GetMongoClient()
	mongo_Database := mongo_client.Database("freel")

	return (mongo_Database)
}

func Get_User_Collection() *mongo.Collection {

	mongo_client := GetMongoClient()
	mongo_Database := mongo_client.Database("freel").Collection("test_data")

	return (mongo_Database)
}

func UploadImagesToPhotoBucket(bucketName, imageFolderPath string) {
	imageFiles, err := ioutil.ReadDir(imageFolderPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, imageFile := range imageFiles {
		if filepath.Ext(imageFile.Name()) == ".webp" {
			imagePath := filepath.Join(imageFolderPath, imageFile.Name())
			uploadImageToBucket_client(bucketName, imagePath)
		}
	}
}

func uploadImageToBucket(bucketName, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", filepath.Base(imagePath))
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(part, file)
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	uploadEndpoint := fmt.Sprintf("https://api.example.com/buckets/%s/upload", bucketName)
	request, err := http.NewRequest("POST", uploadEndpoint, body)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Body.Close()
		fmt.Println("Image uploaded:", imagePath)
	}
}

func uploadImageToBucket_client(bucketName, imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	client := GetMongoClient()
	defer client.Disconnect(context.Background())

	// Get a handle for your database
	db := client.Database(bucketName)

	// Get a handle for your collection
	coll := db.Collection("test_images")

	// Create an io.Reader from the file
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// Insert the image into the collection
	_, err = coll.InsertOne(context.Background(), bson.M{
		"filename": filepath.Base(imagePath),
		"data":     fileBytes,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Image uploaded:", imagePath)
}
type Image struct {
	ID   string `json:"_id,omitempty" bson:"_id,omitempty"`
	URL  string `json:"url,omitempty" bson:"url,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Data []byte `json:"data,omitempty" bson:"data,omitempty"`
}

func GetRandomImage(w http.ResponseWriter, r *http.Request) {
	client, err := GetMongoClient_()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	collection := client.Database("freel").Collection("images")

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
