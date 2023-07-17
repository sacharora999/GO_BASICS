package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Price    int      `json:"price"`
	Password string   `json:"-"`
	Tags     []string `json:"tags omitempty"`
}

func main() {
	fmt.Println("Welcome to Create JSON in Golang")
	DecodeJson()

}

func EncodeJson() {

	lcocourses := []courses{
		{"ReactJS", "sac@gmail.com", 900, "abcd", []string{"mytag1", "mytag2", "mytag3", "mytag4"}},
		{"NodeJS", "sac@gmail.com", 970, "wsss", []string{"mytag8", "mytag2", "mytag9", "mytag7"}},
	}

	finalJson, err := json.MarshalIndent(lcocourses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {

	jsonDataFromWeb := []byte(`

	{
		"name": "ReactJS",
		"email": "sac@gmail.com",
		"price": 900,
		"tags omitempty": [
				"mytag1",
				"mytag2",
				"mytag3",
				"mytag4"
		]
	}`)

	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Print("Valid Json Data")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Printf("JSON WAS NOT VALID")
	}

	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key is %v Value is %v and Type is %T\n", k, v, v)
	}
}
