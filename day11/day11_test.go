package main

import "testing"

const input = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestOccupiedWhenStableNeighbors(t *testing.T) {
	a := newWaitingArea(input)
	res := a.occupiedWhenStableNeighbors()
	if res != 37 {
		t.Errorf("Should return 37 but returned %d", res)
	}
}

func TestOccupiedWhenStableVisibility(t *testing.T) {
	a := newWaitingArea(input)
	res := a.occupiedWhenStableVisibility()
	if res != 26 {
		t.Errorf("Should return 26 but returned %d", res)
	}
}
