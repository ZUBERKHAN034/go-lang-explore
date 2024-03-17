package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("POST request in GO lang...")

	endpoint := "http://localhost:3000/post"
	payloadString := `{"name": "Zuber Khan", "age": 22}`

	PerformPostRequest(endpoint, payloadString)
}

func PerformPostRequest(endpoint string, jsonString string) {

	// sending demo json payload send to a POST request
	// this is a simple way to send a POST request
	parsedJsonPayload := strings.NewReader(jsonString)

	response, responseErr := http.Post(endpoint, "application/json", parsedJsonPayload)
	if responseErr != nil {
		panic(responseErr)
	}

	defer response.Body.Close()

	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		panic(responseBodyErr)
	}

	// using strings.Builder to build the response body
	// this is recommended for large data
	var responseBuilder strings.Builder
	responseBuilder.Write(responseBody)

	fmt.Println("Response is:", responseBuilder.String())
}
