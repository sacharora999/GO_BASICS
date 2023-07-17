package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Slices in GOLANG")

	var fruitList = []int{1, 2, 3, 4, 5}
	fmt.Printf("Type of fruit is %T\n", fruitList)

	fruitList = append(fruitList, 100)
	fmt.Println("Fruits are  : ", fruitList)

	fruitList = append(fruitList[1:6])
	fmt.Println("New Fruits are  : ", fruitList)

	mySlice := make([]int, 4)
	mySlice[0] = 102
	mySlice[1] = 902
	mySlice[2] = 142
	mySlice[3] = 322
	//	mySlice[4] = 686

	mySlice = append(mySlice, 1101, 3737)

	fmt.Println("Make Slice is   : ", mySlice)
	fmt.Println(sort.IntsAreSorted(mySlice))
	sort.Ints(mySlice)

	fmt.Println("Make Slice is   : ", mySlice)
	sort.IntsAreSorted(mySlice)
	fmt.Println(sort.IntsAreSorted(mySlice))

}
