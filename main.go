package main

import (
	"fmt"

	one "github.com/smithbanno/aoc2020/day1"
	two "github.com/smithbanno/aoc2020/day2"
)

var target int = 2020

func main() {
	// Day 1
	nums := one.ReadFile()

	a, b := one.TwoSum(nums, target)
	fmt.Print("Day 1 task 1 answer is: ")
	fmt.Println(nums[a] * nums[b])

	tsn := (one.ThreeSum(nums))
	for _, v := range tsn {
		z := v[0]
		y := v[1]
		x := v[2]
		answer := z * y * x
		fmt.Print("Day 1 task 2 answer is: ")
		fmt.Println(answer)
	}

	// Day 2
	two.CheckValidity1()
	two.CheckValidity2()
}
