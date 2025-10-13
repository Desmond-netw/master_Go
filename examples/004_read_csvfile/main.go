package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// functio to read csv file

func main() {
	// open
	file, err := os.Open("airport-lookup.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	// read the file
	reader := csv.NewReader(file)

	// records to read all
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading this file")
	}

	// create a new csv file and write content of csv to that file
	file, err = os.Create("NewCSV.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("\nCSV Records")
	for _, singleRecord := range records {
		fmt.Println(singleRecord)
	}
}
