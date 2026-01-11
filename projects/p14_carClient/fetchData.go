package main

import (
	"encoding/json"
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

// fetch car Models from json data and decode to struct
func fetchModel() ([]CarModel, error) {
	resp, err := http.Get(APIURL + "/models")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var models []CarModel
	err = json.NewDecoder(resp.Body).Decode(&models)
	return models, err
}
