package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reader is a global variable of type *bufio.Reader
var reader = bufio.NewReader(os.Stdin)

func main() {
	welcome := "Welcome our Pizza app!"
	fmt.Println(welcome)
	inputPromptMsg := "Please rate our Pizza taste from 1 - 5"
	fmt.Println(inputPromptMsg)

	// comma ok || error ok
	inputRating, inputErr := reader.ReadString('\n')
	if inputErr != nil {
		fmt.Println("Error in input", inputErr)
	}

	parsedRating, parseErr := strconv.ParseInt(strings.TrimSpace(inputRating), 10, 32)
	if parseErr != nil {
		fmt.Println("Error in parse", parseErr)
	}

	if parsedRating < 1 || parsedRating > 5 {
		fmt.Println("Invalid rating", parsedRating)
		return
	}

	fmt.Println("Thank you for your rating", parsedRating)

}
