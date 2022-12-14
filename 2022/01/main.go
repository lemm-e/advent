package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	content := strings.Split(input, "\n")
	sum := getSum(content)

	partOne(sum)
	partTwo(sum)
}

func partOne(s []int) {
	fmt.Println("Part 1:", s[0])
}

func partTwo(s []int) {
	fmt.Println("Part 2:", s[0]+s[1]+s[2])
}

func getSum(inp []string) []int {
	var res []int

	var sum int
	for _, v := range inp {
		num, err := strconv.Atoi(v)
		if err != nil {
			res = append(res, sum)
			sum = 0
			continue
		}
		sum += num
	}

	if sum != 0 {
		res = append(res, sum)
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return res
}
