package main

import "testing"

const input = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestRun(t *testing.T) {
	_, res := run(input)
	if res != 5 {
		t.Errorf("Should return 5, but returned %d", res)
	}
}

func TestRepair(t *testing.T) {
	res := repair(input)
	if res != 8 {
		t.Errorf("Should return 8, but returned %d", res)
	}
}

func TestReplaceNth(t *testing.T) {
	const inputReplaced = `nop +0
acc +1
nop +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

	testCases := []struct {
		s        string
		old, new string
		n        int
		expect   string
	}{
		{"aaaa", "a", "b", 0, "baaa"},
		{"aaaa", "a", "b", 1, "abaa"},
		{"aaaa", "a", "b", 2, "aaba"},
		{input, "jmp", "nop", 0, inputReplaced},
	}

	for _, tc := range testCases {
		res := replaceNth(tc.s, tc.old, tc.new, tc.n)
		if res != tc.expect {
			t.Errorf("Should return %s but returned %s", tc.expect, res)
		}
	}
}
