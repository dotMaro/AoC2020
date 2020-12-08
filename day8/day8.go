package main

import (
	"strconv"
	"strings"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	input := utils.ReadFile("day8/input.txt")

	task1, task2 := make(chan struct{}), make(chan struct{})
	go func(chan struct{}) {
		_, acc := run(input)
		utils.Print("Task 1. Acc on duplicate instruction is %d", acc)
		close(task1)
	}(task1)

	go func(chan struct{}) {
		acc := repair(input)
		utils.Print("Task 2. Acc on reparation is %d", acc)
		close(task2)
	}(task2)
	<-task1
	<-task2
}

type interpreter struct {
	pc  int
	acc int
}

// executeOp and return if it was a success.
func (i *interpreter) executeOp(s string) bool {
	switch s[:3] {
	case "acc":
		value, _ := strconv.Atoi(s[5:])
		switch s[4] {
		case '+':
			i.acc += value
		case '-':
			i.acc -= value
		}
		i.pc++
	case "jmp":
		value, _ := strconv.Atoi(s[5:])
		if value == 0 {
			// 0 jump, failure
			return false
		}
		switch s[4] {
		case '+':
			i.pc += value
		case '-':
			i.pc -= value
		}
	case "nop":
		i.pc++
	}

	return true
}

// run a program and return if it was a successful execution and the acc value.
// Even if the execution is not successful acc will be returned.
func run(s string) (bool, int) {
	i := interpreter{}
	instructionsRun := make(map[int]struct{})
	splitLines := utils.SplitLine(s)
	for {
		if i.pc == len(splitLines) {
			// Successful execution
			return true, i.acc
		}
		if i.pc < 0 || i.pc > len(splitLines) {
			// Illegal jmp, failure
			return false, i.acc
		}
		line := splitLines[i.pc]
		_, hasRun := instructionsRun[i.pc]
		if hasRun {
			// Loop, failure
			return false, i.acc
		}
		// Nothing uses the set during execution so it is safe to set before the actual execution
		instructionsRun[i.pc] = struct{}{}
		success := i.executeOp(line)
		if !success {
			return false, i.acc
		}
	}
}

func repair(s string) int {
	lineCount := strings.Count(s, "\n")
	for ix := 0; ix < lineCount; ix++ {
		// First try replacing the nth jmp with nop
		normal, acc := run(replaceNth(s, "jmp", "nop", ix))
		if normal {
			return acc
		}
		// Then the nth nop with jmp
		normal, acc = run(replaceNth(s, "nop", "jmp", ix))
		if normal {
			return acc
		}
	}
	return -1
}

func replaceNth(s string, old, new string, n int) string {
	occurrence := 0
	i := 0
	for i < len(s) {
		offset := strings.Index(s[i:], old)
		i += offset
		if occurrence == n {
			return s[:i] + strings.Replace(s[i:], old, new, 1)
		}
		i += len(old)
		occurrence++
	}
	return s
}
