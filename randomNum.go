package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var rangeNumber int = 100;
	var number int = rand.Intn(rangeNumber);
	fmt.Println("Random number is", number);
}
