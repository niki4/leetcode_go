package medium

import (
	"reflect"
	"testing"
)

var oddEvenListTestCases = []struct {
	input  []int
	output []int
}{
	{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5, 2, 4}},
	{[]int{2, 1, 3, 5, 6, 4, 7}, []int{2, 3, 6, 7, 1, 5, 4}},
}

func TestOddEvenList(t *testing.T) {
	for _, tc := range oddEvenListTestCases {
		inpLL := MakeLinkedListFromSlice(tc.input)
		res := MakeSliceFromLinkedList(oddEvenList(inpLL))
		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("for input=%v\nexpected:\t%v\ngot:\t\t%v", tc.input, tc.output, res)
		}
	}
}
