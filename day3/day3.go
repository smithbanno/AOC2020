package day3

import (
	"strings"
)

// Vec is a struct
type Vec struct {
	x int
	y int
}

// Answer1 computes the first answer for AOC day 3
func Answer1(input string, x int, y int) int {
	input = strings.Replace(input, "\n", "", -1)
	pos := Vec{0, 0}
	trees := 0
	for pos.y*31+pos.x < len(input)-1 {
		if string(input[pos.y*31+pos.x]) == "#" {
			trees++
		}
		pos.x += x
		pos.y += y
		if pos.x > 30 {
			pos.x -= 31
		}
	}
	return trees
}

// Answer2 computes the second answer for AOC day 3
func Answer2(input string) int {
	a := Answer1(input, 1, 1)
	b := Answer1(input, 3, 1)
	c := Answer1(input, 5, 1)
	d := Answer1(input, 7, 1)
	e := Answer1(input, 1, 2)
	return a * b * c * d * e
}
