package main

import (
	"fmt"
)

func main() {
	// Program to demonstrate misc operators
	// & and * operators

	// declare variable
	var a int = 58

	// & operator returns the address of variable a
	z := &a
	fmt.Println("To get value of a print z:", z)

	// * operator
	fmt.Println("To get value of a print *z:", *z)
	
}
