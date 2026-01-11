package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// templates
var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

// Handling list of API
// Making sure
// JSON API data ->Go struct
// Go struct - HTML Templates

// MODEL HANDLER
func modelsHandler(wr http.ResponseWriter, req *http.Request) {
	models, err := fetchModel()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Models fetchd", len(models))
	carsData := struct {
		Title  string
		Models []CarModel
	}{
		Title:  "Car Model",
		Models: models,
	}

	templates.ExecuteTemplate(wr, "layout.html", carsData)
	// if err != nil {
	// 	http.Error(wr, err.Error(), http.StatusInternalServerError)
	// }
}

// home page handler
func homeHandler(wr http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(wr, "layout.html", nil)
}
