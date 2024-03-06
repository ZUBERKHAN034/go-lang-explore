package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func main() {
	fmt.Println("Structs in Golang")

	myDetails := User{"Zuber", "zuberkhan034@gmail.com", 21, true}

	fmt.Println("My details are :", myDetails)
	// %+v gives more information about the struct
	fmt.Printf("My details are : %+v\n", myDetails)
}
