package get_mongo

import (
	"context"
	"encoding/json"
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

func GetDatabaseNames(client *mongo.Client) ([]string, error) {
	// Get the list of database names
	databases, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	return databases, nil
}

func GetCollectionNames(client *mongo.Client, dbName string) ([]string, error) {
	// Get the list of collection names in the specified database
	collections, err := client.Database(dbName).ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	return collections, nil
}

func GetDocuments(client *mongo.Client, dbName string, collectionName string, filter bson.M) ([]bson.M, error) {
	// Get the documents that match the filter
	cur, err := client.Database(dbName).Collection(collectionName).Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var results []bson.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil
}

func ReadCollectionToJson(client *mongo.Client, dbName string, collectionName string) ([]byte, error) {
	// Get the collection
	collection := client.Database(dbName).Collection(collectionName)

	// Find all documents in the collection
	cur, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Loop through the documents and add them to a slice
	var results []bson.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	// Convert the slice to JSON
	jsonBytes, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%d documents found in collection '%s' of database '%s'\n", len(results), collectionName, dbName)

	return jsonBytes, nil
}

func QueryMongoDB(uri, databaseName, collectionName string, query bson.M) ([]bson.M, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(context.Background())

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Get the collection
	collection := client.Database(databaseName).Collection(collectionName)

	// Find documents that match the query
	cur, err := collection.Find(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	// Loop through the documents and add them to a slice
	var results []bson.M
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	fmt.Printf("%d documents found in collection '%s' of database '%s'\n", len(results), collectionName, databaseName)

	return results, nil
}
