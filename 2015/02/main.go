package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type box struct {
	length int
	width  int
	height int
}

func main() {
	boxes := parseInput(input)

	partOne(boxes)
	partTwo(boxes)
}

func parseInput(input string) []box {
	dims := strings.Split(input, "\n")
	var boxes []box
	for _, v := range dims {
		nums := strings.Split(v, "x")

		l, _ := strconv.Atoi(nums[0])
		w, _ := strconv.Atoi(nums[1])
		h, _ := strconv.Atoi(nums[2])

		boxes = append(boxes, box{l, w, h})
	}

	return boxes
}

func partOne(b []box) {
	total := 0

	for _, v := range b {
		lw := v.length * v.width
		wh := v.width * v.height
		hl := v.height * v.length

		smallest := lw
		for _, v := range []int{lw, wh, hl} {
			if v < smallest {
				smallest = v
			}
		}

		total += 2*lw + 2*wh + 2*hl + smallest
	}

	fmt.Println("Part 1:", total)
}

func partTwo(b []box) {
	total := 0

	for _, v := range b {
		lw := 2 * (v.length + v.width)
		wh := 2 * (v.width + v.height)
		hl := 2 * (v.height + v.length)

		smallest := lw
		for _, v := range []int{lw, wh, hl} {
			if v < smallest {
				smallest = v
			}
		}

		total += v.length*v.width*v.height + smallest
	}

	fmt.Println("Part 2:", total)
}
