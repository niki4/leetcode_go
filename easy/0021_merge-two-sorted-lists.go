/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

Constraints:
	The number of nodes in both lists is in the range [0, 50].
	-100 <= Node.val <= 100
	Both list1 and list2 are sorted in non-decreasing order.
*/

package easy

import (
	. "github.com/niki4/leetcode_go/common/tools" //lint:ignore ST1001 dotted import
	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

// Two pointer approach
// Time/Space complexity: O(n + m)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	preHead := result

	// track position at each list and compare values as you go
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			result.Next = &ListNode{Val: list1.Val}
			result = result.Next
			list1 = list1.Next
		} else if list2.Val < list1.Val {
			result.Next = &ListNode{Val: list2.Val}
			result = result.Next
			list2 = list2.Next
		} else { // equal values, add both
			result.Next = &ListNode{Val: list1.Val, Next: &ListNode{Val: list2.Val}}
			result = result.Next.Next
			list1 = list1.Next
			list2 = list2.Next
		}
	}

	// handle case when there's something left in one of the lists
	for list1 != nil {
		result.Next = &ListNode{Val: list1.Val}
		result = result.Next
		list1 = list1.Next
	}
	for list2 != nil {
		result.Next = &ListNode{Val: list2.Val}
		result = result.Next
		list2 = list2.Next
	}

	// Return the head of the merged linked list.
	return preHead.Next
}

// Merge in place. Recursive approach.
// Time complexity: O(m+n)
// Space complexity: O(m+n) because of recursion stack
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			// swap to guarantee list1 always less or equal than list2
			list1, list2 = list2, list1
		}
		list1.Next = mergeTwoLists2(list1.Next, list2)
	}

	// base cases for recursion
	if list1 != nil {
		return list1
	}
	return list2
}

// Merge in place. Iterative approach.
// Time complexity: O(m+n)
// Space complexity: O(1) as we don't use extra space, just relinking nodes in proper order
func mergeTwoLists3(list1 *ListNode, list2 *ListNode) *ListNode {
	return MergeTwoSortedLinkedLists(list1, list2)
}
