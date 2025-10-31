package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// function to create file
/*
Algorithm
1. Create File and Add content
- get file name
-get content to put on file
- check for err creating file and writing file
-
2. Readfile content
- get file name
-open filename
- read file content
- check for err reading file
*/

func createF(fname, content string) {
	// using create os.create to create a file

	path := filepath.Join("temp", fname)
	file, err := os.Create(path) // the path help to handle relative file paths
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// write content to file
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\nError writing content to file  %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n\v\tFile '%v' created sucessfully")
	fmt.Printf("\n\v\tContent added sucessfule")
}

// function to read file

func readF(fname string) {
	// reading file using file name
	fmt.Println("Reading File .....")

	// opening file first
	path := filepath.Join("temp", fname)
	file, es := os.Open(path) // open a file using the file path
	if es != nil {
		fmt.Fprintf(os.Stderr, "Error openning file %v\n", es)
		os.Exit(1)
	}
	defer file.Close()

	//reading file content
	data, es := io.ReadAll(file) // read all file content
	if es != nil {
		fmt.Fprintf(os.Stderr, "Error reading file content %v\n", es)
		os.Exit(1)
	}

	fmt.Printf("Successfully open file :%s \n")
	fmt.Printf("content \n %s \n", string(data))
}

func main() {
	// let initilize our reader
	reader := bufio.NewReader(os.Stdin)

	// prompting
	fmt.Println("Welcome to Terminal App")

	fmt.Println("Enter file name : ")
	fname, _ := reader.ReadString('\n')
	fname = fname[:len(fname)-1] // remove spaces in filename

	// taking input for content
	fmt.Println("Entering file content")
	fcontent, _ := reader.ReadString('\n')

	createF(fname, fcontent)
	readF(fname)

}
