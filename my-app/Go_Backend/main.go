package main

import (
	"Freel.com/freel_api"

)

func main() {

	/*Loads the .env file*/

	/* Loads the MONGO_URI_ From the .env file */

	/* User object id string used to test the function that search via specific user ID */
	//test_user_input_id := "63f565f8df6db2c34aed8997"
	//test_user_input_id2 := "63f5687adcf9b9a96ad516a4"

	/* the Functions that are commented out have been tested Make sure to add the necessary imports above to run code GO like to delete them if they arent used when saved */

	/* ROUTER*/ /* API */
	//router.InitializeRouter()

	/* Serving Files */

	/* MONGO COLLECTIONS */
	//mongo.Get_User_Collection()
	//mongo.Get_Content_Collection()

	/*Mongo Photo Functions */
	//mongo.Upload_Pic(mongoURI)
	//mongo.Get_All_Photos(mongoURI)
	//mongo.Download_Photo(mongoURI)
	//mongo.Delete_Photo(, mongoURI) /* this Test case wont work because object with specific ID has already been deleted */
	//mongo.Find_User_By_ID(test_user_input_id, mongoURI)

	/* User Functions */
	//user.Create_Test_User(mongoURI)
	//user.Update_User_Bio(test_user_input_id, "I Edited this bio with Golang 👌", mongoURI)

	/* Location Functions */
	//location.Add_Location(mongoURI)
	//location.Update_Location(mongoURI, test_user_input_id)
	//location.All_User_In_10km(mongoURI, test_user_input_id)
	//location.Distance_Between_Two_Users(mongoURI, test_user_input_id, test_user_input_id2)

	//post.Create_Fake_Account()

	freel_api.Freel_Api()

	//post.Create_Fake_Account()
}


