package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("POST Form request in GO lang...")

	endpoint := "http://localhost:3000/postform"

	// sending demo form data to a POST request
	demoFormData := url.Values{}
	demoFormData.Add("name", "Zuber Khan")
	demoFormData.Add("age", "22")
	demoFormData.Add("email", "zuberkhan034@gmail.com")

	PerformPostFormRequest(endpoint, demoFormData)
}

func PerformPostFormRequest(endpoint string, formData url.Values) {

	response, responseErr := http.PostForm(endpoint, formData)
	if responseErr != nil {
		panic(responseErr)
	}

	defer response.Body.Close()

	responseBody, responseBodyErr := io.ReadAll(response.Body)
	if responseBodyErr != nil {
		panic(responseBodyErr)
	}

	fmt.Println("Response is:", string(responseBody))

}
