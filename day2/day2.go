package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day2/input.txt")
	occurrenceValidCount := 0
	positionValidCount := 0
	for _, line := range strings.Split(input, "\r\n") {
		pw := newPasswordOccurrencePolicy(line)
		if pw.valid() {
			occurrenceValidCount++
		}
		pw = newPasswordPositionPolicy(line)
		if pw.valid() {
			positionValidCount++
		}
	}
	utils.Print("Task 1. There are %d valid passwords", occurrenceValidCount)
	utils.Print("Task 2. There are %d valid passwords", positionValidCount)
}

func newPasswordOccurrencePolicy(input string) password {
	splitInput := strings.Split(input, " ")
	if len(splitInput) != 3 {
		panic(fmt.Sprintf("incorrect input: %v", input))
	}
	splitLimit := strings.Split(splitInput[0], "-")
	lower, err := strconv.Atoi(splitLimit[0])
	if err != nil {
		panic(err)
	}
	upper, err := strconv.Atoi(splitLimit[1])
	if err != nil {
		panic(err)
	}
	reader := strings.NewReader(splitInput[1])
	r, _, err := reader.ReadRune()
	if err != nil {
		panic(err)
	}

	return password{
		s: splitInput[2],
		policy: occurrencePolicy{
			char:     r,
			minLimit: lower,
			maxLimit: upper,
		},
	}
}

func newPasswordPositionPolicy(input string) password {
	splitInput := strings.Split(input, " ")
	if len(splitInput) != 3 {
		panic(fmt.Sprintf("incorrect input: %v", input))
	}
	splitLimit := strings.Split(splitInput[0], "-")
	pos1, err := strconv.Atoi(splitLimit[0])
	if err != nil {
		panic(err)
	}
	pos2, err := strconv.Atoi(splitLimit[1])
	if err != nil {
		panic(err)
	}

	return password{
		s: splitInput[2],
		policy: positionPolicy{
			char: splitInput[1][0],
			pos1: pos1 - 1, // convert 1-based index to 0-based one
			pos2: pos2 - 1,
		},
	}
}

type password struct {
	s string
	policy
}

func (p password) valid() bool {
	return p.policy.valid(p.s)
}

type policy interface {
	valid(pw string) bool
}

type occurrencePolicy struct {
	char     rune
	minLimit int
	maxLimit int
}

func (p occurrencePolicy) valid(pw string) bool {
	count := 0
	for _, c := range pw {
		if c == p.char {
			count++
			if count > p.maxLimit {
				return false
			}
		}
	}
	return count >= p.minLimit
}

type positionPolicy struct {
	char byte
	pos1 int
	pos2 int
}

func (p positionPolicy) valid(pw string) bool {
	if len(pw) <= p.pos2 {
		return false
	}
	return (pw[p.pos1] == p.char || pw[p.pos2] == p.char) && pw[p.pos1] != pw[p.pos2]
}
