package main

import "fmt"

func main() {
	// in golang we can only call the functions but not define them inside the
	fmt.Println("Functions in Golang")
	greet()

	resultAdd := add(5, 5)
	fmt.Println("Result: ", resultAdd)

	resultAddValues := addValues(1, 2, 3, 4)
	fmt.Println("Result: ", resultAddValues)
}

func greet() {
	fmt.Println("Hello from Golang")
}

func add(val1, val2 int) int {
	return val1 + val2
}

func addValues(values ...int) int {

	sumOfValues := 0

	for _, value := range values {
		sumOfValues += value
	}

	return sumOfValues
}
