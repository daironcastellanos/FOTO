package create_mongo

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

func CreateDatabase(client *mongo.Client, databaseName string) error {
	// Create the database
	err := client.Database(databaseName).CreateCollection(context.Background(), "dummy")
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' created successfully\n", databaseName)

	return nil
}

func CreateCollection(client *mongo.Client, dbName string, collectionName string) error {
	// Create the collection
	err := client.Database(dbName).CreateCollection(context.Background(), collectionName)
	if err != nil {
		return err
	}

	fmt.Printf("Collection '%s' created successfully in database '%s'\n", collectionName, dbName)

	return nil
}

func AddFieldToCollection(uri string, dbName string, collectionName string, fieldName string, fieldValue interface{}) error {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	// Update all documents in the collection to include the new field
	update := bson.M{"$set": bson.M{fieldName: fieldValue}}
	_, err = client.Database(dbName).Collection(collectionName).UpdateMany(context.Background(), bson.M{}, update)
	if err != nil {
		return err
	}

	fmt.Printf("Field '%s' added to all documents in collection '%s' of database '%s'\n", fieldName, collectionName, dbName)

	return nil
}
