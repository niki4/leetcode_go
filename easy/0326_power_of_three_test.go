package easy

import "testing"

var isPowerOfThreeTestCases = []struct {
	num     int
	isPower bool
}{
	{27, true},
	{0, false},
	{9, true},
	{45, false},
	{1, true},
}

var isPowerOfThreeSolutions = []struct {
	fnName string
	fnObj  func(n int) bool
}{
	{"isPowerOfThree", isPowerOfThree},
	{"isPowerOfThreeV2", isPowerOfThreeV2},
}

func TestIsPowerOfThree(t *testing.T) {
	for _, sol := range isPowerOfThreeSolutions {
		for _, tc := range isPowerOfThreeTestCases {
			t.Run(
				sol.fnName,
				func(t *testing.T) {
					if res := sol.fnObj(tc.num); res != tc.isPower {
						t.Errorf("for num %d expected %t, got %t", tc.num, tc.isPower, res)
					}
				})
		}
	}
}
