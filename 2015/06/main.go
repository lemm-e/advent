package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

type command struct {
	ins            string
	x1, y1, x2, y2 int
}

func main() {
	parsed := parse()

	partOne(parsed)
	partTwo(parsed)
}

func partOne(commands []command) {
	board := make([][]int, 1000)
	for i := range board {
		board[i] = make([]int, 1000)
	}

	for _, cmd := range commands {
		for x := cmd.x1; x <= cmd.x2; x++ {
			for y := cmd.y1; y <= cmd.y2; y++ {
				switch cmd.ins {
				case "toggle":
					if board[x][y] == 0 {
						board[x][y] = 1
					} else {
						board[x][y] = 0
					}
				case "on":
					board[x][y] = 1
				case "off":
					board[x][y] = 0
				}
			}
		}
	}

	fmt.Println("Part 1:", evaluate(board))
}

func partTwo(commands []command) {
	board := make([][]int, 1000)
	for i := range board {
		board[i] = make([]int, 1000)
	}

	for _, cmd := range commands {
		for x := cmd.x1; x <= cmd.x2; x++ {
			for y := cmd.y1; y <= cmd.y2; y++ {
				switch cmd.ins {
				case "toggle":
					board[x][y] += 2
				case "on":
					board[x][y]++
				case "off":
					if board[x][y] > 0 {
						board[x][y]--
					}
				}
			}
		}
	}

	fmt.Println("Part 2:", evaluate(board))
}

func parse() []command {
	re := regexp.MustCompile(`(?:(toggle)|turn (on|off)) (\d+),(\d+) through (\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	var commands []command

	for _, match := range matches {
		var cmd command

		ins := match[1]
		if ins == "" {
			ins = match[2]
		}
		x1, _ := strconv.Atoi(match[3])
		y1, _ := strconv.Atoi(match[4])
		x2, _ := strconv.Atoi(match[5])
		y2, _ := strconv.Atoi(match[6])

		cmd = command{ins, x1, y1, x2, y2}
		commands = append(commands, cmd)
	}

	return commands
}

func evaluate(board [][]int) int {
	count := 0
	for _, row := range board {
		for _, cell := range row {
			count += cell
		}
	}

	return count
}