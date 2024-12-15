package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

type point struct {
	x, y int
}

type agent struct {
	name string
	curr point
}

func (a *agent) move(c byte) {
	switch c {
	case '^':
		a.curr.y++
	case 'v':
		a.curr.y--
	case '>':
		a.curr.x++
	case '<':
		a.curr.x--
	}
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	visits := make(map[point]int)
	santa := agent{"Santa", point{0, 0}}

	visits[santa.curr]++
	for _, c := range input {
		santa.move(byte(c))
		visits[santa.curr]++
	}

	fmt.Println("Part 1:", len(visits))
}

func partTwo() {
	visits := make(map[point]int)
	santa := agent{"Santa", point{0, 0}}
	robo := agent{"Robo", point{0, 0}}

	visits[santa.curr]++
	visits[robo.curr]++
	for i, c := range input {
		if i%2 == 0 {
			santa.move(byte(c))
			visits[santa.curr]++
		} else {
			robo.move(byte(c))
			visits[robo.curr]++
		}
	}

	fmt.Println("Part 2:", len(visits))
}
