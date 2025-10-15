package main

// imports
import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Application main entry
func main() {
	// set h flag helper to display application usage
	help := flag.Bool("h", false, "Display Usage")
	flag.Parse()

	//if -h is passe, print usage and exit
	if *help {
		fmt.Println("itinerary usage:")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	// check if exactly 3 arguments are passed (input, output,, lookup)
	args := flag.Args()
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Error : Expected 3 arguments")
		fmt.Println("Usage: go run . ./input.txt ./output.txt ./airport-lookup.txt")
		os.Exit(1)
	}
	// Assign variables to arguments
	inputPath := args[0]
	outputPath := args[1]
	lookupPath := args[2]

	//Load airport lookup data from CSV into a map
	airportMap, err := loadAirportLookup(lookupPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error laoding airport lookup : %v\n", err)
		os.Exit(1)
	}

	// Read the raw itinerary text from the input file
	inputData, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatal("Error Reading lookup file")
		os.Exit(1)
	}

	// split the input into lines for processing
	lines := strings.Split(string(inputData), "\n")
	var outputLines []string

	// process each line : replace airport coes with full names
	for _, eachline := range lines {
		for code, name := range airportMap {
			if strings.Contains(eachline, code) {
				eachline = strings.ReplaceAll(eachline, code, name)
			}
		}
		outputLines = append(outputLines, eachline)
	}

	// Write the processed lines to the output file
	err = os.WriteFile(outputPath, []byte(strings.Join(outputLines, "\n")))

}
