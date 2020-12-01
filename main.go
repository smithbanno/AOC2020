package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

var target int = 2020

func readFile() []int {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var perline int
	var nums []int

	for {
		_, err := fmt.Fscanf(file, "%d\n", &perline) // give a patter to scan

		if err != nil {
			if err == io.EOF {
				break // stop reading the file
			}
			fmt.Println(err)
			os.Exit(1)
		}
		nums = append(nums, perline)
	}
	return nums
}

func twoSum2(nums []int, target int) (int, int) {
	if len(nums) <= 1 {
		return -1, -1
	}
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return j, i
		}
		m[target-v] = i
	}
	return -1, -1
}

func threeSum(nums []int) [][]int {
	var uniqueTriplets [][]int

	// Create a map with our triplets being the key and a bool being the value
	// Note that Golang will only allow a slice to be the key if the slice is given a concrete size
	tripletSet := make(map[[3]int]bool)
	for a := 0; a < len(nums)-2; a++ {
		for b := a + 1; b < len(nums)-1; b++ {
			for c := b + 1; c < len(nums); c++ {
				if nums[a]+nums[b]+nums[c] == target {
					// When we find a sum 2020, create an unsized slice to allow easy sorting
					triplet := []int{nums[a], nums[b], nums[c]}
					// Sort the three numbers
					sort.Ints(triplet)
					// Convert the sorted slice into the sized slice that can be used as a key in our map
					sizedTriplet := [3]int{triplet[0], triplet[1], triplet[2]}
					// Check if the entry already exists in the map before adding it to our results
					_, ok := tripletSet[sizedTriplet]
					if !ok {
						tripletSet[sizedTriplet] = true
						uniqueTriplets = append(uniqueTriplets, triplet)
					}
				}
			}
		}
	}
	return uniqueTriplets
}

func main() {
	nums := readFile()

	a, b := twoSum2(nums, target)
	fmt.Print("Day 1 task 1 answer is: ")
	fmt.Println(nums[a] * nums[b])

	tsn := (threeSum(nums))
	for _, v := range tsn {
		z := v[0]
		y := v[1]
		x := v[2]
		answer := z * y * x
		fmt.Print("Day 1 task 2 answer is: ")
		fmt.Println(answer)
	}
}
