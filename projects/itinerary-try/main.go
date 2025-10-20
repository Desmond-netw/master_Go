package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define commmand-line flags
	// dfine -h help flag to show usage information
	help := flag.Bool("h", false, "Display Usage")
	flag.Parse()

	// if -h flag is set, display usage information and exit
	if *help {
		fmt.Println("itinerary Usage: ")
		fmt.Println("go run . ./input.txt ./output.txt ./airport-lookup.csv")
		return
	}

	// check if exactly 3 arguments are provided
	args := flag.Args()
	if len(args) != 3 {
		fmt.Fprintln(os.Stderr, "Error: Provide exactly 3 arguments")
		fmt.Println("Usage : go run . ./input.txt ./output.txt ./airport-lookup.csv")
		os.Exit(1)
	}

	// Assign arguments to variables
	inputFile := args[0]
	outputFile := args[1]
	airportFile := args[2]

	// Process the itinerary
	err := processItinerary(inputFile, outputFile, airportFile)
	if err != nil {
		// Print error message to stderr and exit with non-zero status
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	fmt.Println("Processing completed successfully.")
}
