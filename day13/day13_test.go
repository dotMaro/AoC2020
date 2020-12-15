package main

import "testing"

const input = `939
7,13,x,x,59,x,31,19`

func TestEarliestDeparture(t *testing.T) {
	resID, resTime := earliestDeparture(input)
	if resID != 59 {
		t.Errorf("Should return ID 59 but returned %d", resID)
	}
	if resTime != 5 {
		t.Errorf("Should return time 5 but returned %d", resTime)
	}
}

func TestSubsquentMatch(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{"17,x,13,19", 3417},
		{"67,7,59,61", 754018},
		{"67,x,7,59,61", 779210},
		{"67,7,x,59,61", 1261476},
		{"1789,37,47,1889", 1202161486},
	}

	for _, tc := range testCases {
		table := newTimetable("\n" + tc.input)
		res := table.subsequentMatch()
		if res != tc.exp {
			t.Errorf("Should return %d but returned %d", tc.exp, res)
		}
	}
}
