package main

import "fmt"

func main() {
	fmt.Println("Welcome to more updates on Slices in GOLAN")
	vegSlice := []int{100, 132, 123, 4, 3, 1, 11, 12, 3322, 322, 42}
	fmt.Print(vegSlice)

	vegSlice = append(vegSlice[:2], vegSlice[3:]...)
	fmt.Print(vegSlice)
}
