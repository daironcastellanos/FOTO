package main

import (
	"fmt"
	"testing"

	"Freel.com/freel_api/delete"
	"Freel.com/freel_api/get"
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



	fmt.Println("Deleting Fake Account ")
	delete.Test_DeleteUser(t)
	fmt.Println()

	



	

}
