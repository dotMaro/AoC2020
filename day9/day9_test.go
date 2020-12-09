package main

import "testing"

const input = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestFirstInvalidNumber(t *testing.T) {
	m := newMessage(input, 5)
	res := m.firstInvalidNumber()
	if res != 127 {
		t.Errorf("Should return 127 but returned %d", res)
	}
}

func TestFindEncryptionWeakness(t *testing.T) {
	m := newMessage(input, 5)
	res := m.findEncryptionWeakness(127)
	if res != 62 {
		t.Errorf("Should return 62 but returned %d", res)
	}
}
