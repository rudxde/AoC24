package main

import (
	"fmt"
	"slices"
)

func main() {
	input, err := parseInput("input.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	leftSorted := input.left
	slices.Sort(leftSorted)
	rightSorted := input.right
	slices.Sort(rightSorted)

	result1 := calculateDistances(leftSorted, rightSorted)

	fmt.Printf("Part 1: The total distance is: %d\n", result1)
	result2 := calculateWeirdOccurrences(leftSorted, rightSorted)

	fmt.Printf("Part 2: The similarity score is: %d\n", result2)

}

func calculateDistances(leftSorted []int, rightSorted []int) int {
	if len(leftSorted) != len(rightSorted) {
		panic(fmt.Errorf("length of left and right don't match! (left: %d right: %d)", len(leftSorted), len(rightSorted)))
	}
	result := 0

	for i := 0; i < len(leftSorted); i++ {
		distance := leftSorted[i] - rightSorted[i]
		if distance > 0 {
			result += distance
		} else {
			result -= distance
		}
	}
	return result
}

func calculateWeirdOccurrences(leftSorted []int, rightSorted []int) int {
	result := 0

	for i := 0; i < len(leftSorted); i++ {
		occurrences := 0
		for j := 0; j < len(rightSorted); j++ {
			if rightSorted[j] == leftSorted[i] {
				occurrences++
			}
		}
		result += occurrences * leftSorted[i]
	}
	return result
}
