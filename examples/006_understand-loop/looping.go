package main

import (
	"fmt"
)

func main() {
	rvariable := []string{"Hi", "coding in", "Go lang"}

	for indx, values := range rvariable {
		fmt.Printf("index %d = %v\n", indx, values)
	}
}
