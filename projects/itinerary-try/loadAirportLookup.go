package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
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

	//create a map to store airport codes names
	lookupMap := make(map[string]string)
	count := 0

	// iterate through records and populate the map
	for indx, eachRecode := range records {
		//skip the header row
		if indx == 0 {
			continue // skip header
		}
		if len(eachRecode) < 5 {
			continue // skip if not enough columns
		}

		// get iata code and airport name
		iata := strings.TrimSpace(eachRecode[4])
		name := strings.TrimSpace(eachRecode[0])

		if iata != "" && name != "" {
			lookupMap[iata] = name
			count++
		}
	}

	fmt.Printf("Succefully laoded %d airport codes from lookup file \n", count)
	return lookupMap, nil
}
