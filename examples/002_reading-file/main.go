package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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
	fmt.Println("Enter the fileanme :")
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1]

	path := filepath.Join("./", filename)
	// open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer file.Close()

	// scan file
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file %v\n", err)
		os.Exit(1)
	}

	// String manipulations
	trimmed := strings.TrimSpace(content)
	toUpper := strings.ToUpper(trimmed)
	// Printing output
	fmt.Println("\n Original Ouput")
	fmt.Println(content)

	// printing manipulations
	fmt.Println("\n Manipulated String contents")
	fmt.Println(trimmed)
	fmt.Println("converted to Upper case")
	fmt.Println(toUpper)
}
