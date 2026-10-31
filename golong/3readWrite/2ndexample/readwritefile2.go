package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 1st func to creat file with os.create method

func createfile(filename, text string) {
	// print output
	fmt.Printf("\t\vWriting a file in Go \n")

	// create file
	file, es := os.Create(filename)
	if es != nil {
		log.Fatalf("Failed creating file %s", es)
	}
	defer file.Close()

	// start writing text to file
	// using writestring method

	len, es := file.WriteString(text)
	if es != nil {
		log.Panicf("writing file fail %s", es)
	}

	fmt.Printf("\nFileName: %s", file.Name())
	fmt.Printf("\nFile len: %d bytes", len)
}

func readfile(filename string) {
	// print out header

	fmt.Printf("\t\vReading Files\n")

	data, es := os.ReadFile(filename)
	if es != nil {
		log.Fatalf("\nReading file fail %s", es)
	}

	// base file info
	fmt.Printf("\nFilename : %s", filename)
	fmt.Printf("\nFilesize: %d bytes", len(data))
	fmt.Printf("\t\vData :\n %s \n", data)
}

// call all funcs

func main() {

	// grap dyname info from cli inputs
	// using the bufio reader and os.stdin
	var filename string
	fmt.Println("Enter file name")
	fmt.Scanln(&filename)

	// use input
	fmt.Printf("\nEnter your file content \n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	createfile(filename, input)

	readfile(filename)
}
