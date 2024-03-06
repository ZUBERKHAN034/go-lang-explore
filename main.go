// example for nested functions
package main;

import "fmt";

func multiplyByTwo(num int) int{
 return num * 2;
}

func add(num1 int, num2 int) int{
	addedValue := num1 + num2;
	return multiplyByTwo(addedValue);
}

func main() {
	result := add(50, 50);
	fmt.Println(result);
}