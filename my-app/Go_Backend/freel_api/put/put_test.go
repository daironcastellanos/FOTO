package put_test

import (
	"testing"

	"Freel.com/freel_api/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Like struct {
    Username string `bson:"username,omitempty" json:"username"`
    Date     string `bson:"date,omitempty" json:"date"`
}

type Comment struct {
    Username string `bson:"username,omitempty" json:"username"`
    Date     string `bson:"date,omitempty" json:"date"`
    Comment  string `bson:"comment,omitempty" json:"comment"`
}

type Post struct {
    gorm.Model
    Title    string     `json:"title"`
    Body     string     `json:"body"`
    Tags     []string   `json:"tags"`
    Date     string     `json:"date"`
    Image    string     `json:"image"`
    Likes    []Like     `bson:"likes,omitempty" json:"likes"`
    Comments []Comment  `bson:"comments,omitempty" json:"comments"`
}

type Location struct {
	Type        string    `bson:"type,omitempty" json:"type"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates"`
}

type User struct {
    ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Name           string             `bson:"name,omitempty" json:"name"`
    Bio            string             `bson:"bio,omitempty" json:"bio"`
    ProfilePicture string             `bson:"profilepicture,omitempty" json:"profilepicture"`
    Posts          []Post             `bson:"posts,omitempty" json:"posts"`
    Location       Location           `bson:"location,omitempty" json:"location"`
    SavedPosts     []Post             `bson:"saved_post,omitempty" json:"saved_post"`
}



