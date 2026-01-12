package main

import (
	"html/template"
	"net/http"
)

// Page Data

type PageData struct {
	Title, User string
}

func main() {
	http.HandleFunc("/", home)
}

// Global scope

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html")
}

// renderTemplate
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Template ParseFiles html file to be rendered
	tpl, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, nil)

}
