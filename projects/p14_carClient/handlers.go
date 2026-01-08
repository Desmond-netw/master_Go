package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

// handler for templates

var templates = template.Must(
	template.ParseGlob("templates/*.html"),
)

func fetchItems(endpoint string) ([]Item, error) {
	resp, err := http.Get(APIURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("HTTP status : %s", resp.Status)
	}

	// var items to item struct
	var items []Item
	err = json.NewDecoder(resp.Body).Decode(&items)
	return items, err
}

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

		templates.ExecuteTemplate(w, "layout.html", data)
	}
}
