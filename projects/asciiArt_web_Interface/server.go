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
func actionHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// get request input value from client side and store their values
	inputText := req.FormValue("inputText")
	mode := req.FormValue("action")

	data := PageData{}

	// validate inputText
	if inputText == "" {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = "Input cannot be empty"
		tmpl.ExecuteTemplate(w, "index.html", data)
	}

	// store result and err
	var (
		result string
		err    error
	)

	// process the mode and data
	switch mode {
	case "decode":
		result, err = resource.Decoder(inputText)
	case "encode":
		result, err = resource.Encoder(inputText)
	default:
		data.Error = "Invalid action selected"
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	// handle error when processing mode
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = err.Error()
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	data.Result = result
	data.Status = "success"

	tmpl.ExecuteTemplate(w, "result.tml", data)
}

// rendering template func
func renderTemplate(w http.ResponseWriter, tmpl string) {
	// Parsing the specified template file bing passed as input
	t := template.Must(template.ParseFiles("templates/" + tmpl)) // t = template
	// execute
	t.Execute(w, nil)
}
