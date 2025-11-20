package main

import (
	"flag"
	"fmt"
	"os"
)

// main func
func main() {

	// parse flags
	flag.Parse()
	// Initilize flags for cli arguments
	arg := flag.Args()
	if len(arg) < 1 {
		fmt.Printf("Usage: go run . '[5 #]' ")
		fmt.Println("Example: go run . '[5 #][5 -_]-[5 #]'")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// let set input to first arguments
	input := arg[0]
	result, err := decoder(input)
	if err != nil {
		fmt.Fprintln(os.Stderr, "decode error:", err)
		os.Exit(1)
	}
	fmt.Println(result)
}

// help flag func
func displayHelp() {
	fmt.Println("Usage:")
	fmt.Println("  go run . '[5 $]'               → Decode single input")
	fmt.Println("  go run . -encode 'argument'     → Encode single input")
	fmt.Println("  go run . -multiline             → Multi-line decode (add done)")
	fmt.Println("  go run . -encode -multiline    → Multi-line encode")
	fmt.Println("  go run . -file=input.txt        → Decode file")
	fmt.Println("  go run . -encode -file=input.txt → Encode file")
}
