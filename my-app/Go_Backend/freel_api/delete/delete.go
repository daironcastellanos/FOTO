package delete

import (
	
	
	"gorm.io/gorm"
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


type User struct {
	gorm.Model
	Name           string `json:"name"`
	Bio            string `json:"bio"`
	ProfilePicture string `json:"profilepicture"`
	Posts          []Post `json:"posts"`
}



func delete_account(){
	client := mongo.GetMongoClient()




}



func delete_profile_info(){
	


}



func delete_pic(){





}




func delete_user_location(){




}