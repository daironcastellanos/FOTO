package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"gorm.io/gorm"
    "github.com/joho/godotenv"
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

func GetMongoClient() *mongo.Client {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
	mongo_uri := os.Getenv("MONGODB_URI")

	

	//log.Printf("MongoDB URI: %s", mongo_uri)

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

func Get_Freel_DataBase() *mongo.Database {

	mongo_client := GetMongoClient()
	mongo_Database := mongo_client.Database("freel")

	return (mongo_Database)
}

func Get_User_Collection() *mongo.Collection {

	mongo_client := GetMongoClient()
	mongo_Database := mongo_client.Database("freel").Collection("users")

	return (mongo_Database)
}

func Get_Photo_Bucket() (*gridfs.Bucket, error) {

	client := GetMongoClient()

	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
