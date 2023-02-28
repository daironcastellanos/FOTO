package delete

import (
	"fmt"
	"net/http"

	"Freel.com/freel_api/mongo"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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



func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameter
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
	http.Error(w, "Invalid ID", http.StatusBadRequest)
	return
	}

	fmt.Println(id);
}



func Delete_Pic(Photo_ID string){

	// Open a GridFS bucket named "photos"
	bucket,err := mongo.Get_Photo_Bucket()

	// Delete the photo with the specified ObjectID
	objectID, err := primitive.ObjectIDFromHex(Photo_ID)
	if err != nil {
		//return err
	}

	err = bucket.Delete(objectID)
	if err != nil {
		//return err
	}

	fmt.Printf("Photo with ID %s deleted successfully\n", Photo_ID)
	

}




func Delete_User_Location(){




}