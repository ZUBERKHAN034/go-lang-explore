package main

import "fmt"

// defer function in Go
// defer is used to ensure that a function call is performed
// later in a program's execution, usually for purposes of cleanup.
func main() {
	fmt.Println("Defer function in Go")

	// defer function call
	defer fmt.Println("AI: Software engineering\n" +
		"is the application\n" +
		"of engineering to the design,\n" +
		"development, implementation,\n" +
		"testing and maintenance\n" +
		"of software in a systematic method.") // last to be executed in defer
	defer fmt.Println("AI: Thinking...")   // 3rd to be executed in defer
	defer fmt.Println("AI: Processing...") // 2nd to be executed in defer
	defer fmt.Println("AI: Searching...")  // 1st to be executed in defer

	fmt.Println("Prompt: What is software engineering?")
	waitDefer()
}

func waitDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Wait for seconds", i)
	}
}
