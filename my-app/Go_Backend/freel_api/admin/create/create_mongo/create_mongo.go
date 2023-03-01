package create_mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Freel.com/freel_api/mongo"
)



func CreateDatabase( databaseName string) error {
	// Create the database
	client := mongo.GetMongoClient()

	err := client.Database(databaseName).CreateCollection(context.Background(), "dummy")
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' created successfully\n", databaseName)

	return nil
}

func CreateCollection(dbName string, collectionName string) error {
	// Create the collection
	client := mongo.GetMongoClient()
	err := client.Database(dbName).CreateCollection(context.Background(), collectionName)
	if err != nil {
		return err
	}

	fmt.Printf("Collection '%s' created successfully in database '%s'\n", collectionName, dbName)

	return nil
}

func AddFieldToCollection(uri string, dbName string, collectionName string, fieldName string, fieldValue interface{}) error {
	// Set client options
	client := mongo.GetMongoClient()

	// Update all documents in the collection to include the new field
	update := bson.M{"$set": bson.M{fieldName: fieldValue}}
	_, err = client.Database(dbName).Collection(collectionName).UpdateMany(context.Background(), bson.M{}, update)
	if err != nil {
		return err
	}

	fmt.Printf("Field '%s' added to all documents in collection '%s' of database '%s'\n", fieldName, collectionName, dbName)

	return nil
}
