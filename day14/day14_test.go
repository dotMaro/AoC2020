package main

import "testing"

const input = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

const input2 = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX1X01
mem[8] = 3
mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX0101
mem[1234] = 24`

const input3 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestInitProgram(t *testing.T) {
	res := initProgram(input)
	if res != 165 {
		t.Errorf("Should return 165 but returned %d", res)
	}
	res = initProgram(input2)
	var exp uint64 = 0b1001 + 0b10101
	if res != exp {
		t.Errorf("Should return %d but returned %d", exp, res)
	}
}

func TestInitProgramV2(t *testing.T) {
	res := initProgramV2(input3)
	if res != 208 {
		t.Errorf("Should return 208 but returned %d", res)
	}
}
