package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Switch case in Golang")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the dice number from 1 to 6")

	input, readErr := reader.ReadString('\n')
	if readErr != nil {
		fmt.Printf("Error while reading the input %v", readErr)
	}

	diceNumber, parseIntErr := strconv.ParseInt(strings.TrimSpace(input), 0, 32)
	if parseIntErr != nil {
		fmt.Printf("Error while converting the input to integer %v", parseIntErr)
	}

	// In golang there is no break to stop the case it stopped automatically
	// If you want to get the falling flow you can use [fallthrough]
	switch diceNumber {
	case 1:
		fmt.Println("Dice number is 1 you can open")
	case 2:
		fmt.Println("You can move 2 steps")
	case 3:
		fmt.Println("You can move 3 steps")
	case 4:
		fmt.Println("You can move 4 steps")
	case 5:
		fmt.Println("You can move 5 steps")
	case 6:
		fmt.Println("You can move 6 steps and 1 more dice roll")
		fallthrough
	case 7:
		fmt.Println("Woo hoo! You have 1 more dice roll!!!")
	}
}
