package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// main entry for runing server
	http.HandleFunc("/", index)
	http.HandleFunc("/process", processor)

	// listen and server on address
	http.ListenAndServe(":5500", nil)

}

// Render about page thus the corresponding html code to execute
func index(wr http.ResponseWriter, req *http.Request) {
	// execute index.html
	err := tmpl.ExecuteTemplate(wr, "index.html", nil)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
}

// processor
func processor(wr http.ResponseWriter, req *http.Request) {
	// check for valid postmethod and process form data
	if req.Method != http.MethodPost {
		http.Redirect(wr, req, "/", http.StatusSeeOther)
		return
	}

	// grap the data var from input
	fname := req.FormValue("firster")
	lname := req.FormValue("laster")

	data := struct {
		FirstName string
		LastName  string
	}{
		FirstName: fname,
		LastName:  lname,
	}

	err := tmpl.ExecuteTemplate(wr, "process.html", data)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
		return
	}
}
