package delete_mongo

import (
	"context"
	"fmt"
	

	"go.mongodb.org/mongo-driver/bson"
	
	

	"Freel.com/freel_api/mongo"
)


func DeleteDatabase( databaseName string) error {
	// Delete the database
	client := mongo.GetMongoClient()
	err := client.Database(databaseName).Drop(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Database '%s' deleted successfully\n", databaseName)

	return nil
}

func DeleteCollection( dbName string, collectionName string) error {
	// Delete the 
	client := mongo.GetMongoClient()
	err := client.Database(dbName).Collection(collectionName).Drop(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Collection '%s' deleted successfully from database '%s'\n", collectionName, dbName)

	return nil
}

func DeleteDocuments( dbName string, collectionName string, filter bson.M) error {
	// Delete the documents that match the filter
	client := mongo.GetMongoClient()
	result, err := client.Database(dbName).Collection(collectionName).DeleteMany(context.Background(), filter)
	if err != nil {
		return err
	}

	fmt.Printf("%d documents deleted from collection '%s' of database '%s'\n", result.DeletedCount, collectionName, dbName)

	return nil
}
