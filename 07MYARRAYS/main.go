package main

import "fmt"

func main() {
	fmt.Println("Welcome to Arrays in GOLANG")
	var myarr [4]string

	myarr[0] = "Applet"
	myarr[1] = "Mango"
	//	myarr[2] = "Banana"
	myarr[3] = "Orange"

	fmt.Println("Fruit list is : ", myarr)      // 3rd elmn is blank
	fmt.Println("Fruit list is : ", len(myarr)) //4

	var vegList = [20]string{"Patato", "Tomato", "Beans"}
	fmt.Println("Veg list is : ", vegList)
}
