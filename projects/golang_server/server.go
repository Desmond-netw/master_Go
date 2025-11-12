package main

import (
	"fmt"
	"net/http"
)

func main() {
	// call http API
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Coding in golang")
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf("Here is a Hi page from GO http")
	})
	//set port
}
