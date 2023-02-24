package delete_mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(uri string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Check the connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func DeleteDatabase(client *mongo.Client, databaseName string) error {
	// Delete the database
	err := client.Database(databaseName).Drop(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' deleted successfully\n", databaseName)

	return nil
}

func DeleteCollection(client *mongo.Client, dbName string, collectionName string) error {
	// Delete the collection
	err := client.Database(dbName).Collection(collectionName).Drop(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Collection '%s' deleted successfully from database '%s'\n", collectionName, dbName)

	return nil
}

func DeleteDocuments(client *mongo.Client, dbName string, collectionName string, filter bson.M) error {
	// Delete the documents that match the filter
	result, err := client.Database(dbName).Collection(collectionName).DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("%d documents deleted from collection '%s' of database '%s'\n", result.DeletedCount, collectionName, dbName)

	return nil
}
