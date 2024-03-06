// Example for multiple returns values
package main;

import "fmt";

func exchangeNumbers(num1, num2 int) (int, int) {
	return num2, num1;
};

func main() {
	num1, num2 := exchangeNumbers(1, 3);
	fmt.Println(num1, num2);
};
