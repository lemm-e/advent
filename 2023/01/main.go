package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	lines := strings.Split(input, "\n")

	partOne(lines)
	partTwo(lines)
}

func partOne(lines []string) {
	re := regexp.MustCompile(`\d`)

	sum := 0
	for _, v := range lines {
		matches := re.FindAllString(v, -1)
		num, err := strconv.Atoi(matches[0] + matches[len(matches)-1])
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}
	fmt.Println("Part 1:", sum)
}

func partTwo(lines []string) {
	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	nums := `one|two|three|four|five|six|seven|eight|nine`
	re := regexp.MustCompile(`(?i)` + nums + `|\d`)
	reRev := regexp.MustCompile(`(?i)` + Reverse(nums) + `|\d`)

	sum := 0
	for _, v := range lines {
		match := re.FindAllString(v, 1)
		matchRev := reRev.FindAllString(Reverse(v), 1)

		first := match[0]
		if len(first) > 1 {
			first = numMap[strings.ToLower(first)]
		}

		last := matchRev[0]
		if len(last) > 1 {
			last = numMap[Reverse(strings.ToLower(last))]
		}

		num, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}

		sum += num
	}
	fmt.Println("Part 2:", sum)
}

func Reverse(s string) string {
	out := ""
	for _, v := range s {
		out = string(v) + out
	}

	return out
}
