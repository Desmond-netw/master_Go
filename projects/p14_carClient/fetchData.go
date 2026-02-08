package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
	BASE API URL : http://localhost:3000/api

	EndPoints:
	GET /models
	GET /manufacturers
	GET /categories
*/

const APIURL = "http://localhost:3000/api"

// Fetching API Data

/*
	Fetch data in Paraller & launch three goroutines,

each loading data from it's API endpoint
*/
func LoadData() ([]Category, []CarModel, []Manufacturer) {

	// result struct
	type result[T any] struct {
		data []T
		err  error
		src  string
	}

	catChanel := make(chan result[Category])
	carModelChanel := make(chan result[CarModel])
	manuChanel := make(chan result[Manufacturer])

	// load each data set
	go func() {
		data, err := fetchCategory()
		catChanel <- result[Category]{data, err, "Categories"}
	}()

	go func() {
		data, err := fetchModel()
		carModelChanel <- result[CarModel]{data, err, "Models"}
	}()

	go func() {
		data, err := fetchManufacturers()
		manuChanel <- result[Manufacturer]{data, err, "Manufacturers"}
	}()
	// storing var to data struct
	var (
		categories    []Category
		models        []CarModel
		manufacturers []Manufacturer
		errors        []APIError
	)

	for i := 0; i < 3; i++ {
		select {
		case r := <-catChanel:
			if r.err != nil {
				errors = append(errors, APIError{r.src, r.err.Error()})
			} else {
				categories = r.data
			}
		case r := <-carModelChanel:
			if r.err != nil {
				errors = append(errors, APIError{r.src, r.err.Error()})
			} else {
				models = r.data
			}

		case r := <-manuChanel:
			if r.err != nil {
				errors = append(errors, APIError{r.src, r.err.Error()})
			} else {
				manufacturers = r.data
			}

		}
	}

	return categories, models, manufacturers
}

// fetch car Models from json data and decode to Go struct
func fetchModel() ([]CarModel, error) {

	resp, err := http.Get(APIURL + "/models")
	if err != nil {
		return nil, fmt.Errorf("models API unreachable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Models API returned %d", resp.StatusCode)
	}
	var models []CarModel
	if err = json.NewDecoder(resp.Body).Decode(&models); err != nil {
		return nil, fmt.Errorf("Invalide models response")
	}

	return models, err
}

/* ---------- Fetch Manufacturers data from API ------*/

func fetchManufacturers() ([]Manufacturer, error) {
	resp, err := http.Get(APIURL + "/manufacturers")
	if err != nil {
		return nil, fmt.Errorf("Manufacturer API unreachable")
	}
	defer resp.Body.Close() // close response body

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("manufacturers API returned %d", resp.StatusCode)
	}

	var manufacturers []Manufacturer // var to store address manuf
	if err = json.NewDecoder(resp.Body).Decode(&manufacturers); err != nil {
		return nil, fmt.Errorf("invalide manufacturers response")
	} // decoding json data from resp.body to Go struct
	return manufacturers, err
}

/*----------- Fetching categories data*/
func fetchCategory() ([]Category, error) {
	resp, err := http.Get(APIURL + "/categories")
	if err != nil {
		return nil, fmt.Errorf("category API unreachable")
	}
	defer resp.Body.Close() // close response body

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("category API returned %d", resp.StatusCode)
	}
	var categories []Category
	if err = json.NewDecoder(resp.Body).Decode(&categories); err != nil {
		return nil, fmt.Errorf("invalide category response")
	}
	return categories, err
}

// func fetchItems(endpoint string) ([]Item, error) {
// 	baseURL := APIURL + endpoint // base url will be /api/models
// 	println("fetching url:", baseURL)
// 	// get response
// 	resp, err := http.Get(baseURL)
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// decode and read json data
// 	var items []Item
// 	err = json.NewDecoder(resp.Body).Decode(&items)

// 	println("Items fetched:", len(items)) // debug

// 	return items, err
// }
