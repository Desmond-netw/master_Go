package main

import (
	"encoding/csv"
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

	//create a new csv reader
	csvReader := csv.NewReader(csvFile)
	//read all records from csv
	records, es := csvReader.ReadAll()
	if es != nil {
		return nil, fmt.Errorf("error reading csv: %v", es)
	}
}
