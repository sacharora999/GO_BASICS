package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to Golan GET, POST Method")
	//PerformGetMethod()
	//PerformPostMethod()
	PerformFormMethod()
}

func PerformGetMethod() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(content))

	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Response String is : ", responseString.String())

}

func PerformPostMethod() {
	const myurl = "http://localhost:8000/post"
	requestBody := strings.NewReader(`
		{
			"Name": "Sachin",
			"Age": 30
		}	
	`)
	contentType := "application/json"

	response, err := http.Post(myurl, contentType, requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(content))

}

func PerformFormMethod() {
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Sachin")
	data.Add("latsName", "Arora")
	data.Add("email", "sachin@gmail.com")

	response, err := http.PostForm(myurl, data)
	defer response.Body.Close()

	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(content))

}
