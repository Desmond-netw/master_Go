package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

// Application to read csv file

func main() {

	// display help flag
	help := flag.Bool("h", false, "Display help and Usage")
	flag.Parse()
	if *help {
		fmt.Println("App usage: ")
		fmt.Println("Application usag go run main.go ")
		return
	}
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
