package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// Application to read csv file

func main() {
	// using os.open to open a file
	file, err := os.Open("student.csv")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	// csv reader
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error readin csv")
	}
	for _, eachrecords := range records {

		fmt.Println(eachrecords)
	}

}
