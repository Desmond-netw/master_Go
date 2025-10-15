package main

import (
	"flag"
	"fmt"
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

}
