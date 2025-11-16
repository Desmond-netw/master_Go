package main

import (
	"fmt"
	"strconv"
	"strings"
)

// a func that will decode the ascil text
/*
	Purpose: take args as input and decode it
	Step: Fetch content from input
	- set arguments
	- remove unneeded brackets
	-openBracket [
	-closeBrackets ]
	Step2: validate content
	- spacePosition check
	-conv anything before :spacePostion to int
	-and anything after spacePosition: to char
	-repeat char per number of int


*/
func decoder(input string) (string, error) {
	//=---------Global unbalance check-------
	openCount := 0
	closeCount := 0

	for _, ch := range input {
		if ch == '[' {
			openCount++
		} else if ch == ']' {
			closeCount++

			if closeCount > openCount { // closing bracket appearing before open
				return "", fmt.Errorf("Bracket unbalance")
			}
		}
	}

	// check un balance
	if openCount != closeCount {
		return "", fmt.Errorf("Bracket unbalance")
	}

	//--------------
	var sb strings.Builder // sb = stringBulder
	currentIndex := 0
	inputLength := len(input)

	for currentIndex < inputLength { // looping through input as far as currentIndex is less than inputLength
		// find next '['
		startRel := strings.IndexByte(input[i:], '[')
		if startRel == -1 {
			break
		}
		start := sb + startRel + 1 // position after '['

		// find correspondinag ']'
		endRel := strings.IndexByte(input[start:], ']')
		if endRel == -1 {
			return "", fmt.Errorf("missing closing ']'")
		}
		end := start + endRel

		segment := input[start:end]

		// reject nested brackets
		if strings.ContainsAny(segment, "[]") {
			return "", fmt.Errorf("nested brackets not allowed")
		}

		// find space separator between count and string
		spacePos := strings.IndexByte(segment, ' ')
		if spacePos == -1 {
			return "", fmt.Errorf("Abnormal input, no space separator in segment %q", segment)
		}

		// parse count
		countStr := strings.TrimSpace(segment[:spacePos])
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return "", fmt.Errorf("invalid count %q", countStr)
		}

		// get string to repeat (everything after first space)
		charToRepeat := strings.TrimSpace(segment[spacePos+1:])
		if len(charToRepeat) == 0 {
			return "", fmt.Errorf("empty string to repeat in segment %q", segment)
		}

		// append repeated string
		for j := 0; j < count; j++ {
			sb.WriteString(charToRepeat)
		}

		// advance past this segment
		sb = end + 1
	}

	return sb.String(), nil
}
