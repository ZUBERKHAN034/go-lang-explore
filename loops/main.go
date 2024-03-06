package main

import "fmt"

func main() {
	fmt.Println("Loops in Golang")

	languages := []string{"hindi", "english", "urdu", "arabic"}

	// Standard for loop
	fmt.Println("Standard for loop")
	for i := 0; i < len(languages); i++ {
		fmt.Printf("index: %d, value: %v\n", i, languages[i])
	}

	// Range for loop
	fmt.Println("Range for loop")
	for r := range languages {
		fmt.Printf("index: %d, value: %v\n", r, languages[r])
	}

	// Foreach like for loop in golang
	fmt.Println("Foreach like for loop in golang")
	for index, value := range languages {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}
}
