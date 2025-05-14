package main

import (
	"reflect"
	"testing"
)

var FindMissingRangesTestCases = []struct {
	nums         []int
	lower, upper int
	expected     []string
}{
	{[]int{0, 1, 3, 50, 75}, 0, 99, []string{"2", "4->49", "51->74", "76->99"}},
	{[]int{}, 1, 1, []string{"1"}},
	{[]int{}, -3, -1, []string{"-3->-1"}},
	{[]int{-1}, -1, -1, []string{}},
	{[]int{-1}, -2, -1, []string{"-2"}},
}

func TestFindMissingRanges(t *testing.T) {
	for _, tc := range FindMissingRangesTestCases {
		if res := findMissingRanges(tc.nums, tc.lower, tc.upper); !reflect.DeepEqual(res, tc.expected) {
			t.Errorf("for input nums=%v lower=%d upper=%d\nexpected\t%#v\n got\t\t%#v",
				tc.nums, tc.lower, tc.upper, tc.expected, res)
		}
	}
}
