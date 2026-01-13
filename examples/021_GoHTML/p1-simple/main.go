package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Page Data

type PageData struct {
	Title, User string
}

func main() {
	http.HandleFunc("/", home)
	fmt.Println("Server running .. localhost:4000")
	http.ListenAndServe(":4000", nil)

}

// Global scope

// home handler
func home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", PageData{Title: "Admin", User: "Analyst"})
}

// renderTemplate
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	// Template ParseFiles html file to be rendered
	tpl, err := template.ParseFiles("views/*" + tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data)

}
