package main

import (
	"strconv"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day12/input.txt")
	utils.Print("Step 1. Manhattan distance is %d", traverse(input))
	utils.Print("Step 2. Manhattan distance is %d", traverseWaypoint(input))
}

type direction int

const (
	eastDirection direction = iota
	southDirection
	westDirection
	northDirection
)

func (d direction) shift(s int) direction {
	dir := direction(s)
	d = (d + dir) % 4
	if d < 0 {
		d += 4
	}
	return d
}

func (d direction) forward(steps int) (east int, north int) {
	switch d {
	case eastDirection:
		east += steps
	case westDirection:
		east -= steps
	case northDirection:
		north += steps
	case southDirection:
		north -= steps
	}
	return
}

func traverse(s string) int {
	var east, north int
	direction := eastDirection
	for _, line := range utils.SplitLine(s) {
		value, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'N':
			north += value
		case 'S':
			north -= value
		case 'E':
			east += value
		case 'W':
			east -= value
		case 'L':
			direction = direction.shift(-value / 90)
		case 'R':
			direction = direction.shift(value / 90)
		case 'F':
			e, n := direction.forward(value)
			east += e
			north += n
		default:
			panic(line[0])
		}
		// utils.Print("East %d North %d", east, north)
	}

	if east < 0 {
		east *= -1
	}
	if north < 0 {
		north *= -1
	}
	return east + north
}

func traverseWaypoint(s string) int {
	waypointEast, waypointNorth := 10, 1
	var shipEast, shipNorth int
	for _, line := range utils.SplitLine(s) {
		value, _ := strconv.Atoi(line[1:])
		switch line[0] {
		case 'N':
			waypointNorth += value
		case 'S':
			waypointNorth -= value
		case 'E':
			waypointEast += value
		case 'W':
			waypointEast -= value
		case 'L':
			waypointEast, waypointNorth = rotateTurns(-value/90, waypointEast, waypointNorth)
		case 'R':
			waypointEast, waypointNorth = rotateTurns(value/90, waypointEast, waypointNorth)
		case 'F':
			for i := 0; i < value; i++ {
				shipEast += waypointEast
				shipNorth += waypointNorth
			}
		default:
			panic(line[0])
		}
		// utils.Print("%s East %d North %d  -  WayEast %d WayNorth %d",
		// 	line, shipEast, shipNorth, waypointEast, waypointNorth)
	}

	if shipEast < 0 {
		shipEast *= -1
	}
	if shipNorth < 0 {
		shipNorth *= -1
	}
	return shipEast + shipNorth
}

func rotateTurns(turns, east, north int) (int, int) {
	for i := 0; i < turns; i++ {
		east, north = rotate(east, north)
	}
	for i := 0; i > turns; i-- {
		north, east = rotate(north, east)
	}
	return east, north
}

func rotate(east, north int) (int, int) {
	return north, -east
}
