package main

import "testing"

const input1 = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

const minimal = `a bags contain 2 b bags.
b bags contain 2 c bags.`

func TestContainedByCount(t *testing.T) {
	rules := parseRules(input1)
	t.Logf("%+v", rules.bags)
	bag := rules.bagWithColor("shiny gold")
	count := bag.containedByCount()
	if count != 4 {
		t.Errorf("Should return 4, but returned %d", count)
	}
}

func TestMustContain(t *testing.T) {
	testCases := []struct {
		input    string
		color    string
		contains int
	}{
		{input1, "shiny gold", 32},
		{minimal, "a", 6}, // 2*2 + 2
	}

	for _, tc := range testCases {
		rules := parseRules(tc.input)
		bag := rules.bagWithColor(tc.color)
		count := bag.mustContain()
		if count != tc.contains {
			t.Errorf("Should return %d, but returned %d", tc.contains, count)
		}
	}
}
