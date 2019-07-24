package main

import (
	"fmt"
	"log"
	"net/http"
)

var Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// var Articles []Article

// A homePage function that will handle all requests to our root URL,

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

//a handleRequests function that will match the URL path hit with a defined function and

func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// a main function which will kick off our API.

func main() {
	handleRequests()
}
