package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// call http API
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Coding in golang")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here is a Hi page from GO http")
	})
	//set port
	port := ":5000"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
