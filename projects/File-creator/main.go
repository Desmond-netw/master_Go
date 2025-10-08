package main

import (
	"bufio"
	"fmt"
	"os"
)

// using bufio  reader for os.stdin

func main() {
	// delcare reader
	reader := bufio.NewReader(os.Stdin)

	//header intro
	fmt.Println("Welcome to Text Base content creating ")

	fmt.Println("Enter Filename ")
	// initila variables
	filename, _ := reader.ReadString('\n')
	filename = filename[:len(filename)-1] // removing white line space

}
