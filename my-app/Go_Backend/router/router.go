package router

import (
	/* go imports */
	"fmt"
	"log"
	"net/http"
	"os"

	/* My Modules */
	"Freel.com/database/data/user"
	"Freel.com/database/mongo"

	/*Additional go packages */
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func serve_pics(w http.ResponseWriter, r *http.Request) {
	// Load the MONGO_URI from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")

	// Get the image data (e.g. from a file or database)
	data, err := mongo.Get_All_Photos(mongoURI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to indicate that this is an image
	w.Header().Set("Content-Type", "image/jpeg")

	// Write the image data to the response
	if _, err := w.Write(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func serve_user_id(w http.ResponseWriter, r *http.Request) {
	test_user_input_id := "63f565f8df6db2c34aed8997"
	//test_user_input_id2 := "63f5687adcf9b9a96ad516a4"
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")

	usr_data, err := mongo.Find_User_By_ID(test_user_input_id, mongoURI)
	if err != nil {
		return
	}
	fmt.Printf("User data:\n%+v\n", usr_data)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Backend Home Page")
}

func update_bio(w http.ResponseWriter, r *http.Request) {
	/* needs user ID and new BIO */
	user.Update_User_Bio(string("new_bio"), string("USER_ID"), string("NEW BIO"))
}

func InitializeRouter() {
	/* Creates new router object */
	r := mux.NewRouter()

	/*Handle func adds a path and a function to the router*/
	/* the function referenced by Handled Func needs the parameters (w http.ResponseWriter, r *http.Request) */

	/* declares basic default page */
	r.HandleFunc("/", home).Methods("GET")

	/* Un comment Rest API path when route have been tested */

	// MONGO COLLECTIONS
	//r.HandleFunc("/mongo/users", mongo.Get_User_Collection).Methods("GET")
	//r.HandleFunc("/mongo/content", mongo.Get_Content_Collection).Methods("GET")

	// MONGO PHOTO FUNCTIONS
	//r.HandleFunc("/mongo/photos/upload", mongo.Upload_Pic).Methods("POST")
	r.HandleFunc("/mongo/photos/all", serve_pics).Methods("GET")
	//r.HandleFunc("/mongo/photos/{id}", mongo.Download_Photo).Methods("GET")
	//r.HandleFunc("/mongo/photos/{id}", mongo.Delete_Photo).Methods("DELETE")

	// USER FUNCTIONS
	r.HandleFunc("/user/{id}", serve_user_id).Methods("GET")
	//r.HandleFunc("/user/create", user.Create_Test_User).Methods("POST")
	r.HandleFunc("/user/{id}/bio", update_bio).Methods("PUT")

	// LOCATION FUNCTIONS
	//r.HandleFunc("/location/add", location.Add_Location).Methods("POST")
	//r.HandleFunc("/location/{id}/update", location.Update_Location).Methods("PUT")
	//r.HandleFunc("/location/{id}/users", location.All_User_In_10km).Methods("GET")
	//r.HandleFunc("/location/{id1}/distance/{id2}", location.Distance_Between_Two_Users).Methods("GET")

	/* This Function starts the Router at (localhost:8081) */
	/* All the Functions and paths added with HandleFunc can be accessed at (localhost:8081/path) */
	/* Log function will be used to print when the Router is closed */
	log.Fatal(http.ListenAndServe(":8081", r))
}
