package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:8000/learn?coursename=reactjs&paymentid=12345"

func main() {
	fmt.Println("Welcome to Handing URLs in GOLANG")
	fmt.Println(myurl)

	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme : ", result.Scheme)
	fmt.Println("Path : ", result.Path)
	fmt.Println("Host : ", result.Host)
	fmt.Println("Port : ", result.Port())
	fmt.Println("RawQuery : ", result.RawQuery)

	qparams := result.Query()

	fmt.Println("Query : ", qparams["coursename"])

	for _, value := range qparams {
		fmt.Println(value)
	}

	constructURL := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=sachin",
	}

	finalURL := constructURL.String()

	fmt.Print(finalURL)

}
