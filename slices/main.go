package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices or Dynamic array in Go language")

	// Standard way to create a slice or a dynamic array
	var fruits = []string{"mango", "apple", "orange"}
	fmt.Println("Fruits list in slices or dynamic array", fruits)
	fruits = append(fruits, "banana", "date")
	fmt.Println("Fruits list after adding more elements", fruits)
	fruits = append(fruits[:3])
	fmt.Println("Fruits list after removing last two elements", fruits)
	sort.Strings(fruits)
	fmt.Println("Fruits list after sorting elements", fruits)

	// Another way to create a slice or a dynamic array
	var scores = make([]int, 3)
	scores[0] = 123
	scores[1] = 456
	scores[2] = 789
	fmt.Println("Scores list in slices or dynamic array", scores)
	scores = append(scores, 987, 654, 321)
	fmt.Println("Scores list after adding more elements", scores)
	sort.Ints(scores)
	fmt.Println("Scores list after sorting elements", scores)
}
