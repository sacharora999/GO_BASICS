package main

import "fmt"

func main() {
	fmt.Println("Defer in golan")

	defer fmt.Println("WORLD")
	defer fmt.Println("ONE")
	defer fmt.Println("TWO")
	fmt.Println("HELLO")
	deferMe()

}

func deferMe() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

// HELLO
// 4
// 3
// 2
// 1
// 0
// TWO
// ONE
// WORLD
