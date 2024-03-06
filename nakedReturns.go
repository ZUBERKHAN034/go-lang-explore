// example for naked function return as variables
package main;

import "fmt";

func nakedReturn(val int) (first int , second int) {
 first = val / 2;
 second = val - first;
 return;
}

func main() {
 fmt.Print(nakedReturn(24));
}