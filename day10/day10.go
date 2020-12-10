package main

import (
	"sort"
	"strconv"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day10/input.txt")

	adapters := newAdapters(input)
	utils.Print("Task 1. Product is %d", adapters.distribution())
	utils.Print("Task 2. There are %d arrangements", adapters.arrangements())
}

type adapters []int

func newAdapters(s string) adapters {
	lines := utils.SplitLine(s)
	adapters := make([]int, len(lines)+1)
	adapters[0] = 0
	for i, line := range lines {
		nbr, _ := strconv.Atoi(line)
		adapters[i+1] = nbr
	}

	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	return adapters
}

func (a adapters) distribution() int {
	oneJoltDiff := 0
	threeJoltDiff := 0
	prev := 0
	for _, jolt := range a {
		switch jolt - prev {
		case 1:
			oneJoltDiff++
		case 3:
			threeJoltDiff++
		}
		prev = jolt
	}
	return oneJoltDiff * threeJoltDiff
}

func (a adapters) arrangements() int {
	return a.arrangementsRecursive(0, make(map[int]int))
}

// arrangementsRecursive checks the amount of arrangements from the startIndex towards the end of the adapter slice.
// It recursively looks at all the available paths (rightwards) and sums up the amount of arrangements.
func (a adapters) arrangementsRecursive(startIndex int, memory map[int]int) int {
	count := 0

	// Check the cache
	if m, ok := memory[startIndex]; ok {
		return m
	}

	if startIndex == len(a)-2 {
		// Second to last adapter, there's only one more path available
		return 1
	}

	// Look at the next three adapters
	for i := startIndex + 1; i <= startIndex+3 && i < len(a); i++ {
		// If they got an okay jolt difference, check the arrangement count along that path
		if a[i]-a[startIndex] <= 3 {
			count += a.arrangementsRecursive(i, memory)
		}
	}
	// Cache the result, all lookups for the same start index will be the same
	memory[startIndex] = count
	return count
}
