package main

import "github.com/dotMaro/AoC2020/utils"

func main() {
	input := utils.ReadFile("day6/input.txt")

	sum := answerSum(input)
	utils.Print("Task 1. Sum is %d", sum)
	collectiveSum := collectiveAnswerSum(input)
	utils.Print("Task 2. Collective sum is %d", collectiveSum)
}

// answerSum returns the sum of yes answers in every group.
func answerSum(s string) int {
	answers := make(map[rune]struct{})
	sum := 0
	for _, line := range utils.SplitLine(s) {
		if len(line) == 0 {
			sum += len(answers)
			answers = make(map[rune]struct{})
			continue
		}

		for _, c := range line {
			answers[c] = struct{}{}
		}
	}
	return sum + len(answers)
}

// collectiveAnswerSum returns the sum of yes answers that every
// individual in a group has answered.
func collectiveAnswerSum(s string) int {
	collectiveAnswers := make(map[rune]struct{})
	sum := 0
	firstInGroup := true
	for _, line := range utils.SplitLine(s) {
		// An empty line indicates we're going to a new group
		if len(line) == 0 {
			sum += len(collectiveAnswers)
			collectiveAnswers = make(map[rune]struct{})
			firstInGroup = true
			continue
		}

		// Go through the individual's answers
		individualAnswers := make(map[rune]struct{})
		for _, c := range line {
			if firstInGroup {
				collectiveAnswers[c] = struct{}{}
			}
			individualAnswers[c] = struct{}{}
		}

		// Check if the individual did not answer yes
		// to any of the collective answers, if so delete them
		for a := range collectiveAnswers {
			if _, ok := individualAnswers[a]; !ok {
				delete(collectiveAnswers, a)
			}
		}
		firstInGroup = false
	}
	return sum + len(collectiveAnswers)
}
