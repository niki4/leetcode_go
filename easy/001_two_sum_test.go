package main

import (
	"fmt"
	"reflect"
	"testing"
)

var twoSumTestCases = []struct {
	inputNums   []int
	inputTarget int
	output      []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{2, 5, 5, 11}, 10, []int{1, 2}},
}

func TestTwoSum(t *testing.T) {
	solutions := []struct {
		fnName string
		fnObj  func(nums []int, target int) []int
	}{
		{"twoSum", twoSum},
		{"twoSum2", twoSum2},
	}
	for _, sol := range solutions {
		for _, tc := range twoSumTestCases {
			t.Run(
				sol.fnName+fmt.Sprint(tc.inputNums),
				func(t *testing.T) {
					res := sol.fnObj(tc.inputNums, tc.inputTarget)
					if !reflect.DeepEqual(res, tc.output) {
						t.Errorf("got %v, want %v", res, tc.output)
					}
				})
		}
	}
}
