package main

import "testing"

var firstUniqCharTestCases = []struct {
	input string
	index int
}{
	{"leetcode", 0},
	{"loveleetcode", 2},
	{"", -1},
	{"foofoo", -1},
}

func TestFirstUniqChar(t *testing.T) {
	for _, tc := range firstUniqCharTestCases {
		if res := firstUniqChar(tc.input); res != tc.index {
			t.Errorf("for input %q expected %d, got %d", tc.input, tc.index, res)
		}
	}
}
