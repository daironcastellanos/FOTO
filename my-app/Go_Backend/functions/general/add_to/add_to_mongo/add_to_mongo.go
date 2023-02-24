package _add_to_mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddToCollection(uri, databaseName, collectionName string, document interface{}) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)

	// Insert the document
	_, err = collection.InsertOne(context.Background(), document)
	if err != nil {
		return err
	}

	fmt.Println("Document added to collection successfully")

	return nil
}
