package freel_api

import (
	"log"
	"net/http"

	"Freel.com/freel_api/delete"
	"Freel.com/freel_api/get"
	"Freel.com/freel_api/post"
	"Freel.com/freel_api/put"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Freel.com/freel_api/admin/location"

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
}

func Freel_Api() {

	// Set up the Gorilla Mux router and define your API routes
	r := mux.NewRouter()

	/* Serves application */
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("../public")))


	/* Gets all users or specific user with unique id */
	r.HandleFunc("/api/users/get", get.Get_Users).Methods("GET")
	r.HandleFunc("/api/users/{id}/get", get.Get_User).Methods("GET")

	r.HandleFunc("/api/users/{id}/photos/get", get.GetUserPosts_Help).Methods("GET")


	//r.HandleFunc("/api/users/{id}/photos/posts/new", post.Upload_Photo).Methods("GET")


	/* create fake account Or create real account with post */
	r.HandleFunc("/api/user/create", post.Create_Fake_Account).Methods("GET")
	r.HandleFunc("/api/users/create_user/post", post.CreateUser).Methods("POST")


	/* update bio or update entire PRofile */
	r.HandleFunc("/api/users/{id}/update/bio", put.Update_Bio).Methods("PUT")
	r.HandleFunc("/api/users/{id}/update/profile/put", put.UpdateUser).Methods("PUT")


	/* deletes a specific user */
	r.HandleFunc("/api/users/{id}/delete", delete.Delete_User).Methods("DELETE")


	/* deletes a specific photo */
	r.HandleFunc("/api/photo/{id}/delete", delete.Delete_Pic).Methods("DELETE")


	/* gets all photos*/
	r.HandleFunc("/api/photos", get.Serve_Pics).Methods("GET")

	
	/* Gets photo by ID */
	//r.HandleFunc("/api/photos/{id}", get.Get_Photo).Methods("GET")


	/* Gets nearby users */
	r.HandleFunc("/api/nearby_users/{id}", location.Get_Nearby_users).Methods("GET")

	// Start the server
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
