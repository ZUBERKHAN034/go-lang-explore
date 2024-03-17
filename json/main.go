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
	coursesJsonString := `
		{
			"name": "Angular",
			"price": 399,
			"platform": "LearnCodeOnline",
			"tags": ["full-stack", "js"]
		}`

	coursesJson := []byte(coursesJsonString)

	isJson := json.Valid(coursesJson)
	fmt.Println("Is JSON valid:", isJson)
	if !isJson {
		panic("Invalid JSON")
	}

	// decoding the JSON string to a slice of courses
	var myCourse course

	decodeErr := json.Unmarshal(coursesJson, &myCourse)
	if decodeErr != nil {
		panic(decodeErr)
	}

	fmt.Println("Decoded courses are:", myCourse)

	// handling in the key value pair
	var coursesMap map[string]interface{}
	json.Unmarshal(coursesJson, &coursesMap)

	fmt.Printf("%#v\n", coursesMap)

	for key, value := range coursesMap {
		fmt.Printf("Key: %v -> Value: %v and Type is %T\n", key, value, value)
	}

}
