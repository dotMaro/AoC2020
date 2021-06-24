package main

import "testing"

func TestMemoryGame(t *testing.T) {
	testCases := []struct {
		startingNumbers []int
		toIter          int
		expect          int
	}{
		{[]int{1, 3, 2}, 2020, 1},
		{[]int{2, 1, 3}, 2020, 10},
		{[]int{1, 2, 3}, 2020, 27},
		// {[]int{0, 3, 6}, 30000000, 175594},
	}

	for _, tc := range testCases {
		res := memoryGame(tc.toIter, tc.startingNumbers)
		if res != tc.expect {
			t.Errorf("Should return %d for startingNumbers %v but returned %d", tc.expect, tc.startingNumbers, res)
		}
	}
}
