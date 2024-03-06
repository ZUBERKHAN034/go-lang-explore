package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("Pointer value: ", ptr) // pointer default value is <nil>

	myAge := 21
	var myAgePtr = &myAge
	fmt.Println("myAgePtr holds memory address as value of another variable: ", myAgePtr)
	fmt.Println("we can access actual value of myAgePtr by using * before pointer variable: ", *myAgePtr)

	*myAgePtr = *myAgePtr + 1 // increment pointer variable BY 1
	fmt.Println("new value of myAge incremented by using myAgePtr: ", myAge)
}
