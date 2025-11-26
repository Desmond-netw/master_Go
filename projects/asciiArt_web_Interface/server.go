package main

import (
	"fmt"
	"html/template"
	"interface/resource"
	"log"
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

	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}

// server handler function
func actionHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// get request input value from client side and store their values
	InputText := req.FormValue("inputText")
	Mode := req.FormValue("action")

	data := PageData{}

	// validate InputText
	if InputText == "" {
		w.WriteHeader(http.StatusBadRequest)
		data.Error = "Input cannot be empty"
		tmpl.ExecuteTemplate(w, "index.gohtml", data)
	}

	// store result and err
	var (
		result string
		err    error
	)

	// process the Mode and data
	switch Mode {
	case "decode":
		result, err = resource.Decoder(InputText)
	case "encode":
		result, err = resource.Encoder(InputText)
	default:
		data.Error = "Invalid action selected"
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	// handle error when processing Mode
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

//---------------------------------
//      SEVER INITIALIZER
//---------------------------------

func RunServer() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/process", actionHandler)

	// server
	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
