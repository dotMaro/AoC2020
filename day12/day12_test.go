package main

import (
	"testing"
)

const input = `F10
N3
F7
R90
F11`

func TestTraverse(t *testing.T) {
	res := traverse(input)
	if res != 25 {
		t.Errorf("Should return 25 but returned %d", res)
	}
}

func TestTraverseWaypoint(t *testing.T) {
	res := traverseWaypoint(input)
	if res != 286 {
		t.Errorf("Should return 286 but returned %d", res)
	}
}

func TestRotateTurns(t *testing.T) {
	testCases := []struct {
		turns, east, north int // input
		expEast, expNorth  int // expected
	}{
		{
			1,
			1, 2,
			2, -1,
		},
		{
			3,
			1, 2,
			-2, 1,
		},
		{
			-1,
			1, 2,
			-2, 1,
		},
	}

	for _, tc := range testCases {
		resEast, resNorth := rotateTurns(tc.turns, tc.east, tc.north)
		if resEast != tc.expEast || resNorth != tc.expNorth {
			t.Errorf("Should return (%d, %d) but returned (%d, %d)", tc.expEast, tc.expNorth, resEast, resNorth)
		}
	}
}
