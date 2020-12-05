package main

import (
	"math"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day5/input.txt")

	highestSeatID, seatIDSet := findHighestSeatID(input)
	utils.Print("Task 1. Highest seat ID is %v", highestSeatID)
	vacantSeat := seatIDSet.findVacantSeat()
	utils.Print("Task 2. Vacant seat ID is %v", vacantSeat)
}

// findHighestSeatID parses the input and returns the highest seat ID and the seatIDSet.
func findHighestSeatID(input string) (float64, seatIDSet) {
	highestSeatID := 0.0
	seatIDSet := newSeatIDSet()
	for _, line := range utils.SplitLine(input) {
		seatID := seatID(line)
		seatIDSet.add(seatID)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	return highestSeatID, seatIDSet
}

// seatID takes an input and returns the seat ID.
func seatID(s string) float64 {
	// Get row
	var l, u float64 = 0, 127
	for _, c := range s[:7] {
		if c == 'F' {
			u = math.Floor((l + u) / 2)
		} else {
			l = math.Ceil((l + u) / 2)
		}
	}
	row := l
	// Get column
	l, u = 0, 7
	for _, c := range s[7:] {
		if c == 'L' {
			u = math.Floor((l + u) / 2)
		} else {
			l = math.Ceil((l + u) / 2)
		}
	}
	column := l
	return row*8 + column
}

type seatIDSet map[float64]struct{}

func newSeatIDSet() seatIDSet {
	return make(map[float64]struct{})
}

func (s seatIDSet) add(id float64) {
	s[id] = struct{}{}
}

// findVacantSeat finds the vacant seat ID.
func (s seatIDSet) findVacantSeat() float64 {
	// Go through all added seat IDs and check if they have
	// a neighbor that is vacant (not added) and itself have both
	// neighbors present.
	vacantSeat := 0.0
	for seatID := range s {
		leftID := seatID - 1
		_, hasLeft := s[leftID]
		if !hasLeft {
			_, leftHasLeft := s[leftID-1]
			if leftHasLeft {
				vacantSeat = leftID
				break
			}
		}
		rightID := seatID + 1
		_, hasRight := s[rightID]
		if !hasRight {
			_, rightHasRight := s[rightID+1]
			if rightHasRight {
				vacantSeat = rightID
				break
			}
		}
	}
	return vacantSeat
}
