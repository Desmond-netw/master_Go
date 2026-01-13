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

// home page handler
func homeHandler(wr http.ResponseWriter, req *http.Request) {
	renderTemplate(wr, "layout.html", nil)
}

// Model handler to send model data to html templ
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

	renderTemplate(wr, "layout.html", carsData)
}

// Manufacturer handler to send manufacture data to html
func manufacturerHandler(wr http.ResponseWriter, req *http.Request) {
	// geting manuf data
	manufacturer, err := fetchManufacturers()
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}

	manufacturerData := struct {
		Title         string
		Manufacturers []Manufacturer
	}{
		Title:         "Manufacturers",
		Manufacturers: manufacturer,
	}

	// render template to html output
	renderTemplate(wr, "layout.html", manufacturerData)

}

// func to handle template rendering
func renderTemplate(wr http.ResponseWriter, tmpl string, data any) {
	//passing specific html templates
	tpl, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
	}

	tpl.Execute(wr, data)
}
