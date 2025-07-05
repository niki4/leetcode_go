package medium

import (
	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

// Runtime: 8 ms, faster than 89.57% of Go
// Memory Usage: 5 MB, less than 20.20% of Go
// Time complexity: O(n)
// Space complexity: O(n) where n - max(len(l1), len(l2))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode, sum := new(ListNode), 0
	for cur := dummyNode; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummyNode.Next
}
