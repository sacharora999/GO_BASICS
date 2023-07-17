package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.linkedin.com/in/sachin-arora-blockchain-developer/"

func main() {
	fmt.Println("Welcome to Golan")
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n ", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(databytes))

}
