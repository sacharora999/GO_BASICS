package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to File Handling in GO lang")
	content := "The Sample text file goes here"

	file, err := os.Create("./sample.txt")
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println(length)
	readFile()
	defer file.Close()

}

func readFile() {
	contbyte, err := ioutil.ReadFile("./sample.txt")
	checkNilError(err)
	fmt.Println(string(contbyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
