package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("URL handling in GO lang...")

	const baseUrl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=123"

	// URL parsing
	parsedUrl, parsedUrlErr := url.Parse(baseUrl)
	if parsedUrlErr != nil {
		panic(parsedUrlErr)
	}

	fmt.Println("Scheme is:", parsedUrl.Scheme)
	fmt.Println("Host is:", parsedUrl.Host)
	fmt.Println("Path is:", parsedUrl.Path)
	fmt.Println("Query is:", parsedUrl.RawQuery)
	fmt.Println("Port is:", parsedUrl.Port()) // port is a method not a property

	qParams := parsedUrl.Query()
	fmt.Println("Query Params:", qParams)

	for key, value := range qParams {
		fmt.Printf("Key: %s -> Value: %s\n", key, value)
	}

	// URL building
	newUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tut",
		RawQuery: "coursename=reactjs&paymentid=123",
	}

	fmt.Println("Built URL is:", newUrl.String())

}
