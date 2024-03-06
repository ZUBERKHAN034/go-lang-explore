package main

import "fmt"

// in global scope use can only use long declaration syntax
var myname string = "Zuber Khan global scope"

// short declaration syntax not allowed like this
// myage := 22

func main() {
	fmt.Println("Variables in golang")
	// var keyword is used to declare variables
	var name string = "Zuber Khan"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	var isLogged = true
	fmt.Println(isLogged)
	fmt.Printf("Variable is of type: %T \n", isLogged)

	var age = 22
	fmt.Println(age)
	fmt.Printf("Variable is of type: %T \n", age)

	var height = 5.7
	fmt.Println(height)
	fmt.Printf("Variable is of type: %T \n", height)

	// default values and some aliases
	var startAge int
	fmt.Println(startAge)
	fmt.Printf("Variable is of type: %T \n", startAge)

	var startName string
	fmt.Println(startName)
	fmt.Printf("Variable is of type: %T \n", startName)

	// short variable declaration
	x := 10
	y := "hello"
	z := true
	fmt.Println(x, y, z)
	fmt.Printf("Variable is of type: %T %T %T \n", x, y, z)
}
