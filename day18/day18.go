package main

import (
	"bufio"
	"strconv"

	"github.com/dotMaro/AoC2020/utils"
)

func main() {
	file := utils.OpenFile("day18/input.txt")
	scanner := bufio.NewScanner(file)
	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum1 += solveExpression(line)
		sum2 += solveWithPrioritizedAddition(line)
	}
	utils.Print("Task 1. The sum is %d", sum1)
	utils.Print("Task 2. The sum is %d", sum2)
}

func solveExpression(s string) int {
	result := 0
	operator := add
	paranthesesDepth := 0
	for i, r := range s {
		if paranthesesDepth > 0 {
			switch r {
			case '(':
				paranthesesDepth++
			case ')':
				paranthesesDepth--
			}
			continue
		}
		switch r {
		case ' ':
			continue
		case '+':
			operator = add
		case '*':
			operator = multiply
		case '(':
			result = operator(result, solveExpression(s[i+1:]))
			paranthesesDepth++
		case ')':
			return result
		default:
			// Assume it is a number.
			value, _ := strconv.Atoi(string(r))
			result = operator(result, value)
		}
	}
	return result
}

type operator func(a, b int) int

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func solveWithPrioritizedAddition(s string) int {
	curSum := 0
	product := 1
	lastOperatorWasAddition := false
	paranthesesDepth := 0

	handleNewValue := func(value int) {
		if curSum == 0 {
			// Beginning of expression or after an asterix.
			curSum = value
		} else if lastOperatorWasAddition {
			curSum += value
		} else {
			product *= value
		}
	}

	for i, r := range s {
		if paranthesesDepth > 0 {
			// The subexpression has been calculated already, skip it.
			switch r {
			case '(':
				paranthesesDepth++
			case ')':
				paranthesesDepth--
			}
			continue
		}
		switch r {
		case ' ':
			continue
		case '+':
			lastOperatorWasAddition = true
		case '*':
			lastOperatorWasAddition = false
			// No more additions that can affect curSum, multiply it into product.
			if curSum != 0 {
				product *= curSum
			}
			curSum = 0
		case '(':
			paranthesesDepth++
			// Recursively solve subexpression within parantheses.
			value := solveWithPrioritizedAddition(s[i+1:])
			handleNewValue(value)
		case ')':
			// Means that this was a subexpression.
			return curSum * product
		default:
			value, _ := strconv.Atoi(string(r))
			handleNewValue(value)
		}
	}

	return curSum * product
}
