package main

import "fmt"

func main() {
	fmt.Println("Object or HashTable or HashMap in Golang")

	// Declaration only
	languages := make(map[string]string)
	// Initialization after declaration
	languages["JS"] = "javascript"
	languages["PHP"] = "php"
	languages["RB"] = "ruby"

	fmt.Println("Languages :", languages)

	delete(languages, "RB")

	fmt.Println("Languages after deleting on entry :", languages)

	// Declaration and initialization
	countries := map[string]string{
		"IN": "India",
		"PK": "Pakistan",
		"US": "United States",
	}

	fmt.Println("Countries :", countries)

	// loops over maps

	for key, val := range countries {
		fmt.Printf("Key is %v, Value is %v\n", key, val)
	}
}
