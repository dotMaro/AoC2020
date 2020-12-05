package main

import "testing"

func TestSeatID(t *testing.T) {
	testCases := []struct {
		input  string
		seatID float64
	}{
		{"BFFFBBFRRR", 567},
		{"FFFBBBFRRR", 119},
		{"BBFFBBFRLL", 820},
	}

	for _, tc := range testCases {
		res := seatID(tc.input)
		if res != tc.seatID {
			t.Errorf("Should return %v but returned %v", tc.seatID, res)
		}
	}
}
