package day_05

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
	orders, page := Parse(input)

	order := map[int][]int{}
	for _, o := range orders {
		order[o[0]] = append(order[o[0]], o[1])
	}

	indices := []int{}
	for i, p := range page {
		if IsValidPage(p, order) {
			indices = append(indices, i)
		}
	}

	sum := 0
	for _, i := range indices {
		sum += MiddleElement(page[i])
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	orders, page := Parse(input)

	order := map[int][]int{}
	for _, o := range orders {
		order[o[0]] = append(order[o[0]], o[1])
	}

	sum := processPages(page, order)

	return strconv.Itoa(sum)
}

func processPages(pages [][]int, dict map[int][]int) int {
	var total int

	// Filter and process pages
	for _, page := range pages {
		if !IsValidPage(page, dict) {
			// Perform insertion sort equivalent
			sortedPage := insertionSort(page, dict)

			// Get middle element
			if len(sortedPage) > 0 {
				middleIndex := len(sortedPage) / 2
				total += sortedPage[middleIndex]
			}
		}
	}

	return total
}

func insertionSort(page []int, dict map[int][]int) []int {
	sorted := []int{}

	for _, e := range page {
		// Find insertion point
		insertIndex := -1
		for i, x := range sorted {
			if containsInDict(dict, e, x) {
				insertIndex = i
				break
			}
		}

		// Insert the element
		if insertIndex == -1 {
			sorted = append(sorted, e)
		} else {
			sorted = append(sorted[:insertIndex], append([]int{e}, sorted[insertIndex:]...)...)
		}
	}

	return sorted
}

func containsInDict(dict map[int][]int, e, x int) bool {
	if dictEntry, exists := dict[e]; exists {
		for _, val := range dictEntry {
			if val == x {
				return true
			}
		}
	}
	return false
}

func validPage(page []int, dict map[int][]int) bool {
	// Implement your validation logic here
	// This is a placeholder implementation
	return len(page) > 0
}

func Parse(input []string) ([][]int, [][]int) {
	split := slices.Index(input, "")
	orderString := input[:split]
	pageString := input[split+1:]

	orders := [][]int{}
	for _, line := range orderString {
		order := strings.Split(line, "|")
		left, _ := strconv.Atoi(order[0])
		right, _ := strconv.Atoi(order[1])
		orders = append(orders, []int{left, right})
	}

	pages := [][]int{}
	for _, str := range pageString {
		p := strings.Split(str, ",")

		numbers := []int{}
		for _, page := range p {
			number, _ := strconv.Atoi(page)
			numbers = append(numbers, number)
		}

		pages = append(pages, numbers)
	}

	return orders, pages
}

func IsValidPage(page []int, order map[int][]int) bool {

	for i, x := range page {
		if i > 0 {
			if deps, exists := order[x]; exists {
				for _, dep := range deps {
					if dep == page[i-1] {
						return false
					}
				}
			}
		}
	}
	return true
}

func MiddleElement(slice []int) int {
	return slice[len(slice)/2]
}
