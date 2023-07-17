package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loop Break and GOTO in GOLANG")

	//days := []string{"Sun", "Mon", "Tue", "Wed", "Thu"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Println(index, " : ", day)
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping to code online")
}
