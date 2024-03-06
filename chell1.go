package main

import (
	"fmt"
)

func main() {
	adams := 42
	fmt.Printf("42 as binary, %b \n", adams)
	fmt.Printf("42 as hexadecimal, %x \n", adams)

	// print these values as both binary & hexadecimal
	a, b, c, d, e, f := 10, 100, 1000, 10000, 100000, 1000000

	fmt.Printf("10, 100, 1000, 10000, 100000, 1000000 as binary, %b %b %b %b %b %b \n", a, b, c, d, e, f)
	fmt.Printf("10, 100, 1000, 10000, 100000, 1000000 as hexadecimal, %x %x %x %x %x %x \n", a, b, c, d, e, f)

}