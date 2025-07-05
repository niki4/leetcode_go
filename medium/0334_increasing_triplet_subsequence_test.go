package medium

import "testing"

var increasingTripletTestCases = []struct {
	input  []int
	output bool
}{
	{[]int{1, 2, 3, 4, 5}, true},
	{[]int{5, 4, 3, 2, 1}, false},
	{[]int{2, 1, 5, 0, 4, 6}, true},
}

func TestIncreasingTriplet(t *testing.T) {
	for _, tc := range increasingTripletTestCases {
		if res := increasingTriplet(tc.input); res != tc.output {
			t.Errorf("for input %v want %t, got %t", tc.input, tc.output, res)
		}
	}
}
