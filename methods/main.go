package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) getStatus() {
	if u.Status {
		fmt.Println("user online")
	} else {
		fmt.Println("user offline")
	}
}

func (u User) setTempEmail(email string) {
	u.Email = email
	fmt.Println("new temporary email is", u.Email)
}

func main() {
	fmt.Println("Methods in Golang")

	myDetails := User{"Zuber", "zuberkhan034@gmail.com", 21, true}

	fmt.Println("My details are :", myDetails)
	// %+v gives more information about the struct
	fmt.Printf("My details are : %+v\n", myDetails)

	// method calling
	myDetails.getStatus()

	// note its not changed original value because
	// that setTempEmail function contains copy of original
	// not the original by using the pointer we manipulate the original vale
	myDetails.setTempEmail("temp@gmail.com")
}
