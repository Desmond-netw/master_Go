package main

import (
	"fmt"
	"os"
)

// loading CSV File and return a map of csv code
// the csv is expectd to have at least 2 columns
func loadAirportLookup(filePath string) (map[string]string, error) {
	// open csv file

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %v", err)
	}
	defer csvFile.Close()

}
