package main

import (
	"log"
	"os"
)

func main() {

	// Empty file creating in Go lang
	newFile, err := os.Create("heyy.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile.Name())
	newFile.Close()

}
