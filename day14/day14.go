package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day14/input.txt")

	sum := initProgram(input)
	utils.Print("Task 1. The sum is %d", sum)
	sum2 := initProgramV2(input)
	utils.Print("Task 2. The sum is %d", sum2)
}

func initProgram(s string) uint64 {
	var mask, overwrite uint64
	mem := make(map[uint64]uint64)
	for _, line := range utils.SplitLine(s) {
		split := strings.Split(line, " ")
		value := split[2]
		firstWord := split[0]
		if firstWord == "mask" {
			// Write overwrite value, e.g. 1X0 -> 100
			ones := strings.ReplaceAll(value, "X", "0")
			overwrite, _ = strconv.ParseUint(ones, 2, 36)
			// Write mask for the overwrite, e.g. 1X0 -> 101
			maskString := strings.ReplaceAll(value, "0", "1")
			maskString = strings.ReplaceAll(maskString, "X", "0")
			mask, _ = strconv.ParseUint(maskString, 2, 36)
		} else if firstWord[:3] == "mem" {
			address, _ := strconv.ParseUint(firstWord[4:len(firstWord)-1], 10, 36)
			unfiltered, _ := strconv.ParseUint(split[2], 10, 36)
			onesAdded := mask | unfiltered
			test := mask ^ overwrite
			mem[address] = uint64(onesAdded & ^test)
		}
	}

	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}

type mask []maskBit

func newMask(s string) mask {
	m := make([]maskBit, len(s))
	for i, r := range s {
		switch r {
		case '0':
			m[i] = Zero
		case '1':
			m[i] = One
		case 'X':
			m[i] = Floating
		default:
			panic(fmt.Sprintf("incorrect mask bit %v", r))
		}
	}
	return m
}

func (m mask) apply(mem map[uint64]uint64, address int64, value uint64) {
	m.applyRecursive(mem, address, value, 0)
}

func (m mask) applyRecursive(mem map[uint64]uint64, address int64, value uint64, fromIndex int) {
	hadFloating := false
	for i := fromIndex; i < len(m); i++ {
		b := m[i]
		switch b {
		case One:
			// Set to one
			address |= 0b1 << int64(len(m)-1-i)
		case Floating:
			hadFloating = true
			m.applyRecursive(mem, address, value, i+1)
			// Flip bit
			address ^= 0b1 << int64(len(m)-1-i)
			m.applyRecursive(mem, address, value, i+1)
		}
	}
	if !hadFloating {
		mem[uint64(address)] = value
	}
}

type maskBit int

const (
	Zero maskBit = iota
	One
	Floating
)

func initProgramV2(s string) uint64 {
	var mask mask
	mem := make(map[uint64]uint64)
	for _, line := range utils.SplitLine(s) {
		split := strings.Split(line, " ")
		value := split[2]
		firstWord := split[0]
		if firstWord == "mask" {
			mask = newMask(value)
		} else if firstWord[:3] == "mem" {
			address, _ := strconv.ParseInt(firstWord[4:len(firstWord)-1], 10, 36)
			valueInt, _ := strconv.ParseUint(value, 10, 64)
			mask.apply(mem, address, valueInt)
		}
	}

	var sum uint64
	for _, v := range mem {
		sum += v
	}
	return sum
}
