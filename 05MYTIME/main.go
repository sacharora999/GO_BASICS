package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study of GOLANG")
	currentTime := time.Now()
	fmt.Printf("Current Time: %v\n", currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.April, 10, 23, 23, 0, 0, time.UTC)
	fmt.Print(createdDate)

}
