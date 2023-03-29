package main

import (
	"fmt"
	"testing"

	"Freel.com/freel_api/get"
	"Freel.com/freel_api/mongo"
	"Freel.com/freel_api/post"
)

func TestAll(t *testing.T) {


	fmt.Println("Starting Test")
	fmt.Println("Creating Fake Account ")
	post.Test_Create_Fake_Account(t)
	fmt.Println()	
	
	fmt.Println("Getting Fake Account ")
	get.TestGet_User(t)
	fmt.Println()

<<<<<<< HEAD


	fmt.Println("Deleting Fake Account ")
	delete.Test_DeleteUser(t)
	fmt.Println()


=======
	fmt.Println("Testing mongo file and random function")
	mongo.TestGetRandomImage(t)
	fmt.Println()	
>>>>>>> 10dd0d00853a2f83ed915b301f65612e7645b1b2
}
