package main

import (
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day11/input.txt")
	waitingArea := newWaitingArea(input)
	utils.Print("Task 1. After stabilizing there are %d occupied seats", waitingArea.occupiedWhenStableNeighbors())
	utils.Print("Task 2. After stabilizing there are %d occupied seats", waitingArea.occupiedWhenStableVisibility())
}

type waitingArea struct {
	seats [][]*bool // nil represents floor, true occupied, false empty
}

func newWaitingArea(s string) waitingArea {
	lines := utils.SplitLine(s)
	seats := make([][]*bool, len(lines))
	for y, line := range lines {
		seats[y] = make([]*bool, len(line))
		for x, c := range line {
			if c == 'L' {
				value := false
				seats[y][x] = &value
			}
		}
	}

	return waitingArea{seats}
}

func noOccupiedNeighbors(a *waitingArea, x, y int) bool {
	return a.occupiedNeighborsCount(x, y) == 0
}

func (a *waitingArea) occupiedNeighborsCount(x, y int) int {
	count := 0
	for checkY := y - 1; checkY <= y+1; checkY++ {
		for checkX := x - 1; checkX <= x+1; checkX++ {
			if checkX == x && checkY == y {
				continue
			}
			if a.occupied(checkX, checkY) {
				count++
			}
		}
	}
	return count
}

func fourOrMoreNeighbors(a *waitingArea, x, y int) bool {
	return a.occupiedNeighborsCount(x, y) >= 4
}

func (a *waitingArea) visibleOccupiedInDirection(x, y int, xDiff, yDiff int) bool {
	checkY := y + yDiff
	checkX := x + xDiff
	for checkY >= 0 && checkY < len(a.seats) && checkX >= 0 && checkX < len(a.seats[0]) {
		if a.occupied(checkX, checkY) {
			return true
		}
		seat := a.seats[checkY][checkX]
		if seat != nil && !*seat {
			return false
		}
		checkY += yDiff
		checkX += xDiff
	}
	return false
}

func (a *waitingArea) visibleNeighborsCount(x, y int) int {
	count := 0
	for yDiff := -1; yDiff <= 1; yDiff++ {
		for xDiff := -1; xDiff <= 1; xDiff++ {
			if xDiff == 0 && yDiff == 0 {
				continue
			}
			if a.visibleOccupiedInDirection(x, y, xDiff, yDiff) {
				count++
			}
		}
	}
	return count
}

func fiveOrMoreVisible(a *waitingArea, x, y int) bool {
	return a.visibleNeighborsCount(x, y) >= 5
}

func noOccupiedVisible(a *waitingArea, x, y int) bool {
	return a.visibleNeighborsCount(x, y) == 0
}

func (a *waitingArea) countOccupied() int {
	count := 0
	for _, row := range a.seats {
		for _, seat := range row {
			if seat != nil && *seat {
				count++
			}
		}
	}
	return count
}

func (a *waitingArea) occupied(x, y int) bool {
	if x < 0 || x >= len(a.seats[0]) ||
		y < 0 || y >= len(a.seats) {
		return false
	}
	occupied := a.seats[y][x]
	return occupied != nil && *occupied
}

func (a *waitingArea) String() string {
	b := strings.Builder{}
	for _, row := range a.seats {
		for _, seat := range row {
			var r rune
			if seat == nil {
				r = '.'
			} else if *seat {
				r = '#'
			} else {
				r = 'L'
			}
			b.WriteRune(r)
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func (a *waitingArea) occupiedWhenStableNeighbors() int {
	return a.nextUntilStable(fourOrMoreNeighbors, noOccupiedNeighbors).countOccupied()
}

func (a *waitingArea) occupiedWhenStableVisibility() int {
	return a.nextUntilStable(fiveOrMoreVisible, noOccupiedVisible).countOccupied()
}

func (a *waitingArea) nextUntilStable(occupiedToEmpty, emptyToOccupied toggleConditionFunc) *waitingArea {
	n := a          // to not modify the original waitingArea
	changed := true // for the initial run
	for changed {
		changed, n = n.next(occupiedToEmpty, emptyToOccupied)
		utils.Print(n.String())
	}
	return n
}

type toggleConditionFunc func(a *waitingArea, x, y int) bool

func (a *waitingArea) next(occupiedToEmpty, emptyToOccupied toggleConditionFunc) (bool, *waitingArea) {
	change := false
	seats := make([][]*bool, len(a.seats))
	for y, row := range a.seats {
		seats[y] = make([]*bool, len(row))
		for x, seat := range row {
			if seat != nil {
				if *seat {
					if occupiedToEmpty(a, x, y) {
						value := false
						seats[y][x] = &value
						change = true
					} else {
						seats[y][x] = seat
					}
				} else {
					if emptyToOccupied(a, x, y) {
						value := true
						seats[y][x] = &value
						change = true
					} else {
						seats[y][x] = seat
					}
				}
			}
		}
	}
	return change, &waitingArea{seats}
}
