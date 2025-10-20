package main

import (
	"fmt"
	"os"
)

// Processing itinerary data
func processItinerary(inputPath, outputPath, lookupPath string) error {
	// check if input file exists
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input not found")
	}
	// check if lookup file exists
	if _, err := os.Stat(lookupPath); os.IsNotExist(err) {
		return fmt.Errorf("lookup file not found")
	}

	// Load airport lookup data from CSV
	airportData, err := loadAirportLookup(lookupPath)
	if err != nil {
		return fmt.Errorf("error loading airport lookup: %v", err)
	}

	// Read the raw itinerary text from the input file
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Process the content line by line
	processedContent := processContent(string(inputData), airportData)

	// write the processed content to the output file
	err = os.WriteFile(outputPath, []byte(processedContent), 0644)
	if err != nil {
		return fmt.Errorf("error writing to output file: %v", err)
	}

	return nil
}

// processContent replaces airport codes in the input text with full airport names
func processContent(content string, airportData *AirportData) string {
	// first, trim any vertical whitespace
	content = trimVerticalWhitespace(content)

	// Process airport codes
	content = processAirportCodes(content, airportData)

	// process dates and times
	content = processDatesAndTimes(content)

	return content
}
