package main

import "fmt"

// This program using structure pointer

// defining structure

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	// passing address of struct -var to empy is a pointer
	empy := &Employee{"Rabia Tu", "Zimbala", 45, 53000}

	// to call first name of empyee struct
	// (*empy).firstName is the syntax to get the value of firstname

	fmt.Println("First name: ", (*empy).firstName)
	fmt.Println("Last name: ", (*empy).lastName)
	fmt.Println("Age: ", (*empy).age)
}
