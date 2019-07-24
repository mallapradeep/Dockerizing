package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("GO Docker")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello WORLD")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

//RUN "docker build -t my-go-app" in terminal
// see images --> docker images
// to run docker --> docker run -d -p 8080:8081 my-go-app
//to see how many docker images are running --> docker ps
//to kill docker image --> docker kill container-id
