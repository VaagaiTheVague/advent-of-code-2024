package day_04

import (
	"fmt"
	"strconv"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	count := CountXMAS(input)
	return strconv.Itoa(count)
}

func CountXMAS(input []string) int {
	total := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[0]); j++ {
			// horizontal
			if j+3 < len(input[0]) {
				if input[i][j] == 'X' && input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
					total++
				}
				if input[i][j] == 'S' && input[i][j+1] == 'A' && input[i][j+2] == 'M' && input[i][j+3] == 'X' {
					total++
				}
			}

			// vertical
			if i+3 < len(input) {
				if input[i][j] == 'X' && input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
					total++
				}
				if input[i][j] == 'S' && input[i+1][j] == 'A' && input[i+2][j] == 'M' && input[i+3][j] == 'X' {
					total++
				}
			}

			// diag right
			if i+3 < len(input) && j+3 < len(input[0]) {
				if input[i][j] == 'X' && input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
					total++
				}
				if input[i][j] == 'S' && input[i+1][j+1] == 'A' && input[i+2][j+2] == 'M' && input[i+3][j+3] == 'X' {
					total++
				}
			}

			// diag left
			if i+3 < len(input) && j-3 >= 0 {
				if input[i][j] == 'X' && input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
					total++
				}
				if input[i][j] == 'S' && input[i+1][j-1] == 'A' && input[i+2][j-2] == 'M' && input[i+3][j-3] == 'X' {
					total++
				}
			}
		}
	}
	return total
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
