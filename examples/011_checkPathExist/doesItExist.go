package main

import (
	"fmt"
	"os"
)

// Inorder to read a file, you need to open it first and chekc it it exists
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide file path")
		return
	}

	// get the file path
	path := arguments[1]

	// check if file exists using os.Stat
	paths, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Path does not exit\n")
			os.Exit(1)
		}
	} else {
		fmt.Println("Details")
		fmt.Println("Path:", paths.Name())
	}

}
