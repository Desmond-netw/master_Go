package main

import (
	"html/template"
	"net/http"
)

// index func to read index.html file
func index(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, "index.html")
}
func main() {
	// web server main entry for handling request and post
	http.HandleFunc("/", index) // pass index fun as hompage
	http.ListenAndServe(":3000", nil)
}
