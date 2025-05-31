/*
Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

package main

/**
* Definition for singly-linked list.
* type ListNode struct {
* 	Val  int
* 	Next *ListNode
* }
 */

// O(n) time, O(n) space
// Hash Set
func hasCycle1(head *ListNode) bool {
	curr := head
	seen := map[*ListNode]struct{}{}

	for curr != nil {
		if _, ok := seen[curr]; ok {
			return true // we seen it before, literally
		}
		seen[curr] = struct{}{}

		curr = curr.Next
	}

	return false
}

// O(n) time, O(1) space
// Floyd's Cycle Detection Algorithm (Floyd's Tortoise and Hare algorithm)
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for slow != fast {
		if fast != nil && fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			return false
		}
	}
	return true
}
