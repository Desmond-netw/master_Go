package main

import (
	"log"
	"os"
)

func main() {
	if es := os.Mkdir("folder", os.ModePerm); es != nil {
		log.Fatal(es)
	}
}
