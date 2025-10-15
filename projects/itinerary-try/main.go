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

}
