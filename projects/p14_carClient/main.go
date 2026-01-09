package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/models", listhandler("/models", "Car Models"))
	http.HandleFunc("/manufacturers", listhandler("/manufacturers", "manufacturers"))
	http.HandleFunc("/categories", listhandler("/categories", "categories"))

	http.ListenAndServe(":8080", nil)
}
