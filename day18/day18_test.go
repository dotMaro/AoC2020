package main

import "testing"

func Test_solveExpression(t *testing.T) {
	testCases := []struct {
		input string
		exp   int
	}{
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, tc := range testCases {
		res := solveExpression(tc.input)
		if res != tc.exp {
			t.Errorf("Should return %d for expression %q, but returned %d", tc.exp, tc.input, res)
		}
	}
}
