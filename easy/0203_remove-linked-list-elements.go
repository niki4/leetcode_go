/*
Given the head of a linked list and an integer val, remove all the nodes of the
linked list that has Node.val == val, and return the new head.

Example 1:
Input: head = [1,2,6,3,4,5,6], val = 6
Output: [1,2,3,4,5]

Example 2:
Input: head = [], val = 1
Output: []

Example 3:
Input: head = [7,7,7,7], val = 7
Output: []

Constraints:
The number of nodes in the list is in the range [0, 104].
1 <= Node.val <= 50
0 <= val <= 50
*/

package easy

import . "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import

// Time complexity: O(n)
// Space complexity: O(1)
func removeElements(head *ListNode, val int) *ListNode {
	preHead := &ListNode{Next: head}
	prev, curr := preHead, preHead.Next

	for curr != nil {
		// remove node that contain Val equal to input
		if curr.Val == val {
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = curr
			curr = curr.Next
		}
	}

	return preHead.Next
}
