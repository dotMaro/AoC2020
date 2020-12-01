package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day1/input.txt")
	ints := convertInputToInts(input)

	term1, term2 := findTwo2020Terms(ints)
	utils.Print("Task 1. Product is %d", term1*term2)

	term1, term2, term3 := findThree2020Terms(ints)
	utils.Print("Task 2. Product is %d", term1*term2*term3)
}

// findTwo2020Terms finds two terms that add up to 2020.
func findTwo2020Terms(entries []int) (int, int) {
	for _, e1 := range entries {
		for _, e2 := range entries {
			if e1+e2 == 2020 {
				return e1, e2
			}
		}
	}

	return 0, 0
}

// findThree2020Terms finds three terms that add up to 2020.
func findThree2020Terms(entries []int) (int, int, int) {
	for _, e1 := range entries {
		for _, e2 := range entries {
			for _, e3 := range entries {
				if e1+e2+e3 == 2020 {
					return e1, e2, e3
				}
			}
		}
	}

	return 0, 0, 0
}

func convertInputToInts(input string) []int {
	var ints []int
	for _, line := range strings.Split(input, "\r\n") {
		i, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	return ints
}
