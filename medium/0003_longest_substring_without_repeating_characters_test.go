package medium

import "testing"

var lengthOfLongestSubstringSolutions = []struct {
	name string
	fn   func(s string) int
}{
	{"lengthOfLongestSubstring", lengthOfLongestSubstring},
	{"lengthOfLongestSubstring2", lengthOfLongestSubstring2},
	{"lengthOfLongestSubstring3", lengthOfLongestSubstring3},
}

var lengthOfLongestSubstringTestCases = []struct {
	input  string
	output int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
	{"", 0},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for _, solution := range lengthOfLongestSubstringSolutions {
		for _, tc := range lengthOfLongestSubstringTestCases {
			t.Run(solution.name, func(t *testing.T) {
				if res := solution.fn(tc.input); res != tc.output {
					t.Errorf("want %d, got %d", tc.output, res)
				}
			})
		}
	}
}
