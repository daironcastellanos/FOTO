package freel_api

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"

	"Freel.com/freel_api/get"
	"Freel.com/freel_api/put"
	"Freel.com/freel_api/admin/location"
	"Freel.com/freel_api/mongo"
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

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"http://192.168.0.178:3000"})


	/* Serves application */
	//r.PathPrefix("/").Handler(http.FileServer(http.Dir("../.next")))
	
	/* These work and have been added to test page */
	r.HandleFunc("/api/users/get", get.GetAllUsers).Methods("GET")
	r.HandleFunc("/api/users/{id}/get", get.GetUserById).Methods("GET")
	
	r.HandleFunc("/api/users/{id}/post/get", get.GetUserById_post).Methods("GET")
	r.HandleFunc("/api/random_pic/get", mongo.GetRandomImage).Methods("GET")
	r.HandleFunc("/api/photos/{id}", get.Get_Photo).Methods("GET")
	
	/* These dont really work */
	r.HandleFunc("/api/users/{id}/update/profile/put", put.Post_Pic).Methods("POST")
	r.HandleFunc("/api/users/{id}/{bio}", put.Update_Bio).Methods("POST")

	/* Gets nearby users */
	r.HandleFunc("/api/nearby_users/{id}", location.Get_Nearby_users).Methods("GET")

	
	/* gets all photos*/
	//r.HandleFunc("/api/photos", get.Serve_Pics).Methods("GET")	
	
	/* Functions we should get working */
	
	/* deletes a specific user */
	//r.HandleFunc("/api/users/{id}/delete", delete.Delete_User).Methods("DELETE")

	/* deletes a specific photo */
	//r.HandleFunc("/api/photo/{id}/delete", delete.Delete_Pic).Methods("DELETE")

	// Start the server
	log.Println("Starting server on :8080")

	log.Fatal(http.ListenAndServe(":8080",handlers.CORS(headers, methods, origins)(r)))
}

