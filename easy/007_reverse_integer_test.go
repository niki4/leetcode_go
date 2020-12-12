package main

import (
	"fmt"
	"testing"
)

var reverseIntegerTC = []struct {
	input  int
	output int
}{
	{123, 321},
	{-123, -321},
	{120, 21},
	{0, 0},
	{1534236469, 0},
}

func TestReverseInteger(t *testing.T) {
	for _, tc := range reverseIntegerTC {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			res := reverse(tc.input)
			if res != tc.output {
				t.Errorf("for input %d - want %d, got %d\n", tc.input, tc.output, res)
			}
		})
	}
}
