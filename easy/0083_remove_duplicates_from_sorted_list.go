package easy

import (
	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 3.1 MB, less than 100.00% of Go
// Time complexity: O(n) because we need visit all nodes, remove dup node takes O(1)
// Space complexity: O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val { // found dup
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
