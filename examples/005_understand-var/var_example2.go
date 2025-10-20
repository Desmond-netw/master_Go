package main

/*
var (
)
      nme1 = initialvalue1
      nme2 = initialvalue2
 T
 he following program declares variables of various kinds using the syn
tax described above.
*/

import "fmt"

func main() {
	// Declare multiple variables using a var block
	var (
		name      string  = "Desmond"
		age       int     = 30
		height    float64 = 5.7
		isStudent bool    = true || false
	)

	//Print the value of the variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)
}
