package freel_api

import (
	"log"
	"net/http"



	"Freel.com/freel_api/delete"
	"Freel.com/freel_api/get"
	"Freel.com/freel_api/post"
	"Freel.com/freel_api/put"

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

func Freel_Api() {

	// Set up the Gorilla Mux router and define your API routes
	r := mux.NewRouter()

	r.HandleFunc("/api/users/get", get.Get_Users).Methods("GET")

	r.HandleFunc("/api/users/{id}/get", get.Get_User).Methods("GET")

	r.HandleFunc("/api/users/{id}/post", post.CreateUser).Methods("POST")

	r.HandleFunc("/api/users/{id}/put", put.UpdateUser).Methods("PUT")

	r.HandleFunc("/api/users/{id}/delete", delete.DeleteUser).Methods("DELETE")

    r.HandleFunc("/api/user/create", post.Create_Fake_Account).Methods("GET")





    r.HandleFunc("/api/photos", get.Get_Photos).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}


