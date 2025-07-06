/*
Given the head of a linked list, remove the nth node from the end of the list and return its head.

Example 1:
Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]

Example 2:
Input: head = [1], n = 1
Output: []

Example 3:
Input: head = [1,2], n = 1
Output: [1]

Constraints:
The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz

Follow up: Could you do this in one pass?
*/

package medium

import . "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import

// One-pass, but extra space for storing refs to ListNode's
// Time/Space complexity: O(n)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	list := []*ListNode{}
	curr := head

	// track all nodes positions
	for curr != nil {
		list = append(list, curr)
		curr = curr.Next
	}

	// corner cases
	targetIdx := len(list) - n
	if len(list) == 0 || targetIdx < 0 || targetIdx > len(list)-1 {
		return nil
	}

	// case when target is current head
	if targetIdx == 0 {
		return head.Next
	}

	// "delete" the node by simply jump over it
	prev := list[targetIdx-1]
	prev.Next = prev.Next.Next
	return head
}

// 2-pointer approach, 1-pass
// Time complexity: O(n)
// Space complexity: O(1)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	preHead := &ListNode{}
	preHead.Next = head

	slow, fast := preHead, preHead

	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	// Now, having calculated the gap, move both pointers till the end.
	// Once the 'fast' one will hit the end, the 'slow' will mark Nth node from the end.
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	// To 'remove' node, simply jump over it
	slow.Next = slow.Next.Next

	return preHead.Next
}
