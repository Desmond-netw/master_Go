package main

import (
	"html/template"
	"interface/resource"
	"net/http"
)

// Template glob
var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// Page data structure
type PageData struct {
	Result string
	Status string
	Error  string
}

// Homepage route handler for home rendering
func homeHandler(w http.ResponseWriter, req *http.Request) {

	tmpl.ExecuteTemplate(w, "index.html", nil)
}

// server handler function
func handler(w http.ResponseWriter, req *http.Request) pageData {

	if req.Method != http.MethodPost {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return pageData{
			Error: "",
		}
	}

	// Render the home page
	http.HandleFunc("/", homepage)

	// get request input from client side and store their values
	inputText := req.FormValue("inputText")
	modeAction := req.FormValue("action")

	// set variables
	myresult := string
	data := pageData{}

	// validate inputText
	if inputText == "" {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = "Please select either Decoder or Encoder"
		tmpl.Execute(w, data)
	}
	//pass data to result
	if modeAction == "decode" {
		myresult, err = resource.Decoder(inputText)
	} else if modeAction == "encode" {
		myresult, err = resource.Encoder(inputText)
	}

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = err.Error()
	}

	return pageData{}
}

// rendering template func
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parsing the specified template file bing passed as input
	t := template.Must(template.ParseFiles("templates/" + tmpl)) // t = template
	// execute
	t.Execute(w, nil)
}
