package main

import (
	"net/http"
)

/*
	BASE API URL : http://localhost:3000/api

	EndPoints:
	GET /models
	GET /manufacturers
	GET /categories
*/

const APIURL = "http://localhost:3000/api"

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image,omitempty"`
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/models", listhandler("/models", "Car Models"))
	http.HandleFunc("/manufacturers", listhandler("/manufacturers", "manufacturers"))
	http.HandleFunc("/categories", listhandler("/categories", "categories"))

	http.ListenAndServe(":8080", nil)
}
