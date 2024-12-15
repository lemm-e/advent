package main

import (
	_ "embed"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type colorCount struct {
	red   int
	green int
	blue  int
}

type resultSet []colorCount

type dataMap map[int]resultSet

//go:embed input.txt
var input string

func main() {
	dmap := ParseData(input)
	resOne := partOne(dmap)
	resTwo := partTwo(dmap)

	fmt.Println("Part 1:", resOne)
	fmt.Println("Part 2:", resTwo)
}

func ParseData(s string) dataMap {
	lines := strings.Split(s, "\n")
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)

	var result dataMap = make(dataMap)
	for i, v := range lines {
		data := strings.Split(v, ": ")[1]
		picks := strings.Split(data, ";")

		var gameInfo resultSet
		for _, k := range picks {
			matches := re.FindAllStringSubmatch(k, -1)

			var pickInfo colorCount
			for _, v := range matches {
				num, err := strconv.Atoi(v[1])
				if err != nil {
					log.Fatal(err)
				}

				if v[2] == "red" {
					pickInfo.red = num
				} else if v[2] == "green" {
					pickInfo.green = num
				} else {
					pickInfo.blue = num
				}
			}
			gameInfo = append(gameInfo, pickInfo)
		}
		result[i+1] = gameInfo
	}

	return result
}

func partOne(d dataMap) int {
	var result int

id:
	for i := 1; i <= 100; i++ {
		for _, v := range d[i] {
			if v.red > 12 || v.green > 13 || v.blue > 14 {
				continue id
			}
		}
		result += i
	}

	return result
}

func partTwo(d dataMap) int {
	var result int

	for _, value := range d {
		var r, g, b int
		for _, v := range value {
			if v.red > r {
				r = v.red
			}
			if v.green > g {
				g = v.green
			}
			if v.blue > b {
				b = v.blue
			}
		}
		result += r * g * b
	}

	return result
}
