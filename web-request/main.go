package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Requesting into web in GO lang...")

	const baseUrl string = "https://lco.dev/"

	response, responseErr := http.Get(baseUrl)
	if responseErr != nil {
		panic(responseErr)
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		panic("Error: Status code not 200")
	}

	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		panic(responseBodyErr)
	}

	responseBoyContent := string(responseBody)
	fmt.Println("Your request url requested website body is:\n", responseBoyContent)
}
