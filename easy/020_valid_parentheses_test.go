package main

import "testing"

var testCases = []struct {
	input  string
	output bool
}{
	{"()", true},
	{"()[]{}", true},
	{"(]", false},
	{"([)]", false},
	{"{[]}", true},
}

func TestIsValidParentheses(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			res := IsValidParenthesesSolution1(tc.input)
			if res != tc.output {
				t.Errorf("for input %q got result %v, want %v", tc.input, res, tc.output)
			}
		})
	}
}
