package main

import "fmt"

// create, and initialize the maps
//import package

func main() {
	// Creating  and initializing empty map.
	var map_1 map[int]int
	// check if the map is nil or not
	if map_1 == nil {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	//Using shorthand-declaration
	// and using map literals
	map_2 := map[int]string{
		90: "Hen",
		91: "Cow",
		93: "Cat",
	}
	fmt.Println("Map[int]string - key :  values")
	fmt.Println("Map 2", map_2)
}
