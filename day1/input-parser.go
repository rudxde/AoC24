package main

import (
	"fmt"
	"os"
	"strconv"
)

type Input struct {
	left  []int
	right []int
}

func parseInput(filepath string) (*Input, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	lines := countLines(&dat)

	left := make([]int, lines)
	right := make([]int, lines)

	filePositionCounter := 0
	for i := 0; i < lines; i++ {
		startSecondNumber := -1
		startOfLine := filePositionCounter
		for ; ; filePositionCounter++ {
			if dat[filePositionCounter] == '\n' {
				return nil, fmt.Errorf("error parsing left input in line %d: Unexpected end of line", i)
			}
			if dat[filePositionCounter] == ' ' {

				leftVal, err := strconv.Atoi(string(dat[startOfLine:filePositionCounter]))

				if err != nil {
					return nil, fmt.Errorf("error parsing left input in line %d: %w", i, err)
				}

				left[i] = leftVal

				filePositionCounter += 3 // expect second space
				startSecondNumber = filePositionCounter
				break
			}
		}

		for ; dat[filePositionCounter] != '\n'; filePositionCounter++ {
		}

		rightVal, err := strconv.Atoi(string(dat[startSecondNumber:filePositionCounter]))

		if err != nil {
			return nil, fmt.Errorf("error parsing right input in line %d: %w", i, err)
		}

		right[i] = rightVal
		filePositionCounter++
	}

	result := Input{
		left:  left,
		right: right,
	}

	return &result, nil
}

func countLines(file *[]byte) int {
	result := 0

	for i := 0; i < len(*file); i++ {
		if (*file)[i] == '\n' {
			result++
		}
	}

	return result
}
