package edit_mongo

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

func UpdateDocuments(client *mongo.Client, dbName string, collectionName string, filter bson.M, update bson.M) error {
	// Update the documents that match the filter
	result, err := client.Database(dbName).Collection(collectionName).UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("%d documents updated in collection '%s' of database '%s'\n", result.ModifiedCount, collectionName, dbName)

	return nil
}
