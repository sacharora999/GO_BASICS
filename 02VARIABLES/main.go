package main

import "fmt"

// https://go.dev/ref/spec
var jwtToken uint64 = 309393

//jwtToken := 38383 (NOT Allowed)
const LoginToken = "My Token"

func main() {
	var username string = "Sachin"
	fmt.Println(username)
	fmt.Printf("username Type is : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn Type is : %T \n", isLoggedIn)

	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("isLoggedIn Type is : %T \n", smallInt)

	var smallVar int
	fmt.Println(smallVar)
	fmt.Printf("Type is : %T \n", smallVar)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("Type is : %T \n", defaultString)

	var defaultBool bool
	fmt.Println(defaultBool)
	fmt.Printf("Type is : %T \n", defaultBool)

	//Implicit Types
	var website = "http://www.google.com"
	fmt.Printf("Type is : %T \n", website)

	//No Var Style

	numOfVars := 50
	fmt.Printf("Type is : %T \n", numOfVars)

	fmt.Printf("Type is : %T \n", LoginToken)

}
