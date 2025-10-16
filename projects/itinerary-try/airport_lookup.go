package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

type AirportInfo struct {
	Name         string
	ISOCountry   string
	Municipality string
	ICAOCode     string
	IATACode     string
	Coordinates  string
}

type AirportData struct {
	ByIATA map[string]*AirportInfo
	ByICAO map[string]*AirportInfo
}

func loadAirportLookup(filePath string) (*AirportData, error) {
	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Airport lookup not found")
	}

	// Open CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Airport lookup not found")
	}
	defer file.Close()

	// Create CSV reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Airport lookup malformed")
	}

	// Check if we have at least a header row
	if len(records) < 1 {
		return nil, fmt.Errorf("Airport lookup malformed")
	}

	airportData := &AirportData{
		ByIATA: make(map[string]*AirportInfo),
		ByICAO: make(map[string]*AirportInfo),
	}

	// Process records (skip header row)
	for i := 1; i < len(records); i++ {
		record := records[i]

		// Check if we have enough columns and no empty cells
		if len(record) < 6 {
			return nil, fmt.Errorf("Airport lookup malformed")
		}

		for _, cell := range record {
			if strings.TrimSpace(cell) == "" {
				return nil, fmt.Errorf("Airport lookup malformed")
			}
		}

		info := &AirportInfo{
			Name:         strings.TrimSpace(record[0]),
			ISOCountry:   strings.TrimSpace(record[1]),
			Municipality: strings.TrimSpace(record[2]),
			ICAOCode:     strings.TrimSpace(record[3]),
			IATACode:     strings.TrimSpace(record[4]),
			Coordinates:  strings.TrimSpace(record[5]),
		}

		// Add to maps if codes are present
		if info.IATACode != "" {
			airportData.ByIATA[info.IATACode] = info
		}
		if info.ICAOCode != "" {
			airportData.ByICAO[info.ICAOCode] = info
		}
	}

	return airportData, nil
}
