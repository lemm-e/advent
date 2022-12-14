package main

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
)

type shapes [3]byte

//go:embed input.txt
var fs embed.FS
var shapeOpp shapes = shapes{'A', 'B', 'C'}
var shapeMe shapes = shapes{'X', 'Y', 'Z'}

func main() {
	input, err := fs.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file. error:", err)
	}
	defer input.Close()

	parsed := Parse(input)

	partOne(parsed)
	partTwo(parsed)
}

func partOne(moves [][2]int) {
	var score int

	for _, v := range moves {
		score += v[1] + 1

		if (v[0]+1)%3 == v[1] {
			score += 6
		} else if v[0] == v[1] {
			score += 3
		}
	}

	fmt.Println("Part 1:", score)
}

func partTwo(moves [][2]int) {
	var score int

	for _, v := range moves {
		score += v[1] * 3

		switch v[1] {
		case 0:
			if v[0]-1 < 0 {
				score += 3
			} else {
				score += (v[0]-1)%3 + 1
			}
			break
		case 1:
			score += v[0] + 1
			break
		case 2:
			score += (v[0]+1)%3 + 1
			break
		}
	}

	fmt.Println("Part 2:", score)
}

func Parse(i io.Reader) [][2]int {
	var moves [][2]int

	reader := bufio.NewScanner(i)
	for reader.Scan() {
		text := reader.Text()
		inp := [2]byte{text[0], text[2]}
		moves = append(moves, [2]int{shapeOpp.IndexOf(inp[0]), shapeMe.IndexOf(inp[1])})
	}

	return moves
}

func (s *shapes) IndexOf(b byte) int {
	for i, v := range s {
		if v == b {
			return i
		}
	}

	// never occurs
	return -1
}
