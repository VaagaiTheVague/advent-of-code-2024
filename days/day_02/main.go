package day_02

import (
	"fmt"
	"strconv"
	"strings"
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
	return strconv.Itoa(SafeReports(input))
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return strconv.Itoa(SafeReportsDamped(input))
}

func SafeReports(input []string) (count int) {
	reports := ParseInput(input)

	for _, level := range reports {
		if IsSafe(level) {
			count++
		}
	}
	return count
}

func IsSafe(level []int) bool {
	if !IsStrictlySloped(level) {
		return false
	} else {
		if !IsGraduallySloped(level) {
			return false
		}
	}
	return true
}

func SafeReportsDamped(input []string) (count int) {
	reports := ParseInput(input)

	for _, level := range reports {
		if IsSafeDamped(level) {
			count++
		}
	}

	return count
}

func IsSafeDamped(input []int) bool {
	if IsSafe(input) {
		return true
	} else {
		for i := range input {
			removedInput := Remove(input, i)
			if IsSafe(removedInput) {
				return true
			}
		}
	}
	return false
}

func Remove(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func IsGraduallySloped(level []int) bool {
	for i := 1; i < len(level); i++ {
		diff := Abs(level[i] - level[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func IsStrictlySloped(level []int) bool {
	isIncreasing := true
	for i := 1; i < len(level); i++ {
		if level[i] <= level[i-1] {
			isIncreasing = false
			break
		}
	}

	isDecreasing := true
	for i := 1; i < len(level); i++ {
		if level[i] >= level[i-1] {
			isDecreasing = false
			break
		}
	}

	return isIncreasing || isDecreasing
}

func ParseInput(input []string) (reports [][]int) {
	for _, line := range input {
		numbers := []int{}
		for _, level := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(level)
			numbers = append(numbers, number)
		}
		reports = append(reports, numbers)
	}
	// fmt.Println(reports)
	return reports
}
