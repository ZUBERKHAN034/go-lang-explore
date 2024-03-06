package main

import (
	"bufio" // Package bufio implements buffered I/O.
	"fmt"   // Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	"os"    // Package os provides a platform-independent interface to operating system functionality.
)

func main() {
	welcome := "Welcome to user input!" // Define a welcome message
	fmt.Println(welcome)                // Print the welcome message

	reader := bufio.NewReader(os.Stdin) // Create a new reader, using the standard input as its source.

	inputPromptMsg := "Please enter your username :" // Define a prompt message for user input
	fmt.Println(inputPromptMsg)                      // Print the prompt message

	input, _ := reader.ReadString('\n')      // Read user input until the first occurrence of '\n' (newline)
	fmt.Println("Your username is: ", input) // Print the user's input
}
