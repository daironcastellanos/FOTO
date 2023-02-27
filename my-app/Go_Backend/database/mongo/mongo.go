package mongo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* declares Structs for data will be moved to seperate file when this is cleaned up */

type Post struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
}

type User struct {
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilepicture"`
	Posts          []Post `json:"posts"`
}

/* This is a function that can be used to initilize the mongo client for future functions*/
/* Returns mongo 🧞‍♂️ client */
func InitiateMongoClient(URI_ string) *mongo.Client {
	var err error
	var client *mongo.Client
	uri := URI_
	opts := options.Client()
	opts.ApplyURI(uri)
	opts.SetMaxPoolSize(5)
	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		fmt.Println(err.Error())
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Pinged mongo")
	}

	return client
}

/* This Function accesses the Database freel and conects to the user collection successfully */
func Get_User_Collection() {

	url := "https://us-east-1.aws.data.mongodb-api.com/app/data-vufcj/endpoint/data/v1/action/findOne"
	method := "POST"

	payload := strings.NewReader(`{
			"collection":"users",
			"database":"freel",
			"dataSource":"Cluster0",
			"projection": {"_id": 1}
		}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", "OoktA3ZOGR5D2DHJZhPhi4wupxtKB5YYcyhjNwdZt8kJGtHwrnxxR3RgypKimp1v")
	req.Header.Add("Accept", "application/ejson")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/* This Function accesses the Database freel and conects to the Content collection successfully currently no data🤬*/
func Get_Content_Collection() {

	url := "https://us-east-1.aws.data.mongodb-api.com/app/data-vufcj/endpoint/data/v1/action/findOne"
	method := "POST"

	payload := strings.NewReader(`{
			"collection":"Content",
			"database":"freel",
			"dataSource":"Cluster0",
			"projection": {"_id": 1}
		}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Access-Control-Request-Headers", "*")
	req.Header.Add("api-key", "OoktA3ZOGR5D2DHJZhPhi4wupxtKB5YYcyhjNwdZt8kJGtHwrnxxR3RgypKimp1v")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

/* Function to upload pic to database freel and photo bucket "photos" 👌 */
func Upload_Pic(URI_ string) {
	// Set client options
	clientOptions := options.Client().ApplyURI(URI_)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// Open the file

	file, err := ioutil.ReadFile("database/test_photo/photo1.jpg")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Create a GridFS bucket
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		fmt.Println("Error creating GridFS bucket:", err)
		return
	}

	// Upload the file
	uploadStream, err := bucket.OpenUploadStream("photo.jpg")
	if err != nil {
		fmt.Println("Error opening upload stream:", err)
		return
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(file)
	if err != nil {
		fmt.Println("Error writing file to stream:", err)
		return
	}

	fmt.Println("File uploaded successfully.")
}

func Get_All_Photos(URI_ string) ([]byte, error) {
	// Set up a MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI(URI_)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Open a GridFS bucket named "photos"
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		return nil, err
	}

	// Retrieve a photo by filename
	filename := "photo.jpg"
	//FileID := "63f55fb60620dbf31b5b31ba"
	//file, err := bucket.OpenUploadStreamWithID(FileID)
	file, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the photo data into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Print the length of the data to confirm that we have read the file
	fmt.Printf("Read %d bytes\n", len(data))

	return data, nil
}

/* function to Download a photo currently referenced through the name of photo but should be photoid 🧞‍♂️ */
func Download_Photo(URI_ string) {
	// Set up a MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI(URI_)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Open a GridFS bucket named "photos"
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Download a photo by filename
	filename := "photo.jpg"
	downloadStream, err := bucket.OpenDownloadStreamByName(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer downloadStream.Close()

	// Read the photo data into a byte slice
	data, err := ioutil.ReadAll(downloadStream)
	if err != nil {
		log.Fatal(err)
	}

	// Write the photo data to a file
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Photo downloaded and saved to %s\n", filename)
}

func Delete_Photo(Photo_ID string, URI_ string) error {
	// Set up a MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI(URI_)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Open a GridFS bucket named "photos"
	bucket, err := gridfs.NewBucket(
		client.Database("freel"),
		options.GridFSBucket().SetName("photos"),
	)
	if err != nil {
		return err
	}

	// Delete the photo with the specified ObjectID
	objectID, err := primitive.ObjectIDFromHex(Photo_ID)
	if err != nil {
		return err
	}

	err = bucket.Delete(objectID)
	if err != nil {
		return err
	}

	fmt.Printf("Photo with ID %s deleted successfully\n", Photo_ID)
	return nil
}


/*This function takes in user mongo Id and returns all user data */
func Find_User_By_ID(userID string, URI_ string) (*User, error) {
	// Set up a MongoDB client and connect to the database
	clientOptions := options.Client().ApplyURI(URI_)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Find the user with the specified ObjectID and return their data
	objectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	collection := client.Database("freel").Collection("users")
	filter := bson.M{"_id": objectID}
	var user User
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	/* will delete print after test is done */
	fmt.Printf("User data:\n%+v\n", user)

	return &user, nil
}
