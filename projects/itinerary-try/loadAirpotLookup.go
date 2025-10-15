package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
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
	csvReader := csv.NewReader(myfile)

	// Read all records from csv file
	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
		return nil, err // Return error and log if reading fails
	}

	// initialize a map to store airport names
	lookupMap := make(map[string]string)

	for _, eachRecord := range csvRecords {
		if len(eachRecord) >= 2 {
			code := strings.TrimSpace(eachRecord[0]) // eg. Airpot Code "JFK"
			name := strings.TrimSpace(eachRecord[1]) // eg. full name John F. Kennedy"
			lookupMap[code] = name
		}
	}

	return lookupMap, nil // return the populated map

}
