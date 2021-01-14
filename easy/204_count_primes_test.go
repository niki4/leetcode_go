package main

import "testing"

var countPrimesTestCases = []struct {
	num   int
	count int
}{
	{10, 4},
	{0, 0},
	{1, 0},
	{3, 1},
	{10000, 1229},
}

func TestCountPrimes(t *testing.T) {
	for _, tc := range countPrimesTestCases {
		if res := countPrimes(tc.num); res != tc.count {
			t.Errorf("for input num %d expected %d, got %d", tc.num, tc.count, res)
		}
	}
}
