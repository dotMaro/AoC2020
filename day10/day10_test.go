package main

import "testing"

const input1 = `16
10
15
5
1
11
7
19
6
12
4`

const input2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

func TestDistribution(t *testing.T) {
	a := newAdapters(input1)
	res := a.distribution()
	if res != 22 {
		t.Errorf("Should return 22 but returned %d", res)
	}
}

func TestArrangements(t *testing.T) {
	a := newAdapters(input1)
	res := a.arrangements()
	if res != 8 {
		t.Errorf("Should return 8 but returned %d", res)
	}
	a = newAdapters(input2)
	res = a.arrangements()
	if res != 19208 {
		t.Errorf("Should return 19208 but returned %d", res)
	}
}
