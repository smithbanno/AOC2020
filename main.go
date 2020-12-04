package main

import (
	"fmt"
	"io/ioutil"

	one "github.com/smithbanno/aoc2020/day1"
	two "github.com/smithbanno/aoc2020/day2"
	three "github.com/smithbanno/aoc2020/day3"
)

var target int = 2020

func main() {
	// Day 1
	fmt.Println("Day 1 answers:")
	nums := one.ReadFile()

	a, b := one.TwoSum(nums, target)
	fmt.Print("Task 1 answer is ")
	fmt.Println(nums[a] * nums[b])

	tsn := (one.ThreeSum(nums))
	for _, v := range tsn {
		z := v[0]
		y := v[1]
		x := v[2]
		answer := z * y * x
		fmt.Print("Task 2 answer is ")
		fmt.Println(answer)
	}

	// Day 2
	fmt.Println("Day 2 answers:")
	two.CheckValidity1()
	two.CheckValidity2()

	//Day 3
	input, err := ioutil.ReadFile("day3.txt")
	if err != nil {
		panic(err)
	}
	trees := three.Answer1(string(input), 3, 1)
	fmt.Println("Day 3 answers:")
	fmt.Printf("Task 1 answer is %d\n", trees)
	multiTree := three.Answer2(string(input))
	fmt.Printf("Task 2 answer is %d\n", multiTree)
}
