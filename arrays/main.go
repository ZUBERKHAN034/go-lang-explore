package main

import "fmt"

func main() {
	fmt.Println("Arrays in Go language")

	var fruits [4]string
	fruits[0] = "mango"
	fruits[1] = "banana"
	fruits[3] = "apple"

	fmt.Println("Fruits list in an Arrays", fruits)
	fmt.Println("Fruits list length in an Arrays", len(fruits))

	var vegetables = [5]string{"ladyfinger", "tomato", "potato"}
	fmt.Println("Vegetables list in an Arrays", vegetables)
	fmt.Println("Vegetables list length in an Arrays", len(vegetables))
}
