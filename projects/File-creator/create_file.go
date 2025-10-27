package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Creat File function
// it takes 2 arguments >> filename ja content
func createFile(filename, content string) {
	// printing operations
	fmt.Println("Writing to file ......")

	// using os.Create to truncate the file
	path := filepath.Join("temp", filename) // filepath Join to store the path
	myfile, es := os.Create(path)
	if es != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %v\n", es)
		log.Fatal(es)
		os.Exit(1)
	}
	defer myfile.Close() // this defer gaurantee the closer of the application

	// write content to the file
	_, es = myfile.WriteString(content)
	if es != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file %v\n", es)
		log.Fatal(es)
		os.Exit(1)
	}
	// success message
	fmt.Printf("\t\v\nFile %s Writen successfully (%s bytes)", filename, myfile)
}
