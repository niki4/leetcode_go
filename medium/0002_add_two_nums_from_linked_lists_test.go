package medium

import (
	"reflect"
	"testing"

	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

var addTwoNumbersTestCases = []struct {
	l1Values []int
	l2Values []int
	result   []int
}{
	{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
}

// MakeLinkedListFromSlice converts slice of integers to Linked List returning
// a head of such list. The order of values from original slice is kept in the list.
func MakeLinkedListFromSlice(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	dummyNode := new(ListNode)
	node := dummyNode
	for _, v := range values {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return dummyNode.Next
}

// MakeSliceFromLinkedList converts Linked List to a Slice of integers.
// The order of values from original Linked List is kept in the resulted slice.
func MakeSliceFromLinkedList(head *ListNode) (res []int) {
	node := head
	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}
	return
}

func TestAddTwoNumbers(t *testing.T) {
	for _, tc := range addTwoNumbersTestCases {
		inpL1 := MakeLinkedListFromSlice(tc.l1Values)
		inpL2 := MakeLinkedListFromSlice(tc.l2Values)
		res := MakeSliceFromLinkedList(addTwoNumbers(inpL1, inpL2))
		if !reflect.DeepEqual(res, tc.result) {
			t.Errorf("for input l1=%v l2=%v\nexpected:\t%v\ngot:\t\t%v",
				tc.l1Values, tc.l2Values, tc.result, res)
		}
	}
}
