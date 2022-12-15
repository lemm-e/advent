package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	partOne(input)
	partTwo(input)
}

func partOne(p string) {
	s, m := parse(p)
	for _, v := range m {
		count, from, to := v[0], v[1], v[2]
		for i := 0; i < count; i++ {
			temp := s[from][len(s[from])-1]
			s[from] = s[from][:len(s[from])-1]
			s[to] = append(s[to], temp)
		}
	}

	var out string
	for i := 1; i <= len(s); i++ {
		out += s[i][len(s[i])-1]
	}

	fmt.Println("Part 1:", out)
}

func partTwo(p string) {
	s, m := parse(p)
	for _, v := range m {
		count, from, to := v[0], v[1], v[2]
		temp := s[from][len(s[from])-count:]
		s[from] = s[from][:len(s[from])-count]
		s[to] = append(s[to], temp...)
	}

	var out string
	for i := 1; i <= len(s); i++ {
		out += s[i][len(s[i])-1]
	}

	fmt.Println("Part 2:", out)
}

func parse(input string) (map[int][]string, [][3]int) {
	var parts = strings.Split(input, "\n\n")
	re := regexp.MustCompile(`\w+\s(\d+)\s\w+\s(\d+)\s\w+\s(\d+)`)
	var moves [][3]int
	for _, v := range re.FindAllStringSubmatch(parts[1], -1) {
		num, _ := strconv.Atoi(v[1])
		from, _ := strconv.Atoi(v[2])
		to, _ := strconv.Atoi(v[3])
		moves = append(moves, [3]int{num, from, to})
	}

	stacks := make(map[int][]string)
	lines := strings.Split(parts[0], "\n")
	count, _ := strconv.Atoi(string(parts[0][len(parts[0])-2]))
	for i := len(lines) - 1; i >= 0; i-- {
		v := lines[i]
		for j := 0; j < 4*count-1; j += 4 {
			if v[j] == '[' {
				stacks[j/4+1] = append(stacks[j/4+1], string(v[j+1]))
			}
		}
	}

	return stacks, moves
}