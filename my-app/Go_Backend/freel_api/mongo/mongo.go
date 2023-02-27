package mongo 


import(
	"context"
    "encoding/json"
    "log"
    "net/http"
	"os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "go.mongodb.org/mongo-driver/mongo/gridfs"

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


func GetMongoClient() *mongo.Client {
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


func Get_Mongo_Data_Base(){

    mongo_client := GetMongoClient()
    mongo_Database = client.Database("freel")
    return (mongo_Database)
}

func GetMongoCollection(coll_name string){

    mongo_collection := Get_Mongo_Data_Base().Collection(coll_name);

    return(mongo_collection);
}


func Get_Photo_Bucket(){

    client := GetMongoClient()

    bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		return nil, err
	}
}

