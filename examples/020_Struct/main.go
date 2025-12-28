package main

// Examples of Structure
import "fmt"

type Address struct {
	name     string
	streetNo string
	city     string
	state    string
	pin_code int
}

// Declaring variable of the struct type .
// All struct fields initialized value is set to zero
func main() {

	var x Address

	fmt.Printf("%v\n", x) // This will print { 0}
	// Delcaring and initilizing struct using struct literal
	x1 := Address{
		name:     "John Petter",
		streetNo: "Bonyan",
		city:     "Techiman",
		state:    "Bono",
		pin_code: 12321,
	}

	fmt.Printf("my Details %v\n", x1) // this print the value of struct

	// Using pointer
	addr := &Address{name: "Filip", city: "Accra", pin_code: 2025}
	println("Printing Pointer to a struct")
	// addr is an address to Address
	// so to get the value of address use *addr
	fmt.Println("First Name: ", (*addr).name)
	fmt.Println("City: ", (*addr).city)
	fmt.Println("pin_code: ", (*addr).pin_code)
}
