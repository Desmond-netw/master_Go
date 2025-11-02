package main

import (
	"fmt"
	"os"
	"strings"
)

// define const for artheight of each char
const (
	artHeight = 8
)

// asciiArtFont is a var that hold the 8-line height of ASCII uppercase
// each char is 8 character wide
var asciiArtFont = map[rune][]string{
	'A': {
		"   ##   ",
		"  #  #  ",
		" #    # ",
		" ###### ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
	},
	'B': {
		" #####  ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #####  ",
	},
	'C': {
		"  ##### ",
		" #     #",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #     #",
		"  ##### ",
	},
	'D': {
		" #####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #####  ",
	},
	'E': {
		" #######",
		" #      ",
		" #      ",
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		" #######",
	},
	'F': {
		" #######",
		" #      ",
		" #      ",
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
	},
	'G': {
		"  ##### ",
		" #     #",
		" #      ",
		" #  ####",
		" #     #",
		" #     #",
		" #     #",
		"  ##### ",
	},
	'H': {
		" #    # ",
		" #    # ",
		" #    # ",
		" ###### ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
	},
	'I': {
		" #######",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		" #######",
	},
	'J': {
		"       #",
		"       #",
		"       #",
		"       #",
		"       #",
		" #     #",
		" #     #",
		"  ##### ",
	},
	'K': {
		" #    # ",
		" #   #  ",
		" #  #   ",
		" ###    ",
		" #  #   ",
		" #   #  ",
		" #    # ",
		" #    # ",
	},
	'L': {
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #      ",
		" #######",
	},
	'M': {
		" #    # ",
		" ##  ## ",
		" # ## # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #    # ",
	},
	'N': {
		" #    # ",
		" ##   # ",
		" # #  # ",
		" #  # # ",
		" #   ## ",
		" #    # ",
		" #    # ",
		" #    # ",
	},
	'O': {
		"  ##### ",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		"  ##### ",
	},
	'P': {
		" #####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #      ",
		" #      ",
		" #      ",
	},
	'Q': {
		"  ##### ",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #  #  #",
		" #   # #",
		"  #### #",
	},
	'R': {
		" #####  ",
		" #    # ",
		" #    # ",
		" #    # ",
		" #####  ",
		" #   #  ",
		" #    # ",
		" #    # ",
	},
	'S': {
		"  ##### ",
		" #     #",
		" #      ",
		"  ##### ",
		"       #",
		" #     #",
		" #     #",
		"  ##### ",
	},
	'T': {
		" #######",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
	},
	'U': {
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		"  ##### ",
	},
	'V': {
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		"  #   # ",
		"   # #  ",
		"    #   ",
	},
	'W': {
		" #     #",
		" #     #",
		" #     #",
		" #     #",
		" #  #  #",
		" # ## # ",
		"  #  #  ",
		"  #  #  ",
	},
	'X': {
		" #     #",
		"  #   # ",
		"   # #  ",
		"    #   ",
		"   # #  ",
		"  #   # ",
		" #     #",
		" #     #",
	},
	'Y': {
		" #     #",
		"  #   # ",
		"   # #  ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
		"    #   ",
	},
	'Z': {
		" #######",
		"       #",
		"      # ",
		"     #  ",
		"    #   ",
		"   #    ",
		"  #     ",
		" #######",
	},
	' ': {
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
		"        ",
	},
	'?': {
		"  ##### ",
		" #     #",
		" #     #",
		"    #   ",
		"   #    ",
		"   #    ",
		"        ",
		"   #    ",
	},
}

/*
-1. Get the arg flags
-2. converts inputText to uppercase
-3.
*/

func main() {
	// get input argument from command line including programname
	if len(os.Args) < 2 {
		fmt.Println("usage: go run main.go \"Your Text\" ")
		fmt.Println("Exampel: go run asciiArtFont.go \"Hello Word\"")
		return
	}
	// get inputText by joining all arguments
	input := os.Args[1:]
	inputText := strings.Join(input, "")
	// outputlines will have artHeight number of lines
	// for the final output
	outputLines := make([]string, artHeight)
	// convert inputText
	inputText = strings.ToUpper(inputText)

	// 3: iterate throght each char
	for _, char := range inputText {
		// lookup the ASCII art
		art, exists := asciiArtFont[char]
		if !exists {
			art = asciiArtFont['?']
		}

		// 6. Append each line of the character's
		for i := 0; i < artHeight; i++ {
			outputLines[i] += art[i]
		}
	}

	// Print the final, combined output
	for _, line := range outputLines {
		fmt.Println(line)
	}

}
