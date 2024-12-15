package main

import (
	"crypto/md5"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	partOne()
	partTwo()
}

func partOne() {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", strings.TrimSpace(input), i)))
		a := fmt.Sprintf("%x", hash)
		if a[:5] == "00000" {
			fmt.Println("Part 1:", i)
			break
		}
	}
}

func partTwo() {
	for i := 0; ; i++ {
		hash := md5.Sum([]byte(fmt.Sprintf("%s%d", strings.TrimSpace(input), i)))
		a := fmt.Sprintf("%x", hash)
		if a[:6] == "000000" {
			fmt.Println("Part 1:", i)
			break
		}
	}
}
