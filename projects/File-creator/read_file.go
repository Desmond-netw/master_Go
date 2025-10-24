package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func readFile(filename string) {
	fmt.Println("\nReading from file ......")

	// Open function for reading
	// for proper file path handling in all OS
	path := filepath.Join("temp", filename)
	myfile, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file : %v\n", err)
		os.Exit(1)
	}
	defer myfile.Close()

	// read the file content
	data, err := io.ReadAll(myfile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %V\n", err)
		log.Fatal(err)
		os.Exit(1)
	}

	// success
	fmt.Printf("Content of '%s': \n%s\n", filename, string(data))
}
