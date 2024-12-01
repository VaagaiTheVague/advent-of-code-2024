package day_01

import (
	"fmt"
	"slices"
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
	distance := sum(CalculateDistances(input))
	return strconv.Itoa(distance)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}

func CalculateDistances(input []string) (distances []int) {
	left := []int{}
	right := []int{}

	for _, line := range input {
		splitStrings := strings.Split(line, " ")
		ln, _ := strconv.Atoi(splitStrings[0])
		left = append(left, ln)
		rn, _ := strconv.Atoi(splitStrings[len(splitStrings)-1])
		right = append(right, rn)
	}

	slices.Sort(left)
	slices.Sort(right)

	for i := range left {
		distances = append(distances, Abs(left[i]-right[i]))
	}

	return distances
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func Abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
