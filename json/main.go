package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"` // this field will be ignored
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("JSON handling in GO lang...")

	EncodeJson()
}

func EncodeJson() {

	courses := []course{
		{"ReactJS", 299, "LearnCodeOnline", "abc@123", []string{"web-dev", "js"}},
		{"MERN", 199, "LearnCodeOnline", "abc@123", []string{"full-stack", "js"}},
		{"Angular", 399, "LearnCodeOnline", "abc@123", nil},
	}

	// encoding the courses slice to JSON
	// and printing the JSON
	coursesJson, coursesJsonErr := json.MarshalIndent(courses, "", "\t")
	if coursesJsonErr != nil {
		panic(coursesJsonErr)
	}

	fmt.Println("Encoded JSON is:", string(coursesJson))

}
