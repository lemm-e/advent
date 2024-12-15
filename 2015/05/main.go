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
	var nice int
	for _, s := range strings.Split(input, "\n") {
		if !(strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") || strings.Count(s, "a")+strings.Count(s, "e")+strings.Count(s, "i")+strings.Count(s, "o")+strings.Count(s, "u") < 3 || !letterRepeat(s, 0)) {
			nice++
		}
	}
	fmt.Println("Part 1:", nice)
}

func partTwo() {
	var nice int
	for _, s := range strings.Split(input, "\n") {
		if nonOverlappingPair(s) && letterRepeat(s, 1) {
			nice++
		}
	}
	fmt.Println("Part 2:", nice)
}

func letterRepeat(s string, skip int) bool {
	for i := 0; i < len(s) - skip - 1; i++ {
		if s[i] == s[i+1+skip] {
			return true
		}
	}
	return false
}

func nonOverlappingPair(s string) bool {
	for i := 0; i < len(s) - 1; i++ {
		if strings.Count(s, s[i:i+2]) > 1 {
			return true
		}
	}
	return false
}
