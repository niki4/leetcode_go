package easy

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

var countPrimesSolutions = []struct {
	fnName string
	fnObj  func(n int) (pCounter int)
}{
	{"countPrimesSieveOfEratosthenes", countPrimes},
	{"countPrimesMathBig", countPrimesMathBig},
}

func TestCountPrimes(t *testing.T) {
	for _, sol := range countPrimesSolutions {
		for _, tc := range countPrimesTestCases {
			t.Run(
				sol.fnName,
				func(t *testing.T) {
					if res := sol.fnObj(tc.num); res != tc.count {
						t.Errorf("for input num %d expected %d, got %d", tc.num, tc.count, res)
					}
				})
		}
	}
}
