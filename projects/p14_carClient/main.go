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

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image,omitempty"`
}

func fetchItems(endpoint string) ([]Item, error) {
	resp, err := http.Get(APIURL + endpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Errorf("HTTP status : %s", resp.Status)
	}

	// var items to items struct
	var items []Item
	err = json.NewDecoder(resp.Body).Decode(&items)
	return items, err
}
