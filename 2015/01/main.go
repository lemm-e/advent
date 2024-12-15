package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	fmt.Println("Part 1:", strings.Count(input, "(")-strings.Count(input, ")"))
}

func partTwo() {
	var floor int
	for i, v := range input {
		if fmt.Sprintf("%c", v) == "(" {
			floor++
		} else {
			floor--
		}

		if floor == -1 {
			fmt.Println("Part 2:", i+1)
			break
		}
	}
}
