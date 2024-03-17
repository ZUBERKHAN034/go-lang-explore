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
	DecodeJson()
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

func DecodeJson() {

	// a simple JSON string
	coursesJsonString := `[
		{
			"name": "ReactJS",
			"price": 299,
			"platform": "LearnCodeOnline",
			"tags": ["web-dev", "js"]
		},
		{
			"name": "MERN",
			"price": 199,
			"platform": "LearnCodeOnline",
			"tags": ["full-stack", "js"]
		},
		{
			"name": "Angular",
			"price": 399,
			"platform": "LearnCodeOnline"
		}
	]`

	coursesJson := []byte(coursesJsonString)

	isJson := json.Valid(coursesJson)
	if !isJson {
		panic("Invalid JSON")
	}

	// decoding the JSON string to a slice of courses
	var courses []course

	decodeErr := json.Unmarshal(coursesJson, &courses)
	if decodeErr != nil {
		panic(decodeErr)
	}

	fmt.Println("Decoded courses are:", courses)
}
