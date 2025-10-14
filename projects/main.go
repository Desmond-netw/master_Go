package main

// imports
import (
	"flag"
	"fmt"
	"os"
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

	//Load airport lookup data from CSV into a map

}
