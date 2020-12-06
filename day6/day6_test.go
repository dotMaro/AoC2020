package main

import "testing"

const input = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestAnswerSum(t *testing.T) {
	res := answerSum(input)
	if res != 11 {
		t.Errorf("Should return 11, but returned %d", res)
	}
}

func TestCollectiveAnswerSum(t *testing.T) {
	res := collectiveAnswerSum(input)
	if res != 6 {
		t.Errorf("Should return 6, but returned %d", res)
	}
}
