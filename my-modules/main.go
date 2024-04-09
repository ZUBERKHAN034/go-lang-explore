package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Module in GO lang...")
	severing()

	router := mux.NewRouter()
	router.HandleFunc("/", serveHome).Methods("GET")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func severing() {
	fmt.Println("Server is running on port 3000...")
}

func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Hello from go module!<h1>"))
}
