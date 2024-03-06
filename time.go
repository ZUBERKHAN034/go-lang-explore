package main;

import "fmt";
import "time";

func main() {
var dateAndTime time.Time = time.Now();
fmt.Printf("Current Date & Time is %s", dateAndTime);
}