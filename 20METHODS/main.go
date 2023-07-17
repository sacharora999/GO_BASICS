package main

import "fmt"

func main() {
	fmt.Println("Methods in GOLANG")

	Sachin := User{"Arora", 10, true, "Lead"}

	fmt.Println(Sachin)
	fmt.Println(Sachin.Designation)
	fmt.Printf("Details are %+v\n", Sachin)
	Sachin.getStarted()

	Sachin.setAge(20)
	Sachin.getAge()
	fmt.Printf("Details are %+v\n", Sachin)

}

type User struct {
	Name        string
	Age         int64
	Status      bool
	Designation string
}

func (u User) getStarted() {
	fmt.Println("Is user active", u.Status)
}

func (u User) getAge() {
	fmt.Println("User Age is ", u.Age)
}

func (u User) setAge(myAge int64) {
	u.Age = myAge
}
