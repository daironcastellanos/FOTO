package main

import (
	"bytes"
	"context"
	
	
	"io"
	"mime/multipart"
	"net/http"
	
	"strings"
	"testing"

	//"Freel/freel_api/get"
	
	"go.mongodb.org/mongo-driver/bson"
	
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateBio updates the user's bio in the database.
func UpdateBio(FireID string, Bio string) (*mongo.UpdateResult, error) {
	client := GetMongoClient()
	collection := client.Database("freel").Collection("users")
	filter := bson.M{"FireID": FireID}

	update := bson.M{"$set": bson.M{"Bio": Bio}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)

	return updateResult, err
}

func Update_Bio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	FireID := params["fireID"]

	var requestBody map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	Bio := requestBody["bio"]

	updateResult, err := UpdateBio(FireID, Bio)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("User with fireID %s updated successfully. Modified count: %d\n", FireID, updateResult.ModifiedCount)
	fmt.Fprintf(w, "User with fireID %s updated successfully\n", FireID)
}

func TestUpdateBio(t *testing.T) {
	var err error
	var client *mongo.Client
	var collection *mongo.Collection
	var ctx = context.Background()

	// Prepare test data
	FireID := "test_fire_id"
	Bio := "Test bio"
	user := User{
		FireID: FireID,
		Bio:    "Original bio",
	}

	if client, err = GetMongoClient(); err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect(ctx)

	collection = client.Database("freel").Collection("users")
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		t.Fatal(err)
	}

	// Test UpdateBio function
	updateResult, err := UpdateBio(FireID, Bio)
	if err != nil {
		t.Fatal(err)
	}
	if updateResult.ModifiedCount != 1 {
		t.Errorf("Expected modified count to be 1, but got %d", updateResult.ModifiedCount)
	}

	// Check if the bio was updated
	var updatedUser User
	filter := bson.M{"FireID": FireID}
	err = collection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		t.Fatal(err)
	}

	if updatedUser.Bio != Bio {
		t.Errorf("Expected Bio to be %s, but got %s", Bio, updatedUser.Bio)
	}

	// Clean up the test data
	_, err = collection.DeleteOne(ctx, bson.M{"FireID": FireID})
	if err != nil {
		t.Fatal(err)
	}
}
