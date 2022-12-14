package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	rucksacks := strings.Split(input, "\n")

	partOne(rucksacks)
	partTwo(rucksacks)
}

func partOne(sacks []string) {
	var score int
	for _, sack := range sacks {
		a, b := sack[:len(sack)/2], sack[len(sack)/2:]
		for _, char := range a {
			inx := strings.Index(b, string(char))
			if inx != -1 {
				score += getPriority(char)
				break
			}
		}
	}

	fmt.Println("Part 1:", score)
}

func partTwo(sacks []string) {
	var score int

	for i := 0; i < len(sacks)-2; i += 3 {
		a, b, c := sacks[i], sacks[i+1], sacks[i+2]
		for _, char := range a {
			inxB := strings.Index(b, string(char))
			inxC := strings.Index(c, string(char))
			if inxB != -1 && inxC != -1 {
				score += getPriority(char)
				break
			}
		}
	}

	fmt.Println("Part 2:", score)
}

func getPriority(c rune) int {
	var priority int
	if string(c) == strings.ToLower(string(c)) {
		priority = int(c) - 96
	} else {
		priority = int(c) - 38
	}
	return priority
}
