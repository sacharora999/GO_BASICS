package main

import "fmt"

func main() {
	fmt.Println("Welcome to Functions in GOLANG")
	result := addr(3, 8)
	fmt.Println("Result of addition is", result)

	result1, _ := proAddr(3, 8, 9, 20)
	fmt.Println("Result of addition is", result1)
}

func addr(val1 int, val2 int) int {
	return val1 + val2
}

func proAddr(values ...int) (int, string) {

	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Hello world"
}
