package _add_to_mongo

import (
	"context"
	"fmt"
	"Freel.com/freel_api/mongo"
)

func AddToCollection( databaseName, collectionName string, document interface{}) error {
	
	client := mongo.GetMongoClient()
	// Connect to Mongo

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
