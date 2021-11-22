package main

import "testing"

const input = `.#.
..#
###`

func TestConway3DStep(t *testing.T) {
	c := parseConway3D(input)
	c.steps(6)
	active := c.totalActive()
	if active != 112 {
		t.Errorf("Should have 112 active cubes after 6 steps but had %d", active)
	}
}

func TestConway3DNeighborCount(t *testing.T) {
	testCases := []struct {
		x, y int
		exp  int
	}{
		{x: 0, y: 0, exp: 1},
		{x: 1, y: 1, exp: 5},
		{x: -1, y: 3, exp: 1},
	}

	c := parseConway3D(input)
	for _, tc := range testCases {
		count := c.neighborCount(0, tc.y, tc.x)
		if count != tc.exp {
			t.Errorf("Should return %d for input x %d y %d, but returned %d", tc.exp, tc.x, tc.y, count)
		}
	}
}

func TestConway4DSteps(t *testing.T) {
	c := parseConway4D(input)
	c.steps(6)
	active := c.totalActive()
	if active != 848 {
		t.Errorf("Should have 848 active cubes after 6 steps but had %d", active)
	}
}
