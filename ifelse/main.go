package main

import "fmt"

func main() {
	fmt.Println("if else in Golang")

	if position := 3; position == 1 {
		fmt.Println("Gold with position: ", position)
	} else if position == 2 {
		fmt.Println("Silver with position: ", position)
	} else if position == 3 {
		fmt.Println("Browne with position: ", position)
	} else {
		fmt.Println("Not valid position: ", position)
	}
}
