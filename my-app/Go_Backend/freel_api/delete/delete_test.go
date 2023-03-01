package delete_test

import (
	"context"
	"testing"

	"Freel.com/freel_api/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func DeleteUser(t *testing.T) error {

	id := "63f5687adcf9b9a96ad516a4"
    // Convert the ID string to a MongoDB ObjectID
    objectID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return err
    }

    client := mongo.GetMongoClient()
    

    // Get the "users" collection from the "test" database
    collection := client.Database("test").Collection("users")

    // Delete the user with the given ID
    _, err = collection.DeleteOne(context.Background(), bson.M{"_id": objectID})
    if err != nil {
        return err
    }

    // Close the connection
    err = client.Disconnect(context.Background())
    if err != nil {
		return err
    }

    return nil
}