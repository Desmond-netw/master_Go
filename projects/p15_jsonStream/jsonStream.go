package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// json stream

func main() {
	const JsonCountry = `
	{"Country" : "Ghana", "City": "Accra"}
	{"Country": "Finland","City": "Helsinki"}
	{"Country": "USA","City": "Washton"}
	{"Country": "United Kingdom","City": "London"}
	{"Country": "Germany","City": "Belin"}
	{"Country": "Italy","City": "Rome"}
	`

	// type to get contry
	type Map struct {
		Country, City string
	}

	dec := json.NewDecoder(strings.NewReader(JsonCountry))

	// loop while reading each line into struct
	for {
		var m Map

		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("country: %s, City: %s\n", m.Country, m.City)
	}
}
