package main

import (
	"strconv"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day9/input.txt")

	m := newMessage(input, 25)
	invalid := m.firstInvalidNumber()
	utils.Print("Task 1. First invalid number is %d", invalid)
	utils.Print("Task 2. First invalid number is %d", m.findEncryptionWeakness(invalid))
}

type message struct {
	numbers  []int
	preamble int
}

func newMessage(s string, preamble int) message {
	split := utils.SplitLine(s)
	numbers := make([]int, len(split))

	for i, line := range utils.SplitLine(s) {
		nbr, _ := strconv.Atoi(line)
		numbers[i] = nbr
	}

	return message{
		numbers:  numbers,
		preamble: preamble,
	}
}

func (m *message) firstInvalidNumber() int {
	for i, nbr := range m.numbers {
		if !m.isValid(nbr, i) {
			return nbr
		}
	}
	return -1
}

func (m *message) isValid(value, index int) bool {
	if index < m.preamble {
		// Preamble number
		return true
	}
	for i1 := index - m.preamble; i1 < index; i1++ {
		for i2 := index - m.preamble; i2 < index; i2++ {
			if i1 == i2 {
				continue
			}
			if m.numbers[i1]+m.numbers[i2] == value {
				return true
			}
		}
	}
	return false
}

func (m *message) findEncryptionWeakness(invalid int) int {
	type sum struct {
		total    int
		smallest int
		largest  int
	}
	sums := []*sum{}
	sumStart := 0
	for _, nbr := range m.numbers {
		for _, s := range sums[sumStart:] {
			s.total += nbr
			if s.total > invalid {
				sumStart++
				continue
			}
			if nbr > s.largest {
				s.largest = nbr
			}
			if nbr < s.smallest {
				s.smallest = nbr
			}
			if s.total == invalid {
				return s.smallest + s.largest
			}
		}
		sums = append(sums, &sum{nbr, nbr, nbr})
	}
	return -1
}
