package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 40 ms, faster than 83.97% of Go.
// Memory Usage: 8.5 MB, less than 26.82% of Go.
// Time complexity: O(m+n)
// Space complexity: O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ptrA := headA
	ptrB := headB
	for ptrA != ptrB {
		if ptrA != nil {
			ptrA = ptrA.Next
		} else {
			ptrA = headB
		}
		if ptrB != nil {
			ptrB = ptrB.Next
		} else {
			ptrB = headA
		}
	}
	return ptrA
}
