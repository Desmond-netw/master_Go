package main

import (
	"log"
	"os"
)

// Loaing CSV File and return a map of csv code
// the csv is expected to have at least 2 columns
func loadAirportLookup(filePath string) (map[string]string, error) {
	// open csv file
	myfile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err // return an error if file can't be open
	}
	defer myfile.Close() // Ensure the file is closed function exit

	// create a new CSV reader
}
