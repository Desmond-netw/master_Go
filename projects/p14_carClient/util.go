package main

import "html/template"

// Building lookup maps

func manufacturerMap(manufacturer []Manufacturer) map[int]string {
	manuMap := make(map[int]string)
	for _, manuf := range manufacturer {
		manuMap[manuf.ID] = manuf.Name
	}
	return manuMap
}

// category map
func categoryMap(category []Category) map[int]string {
	catMap := make(map[int]string)
	for _, cats := range category {
		catMap[cats.ID] = cats.Name
	}
	return catMap
}

// -------------------
// Global Variables

var (
	homeTmpl  = template.Must(template.ParseFiles("templates/layout.html", "templates/views.html"))
	modelTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/models.html"))
	categTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/category.html"))
	manufTmpl = template.Must(template.ParseFiles("templates/layout.html", "templates/manufacturer.html"))
)
