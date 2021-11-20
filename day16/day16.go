package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

// This implementation is incredibly ugly. I started on this last year in 2020 and just got back to it in now in
// Nov 2021. I finished the last stuff and got it to work but I'm too lazy to refactor. Don't judge me!

func main() {
	input := utils.ReadFile("day16/input.txt")
	errorRate := parseNotes()
	utils.Print("Task 1. Error rate is %d", errorRate)
	product := parseNotes2(input)
	utils.Print("Task 2. Error rate is %d", product)
}

func parseNotes() int {
	var fields []fieldRange

	input := utils.ReadFile("day16/input.txt")
	var (
		nearbyTickets   bool
		errorRate       int
		expectOurTicket bool
		ourTicket       []int
	)
	lines := utils.SplitLine(input)
	ourTicket = make([]int, len(lines))
	for _, line := range lines {
		if expectOurTicket {
			values := strings.Split(line, ",")
			for i, value := range values {
				valueInt, _ := strconv.Atoi(value)
				ourTicket[i] = valueInt
			}
			expectOurTicket = false
		} else if nearbyTickets {
			values := strings.Split(line, ",")
			for _, value := range values {
				valueInt, _ := strconv.Atoi(value)
				inRange := false
				for _, f := range fields {
					if f.isInRange(valueInt) {
						inRange = true
						break
					}
				}
				if !inRange {
					errorRate += valueInt
				}
			}
		} else {
			splitLine := strings.Split(line, ": ")
			if len(splitLine) == 2 {
				for _, ranges := range strings.Split(splitLine[1], " or ") {
					values := strings.Split(ranges, "-")
					lower, _ := strconv.Atoi(values[0])
					upper, _ := strconv.Atoi(values[1])
					r := fieldRange{
						lower: lower,
						upper: upper,
					}
					fields = append(fields, r)
				}
			} else if line == "your ticket:" {
				expectOurTicket = true
			} else if line == "nearby tickets:" {
				nearbyTickets = true
			}
		}
	}

	return errorRate
}

func parseNotes2(input string) int {
	var fields []field
	lines := utils.SplitLine(input)

	var (
		nearbyTickets   bool
		expectOurTicket bool
		ourTicket       []int = make([]int, len(lines))
		possibleColumns [][]field
	)
lineLoop:
	for _, line := range lines {
		if expectOurTicket {
			expectOurTicket = false
			values := strings.Split(line, ",")
			possibleColumns = make([][]field, len(values)) // col -> possible fields
			for i, value := range values {
				valueInt, _ := strconv.Atoi(value)
				ourTicket[i] = valueInt
			}
		} else if nearbyTickets {
			values := strings.Split(line, ",")
			for col, value := range values {
				possibleFields := make(map[string]field) // name -> field  - used for?
				valueInt, _ := strconv.Atoi(value)
				// Go through each field and check if they're compatible with the current value.
				hasAnyCompatibleField := false
				for _, f := range fields {
					if f.isInRange(valueInt) {
						possibleFields[f.name] = f
						hasAnyCompatibleField = true
					}
				}
				// Skip tickets that have values that are not compatible with any field.
				if !hasAnyCompatibleField {
					continue lineLoop
				}
				possibleCol := possibleColumns[col]
				var newPossibleCol []field
				// Filter out fields that are not possible in the current ticket.
				if len(possibleCol) != 0 {
					for _, p := range possibleCol {
						_, ok := possibleFields[p.name]
						if ok {
							newPossibleCol = append(newPossibleCol, p)
						}
					}
				} else {
					for _, f := range possibleFields {
						newPossibleCol = append(newPossibleCol, f)
					}
				}
				possibleColumns[col] = newPossibleCol
			}
		} else {
			switch line {
			case "your ticket:":
				expectOurTicket = true
			case "nearby tickets:":
				nearbyTickets = true
			case "":
				continue
			default:
				fields = append(fields, parseField(line))
			}
		}
	}

	foundFields := make(map[string]int) // name -> col
	for len(foundFields) < len(fields) {
		for c, p := range possibleColumns {
			possibleFieldsCount := 0
			notFoundFieldIndex := 0
			for i, f := range p {
				_, alreadyFound := foundFields[f.name]
				if !alreadyFound {
					possibleFieldsCount++
					notFoundFieldIndex = i
				}
			}
			if possibleFieldsCount == 1 {
				foundFields[p[notFoundFieldIndex].name] = c
			}
		}
	}

	product := 1
	for fieldName, col := range foundFields {
		if strings.HasPrefix(fieldName, "departure") {
			product *= ourTicket[col]
		}
	}

	return product
}

func parseField(s string) field {
	var f field
	splitLine := strings.Split(s, ": ")
	if len(splitLine) == 2 {
		var fieldRanges [2]fieldRange
		for i, ranges := range strings.Split(splitLine[1], " or ") {
			fieldRanges[i] = parseRange(ranges)
		}
		f = field{
			name:   splitLine[0],
			ranges: fieldRanges,
		}
	}
	return f
}

func parseRange(s string) fieldRange {
	values := strings.Split(s, "-")
	lower, _ := strconv.Atoi(values[0])
	upper, _ := strconv.Atoi(values[1])
	return fieldRange{
		lower: lower,
		upper: upper,
	}
}

type field struct {
	name   string
	ranges [2]fieldRange
}

func (f field) isInRange(i int) bool {
	return f.ranges[0].isInRange(i) || f.ranges[1].isInRange(i)
}

type fieldRange struct {
	lower, upper int
}

func (r fieldRange) isInRange(i int) bool {
	return i >= r.lower && i <= r.upper
}
