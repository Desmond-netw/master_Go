package main

import (
	"html/template"
	"net/http"
)

// handler for templates
var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

// home page handler
func homeHandler(wr http.ResponseWriter, req *http.Request) {
	templates.ExecuteTemplate(wr, "layout.html", nil)
}

func listhandler(endpoint string, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		items, err := fetchItems(endpoint)
		if err != nil {
			http.Error(w, "failed to fetch data", http.StatusInternalServerError)
			return
		}

		data := struct {
			Title string
			Items []Item
		}{
			Title: title,
			Items: items,
		}

		err = templates.ExecuteTemplate(w, "layout.html", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
