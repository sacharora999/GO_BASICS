package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers in GOLANG")

	k := 100
	var ptr *int = &k
	var ptr2 = &k

	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of 2nd pointer is ", ptr2)

	fmt.Println("Value of pointer is ", *ptr)
	fmt.Println("Value of 2nd pointer is ", *ptr2)

}
