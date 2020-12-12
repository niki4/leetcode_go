package main

import (
	"math"
	"strconv"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within
the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem,
assume that your function returns 0 when the reversed integer overflows.
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2.2 MB, less than 100.00% of Go
func reverse(x int) int {
	xStr := strconv.Itoa(abs(x))
	runes := []rune(xStr)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	res, _ := strconv.Atoi(string(runes))

	// math.MinInt32 == -(1 << 31)
	// math.MaxInt32 == (1 << 31) - 1
	if res < math.MinInt32 || res > math.MaxInt32 { // overflow
		return 0
	}

	if x < 0 {
		return -res
	}
	return res
}
