package main

import (
	"fmt"
	"log"
	"os"
)

// create file func
func createfile() {

	// /home/desmond/Documents/
	file, es := os.Create("/home/desmond/Documents/testfile/logfile.txt")
	if es != nil {
		log.Fatal(es)
	}
	defer file.Close()

	// capture len of string in file
	len, es := file.WriteString("welcome everyone" + "Program demonstrate write and read")
	if es != nil {
		log.Fatalf("failed writing file: %s", es)
	}

	fmt.Printf("\nFile name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

// read file
func readfile() {
	fmt.Println("Reading file ")

	filename := "logfile.txt"

	data, es := os.ReadFile("/home/desmond/Documents/testfile/logfile.txt")
	if es != nil {
		log.Panicf("Failed reading file %s", es)
	}
	fmt.Printf("\nFile name is : %s ", filename)
	fmt.Printf("\n\nData is : %s \n", data)
}

func main() {
	createfile()
	readfile()
}
