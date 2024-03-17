package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("GET request in GO lang...")

	endpoint := "http://localhost:3000/get"
	PerformGetRequest(endpoint)
}

func PerformGetRequest(endpoint string) {

	response, responseErr := http.Get(endpoint)
	if responseErr != nil {
		panic(responseErr)
	}

	defer response.Body.Close()

	fmt.Println("Status code is:", response.StatusCode)
	fmt.Println("Content length is:", response.ContentLength)

	// Reading the response body
	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		panic(responseBodyErr)
	}

	// easy way to convert byte array to string
	// but not recommended for large data
	responseBodyContent := string(responseBody)
	fmt.Println("Response body is:", responseBodyContent)

	// using strings.Builder to build the response body
	// this is recommended for large data
	var responseBuilder strings.Builder
	byteCount, byteCountErr := responseBuilder.Write(responseBody)
	if byteCountErr != nil {
		panic(byteCountErr)

	}

	fmt.Println("Byte count is:", byteCount)
	fmt.Println("Response body is:", responseBuilder.String())

}
