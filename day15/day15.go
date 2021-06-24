package main

import "github.com/dotMaro/AoC2020/utils"

func main() {
	res1 := memoryGame(2020, []int{6, 13, 1, 15, 2, 0})
	utils.Print("Task 1. The 2020th number is %d", res1)
	res2 := memoryGame(30000000, []int{6, 13, 1, 15, 2, 0})
	utils.Print("Task 2. The 30000000th number is %d", res2)
}

func memoryGame(stopAfterIteration int, startingNumbers []int) int {
	lastEncounters := make(map[int]int)
	// Add all but the last starting numbers
	for i := 0; i < len(startingNumbers)-1; i++ {
		v := startingNumbers[i]
		lastEncounters[v] = i
	}
	lastNumber := startingNumbers[len(startingNumbers)-1]
	for i := len(startingNumbers); i < stopAfterIteration; i++ {
		lastEncounter, hasEncountered := lastEncounters[lastNumber]
		// Set last encounter after checking if it has been encountered before
		lastEncounters[lastNumber] = i - 1
		if hasEncountered {
			lastNumber = i - 1 - lastEncounter // last index - last encounter
		} else {
			lastNumber = 0
		}
	}

	return lastNumber
}
