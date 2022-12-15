package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

func main() {
	partOne(input)
	partTwo(input)
}

func partOne(input string) {
	fmt.Println("Part 1:", distinctChars(input, 4))
}

func partTwo(input string) {
	fmt.Println("Part 2:", distinctChars(input, 14))
}

func distinctChars(input string, n int) int {
	var ret int = -1
	for i := 0; i < len(input)-n; i++ {
		count := make(map[byte]int, n)
		for j := i; j < i+n; j++ {
			count[input[j]]++
		}

		var flag bool
		for _, v := range count {
			if v != 1 {
				flag = true
				break
			}
		}

		if !flag {
			return i+n
		}
	}

	return ret
}