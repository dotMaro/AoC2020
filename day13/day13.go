package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day13/input.txt")

	idClosest, closest := earliestDeparture(input)
	utils.Print("Task 1. The earliest bus you can take is %d after %d minutes (%d)", idClosest, closest, closest*idClosest)
	table := newTimetable(input)
	utils.Print("Task 2. Earliest subsequent match is %d", table.subsequentMatch())
}

func earliestDeparture(s string) (int, int) {
	lines := utils.SplitLine(s)
	departure, _ := strconv.Atoi(lines[0])

	closest := 99999999
	idClosest := -1
	for _, id := range strings.Split(lines[1], ",") {
		if id == "x" {
			continue
		}
		idInt, _ := strconv.Atoi(id)
		afterLast := departure % idInt
		wait := idInt - afterLast
		if wait < closest {
			closest = wait
			idClosest = idInt
		}
	}
	return idClosest, closest
}

func newTimetable(s string) timetable {
	lines := utils.SplitLine(s)
	departure, _ := strconv.Atoi(lines[0])

	splitIDs := strings.Split(lines[1], ",")
	ids := make([]busID, len(splitIDs))
	for i, id := range splitIDs {
		var id2 busID
		if id == "x" {
			id2 = wildcard
		} else {
			idInt, _ := strconv.Atoi(id)
			id2 = busID(idInt)
		}
		ids[i] = id2
	}
	return timetable{departure, ids}
}

type timetable struct {
	departure int
	buses     []busID
}

type busID int

const wildcard busID = -1

func (table timetable) subsequentMatch() int {
	highest := busID(0)
	highestIndex := 0
	for i, id := range table.buses {
		if id > highest {
			highest = id
			highestIndex = i
		}
	}

	i := 1
	for {
		t := int(highest)*i - highestIndex
		ok := true
		for ix, id := range table.buses {
			if id != wildcard && (t+ix)%int(id) != 0 {
				ok = false
				break
			}
		}
		if ok {
			return t
		}
		i++
	}
}
