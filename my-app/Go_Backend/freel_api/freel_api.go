package freel_api 


import(
	"context"
    "encoding/json"
    "log"
    "net/http"
	"os"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"


	"Freel.com/freel_api/get"
	"Freel.com/freel_api/post"
	"Freel.com/freel_api/put"
	"Freel.com/freel_api/delete"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
	Date  string   `json:"date"`
	Image string   `json:"image"`
}


type User struct {
	gorm.Model
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilepicture"`
	Posts          []Post `json:"posts"`
}



var client *mongo.Client


func freel_api() {

	
    // Set up the Gorilla Mux router and define your API routes
    r := mux.NewRouter()

    r.HandleFunc("/api/users", get_users).Methods("GET")
    r.HandleFunc("/api/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/api/users", createUser).Methods("POST")
    r.HandleFunc("/api/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/api/users/{id}", deleteUser).Methods("DELETE")

    // Start the server
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

func getUser(){

}