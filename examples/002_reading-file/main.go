package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Applicatio read file
// step
// 1. get the user input filename
// 2. using os.open to open the file
// 3. scann the content of the open file
// 4. print the scan content

func main() {

	// reader var
	reader := bufio.NewReader(os.Stdin)
	// Print prompt
	fmt.Println("Enter the fileanme : \n")
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1]

	path := filepath.Join("temp", filename)
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		fmt.Fprintf(os.Stderr, "Error opening file")
		os.Exit(1)
	}
	defer file.Close()

}
