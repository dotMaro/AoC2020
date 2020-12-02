package main

import "testing"

func TestPositionPolicy(t *testing.T) {
	testCases := []struct {
		input string
		valid bool
	}{
		{
			input: "1-3 a: abcde",
			valid: true,
		},
		{
			input: "1-3 b: cdefg",
			valid: false,
		},
		{
			input: "2-9 c: ccccccccc",
			valid: false,
		},
	}

	for _, tc := range testCases {
		pw := newPasswordPositionPolicy(tc.input)
		if pw.valid() != tc.valid {
			t.Errorf("input %s should have validity %t", tc.input, tc.valid)
		}
	}
}
