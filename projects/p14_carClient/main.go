package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/models", modelsHandler)
	http.HandleFunc("/manufacturers", manufacturerHandler)
	http.HandleFunc("/categories", categoryHandler)

	fmt.Println("Starting Server.....")
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
