package easy

import (
	"fmt"
	"testing"
)

var palindromeNumTC = []struct {
	input  int
	output bool
}{
	{121, true},
	{-121, false},
	{10, false},
	{-101, false},
}

func TestIsPalindrome(t *testing.T) {
	for _, tc := range palindromeNumTC {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			if res := isPalindrome(tc.input); res != tc.output {
				t.Errorf("want %t, got %t", tc.output, res)
			}
		})
	}
}
