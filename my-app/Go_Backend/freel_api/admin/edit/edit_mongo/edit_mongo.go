package edit_mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"Freel.com/freel_api/mongo"
)

func UpdateDocuments(dbName string, collectionName string, filter bson.M, update bson.M) error {
	// Update the documents that match the filter
	client := mongo.GetMongoClient()
	result, err := client.Database(dbName).Collection(collectionName).UpdateMany(context.Background(), filter, update)
	if err != nil {
		return err
	}

	fmt.Printf("%d documents updated in collection '%s' of database '%s'\n", result.ModifiedCount, collectionName, dbName)

	return nil
}
