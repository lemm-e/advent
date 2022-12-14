package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	var pairs = strings.Split(input, "\n")

	parsed := parse(pairs)
	partOne(parsed)
	partTwo(parsed)
}

func partOne(spaces [][2][2]int) {
	var count int
	for _, v := range spaces {
		if overlapFull(v) {
			count++
		}
	}

	fmt.Println("Part 1:", count)
}

func partTwo(spaces [][2][2]int) {
	var count int
	for _, v := range spaces {
		if overlapPartial(v) {
			count++
		}
	}

	fmt.Println("Part 2:", count)
}

func overlapFull(d [2][2]int) bool {
	return (d[0][0] <= d[1][0] && d[0][1] >= d[1][1]) || (d[1][0] <= d[0][0] && d[1][1] >= d[0][1])
}

func overlapPartial(d [2][2]int) bool {
	return (d[1][0] <= d[0][0] && d[0][0] <= d[1][1]) || (d[0][0] <= d[1][0] && d[1][0] <= d[0][1])
}

func parse(s []string) [][2][2]int {
	var parsed [][2][2]int

	for _, line := range s {
		var pairs [2][2]int
		ranges := strings.Split(line, ",")
		for i, v := range ranges {
			limits := strings.Split(v, "-")

			low, _ := strconv.Atoi(limits[0])
			high, _ := strconv.Atoi(limits[1])

			pairs[i] = [2]int{low, high}
		}
		parsed = append(parsed, pairs)
	}

	return parsed
}