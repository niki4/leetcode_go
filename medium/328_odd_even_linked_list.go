package main

/*
Given a singly linked list, group all odd nodes together followed by the even nodes.
Please note here we are talking about the node number and not the value in the nodes.

You should try to do it in place. The program should run in O(1) space complexity and O(nodes) time complexity.

Example 1:
Input: 1->2->3->4->5->NULL
Output: 1->3->5->2->4->NULL

Example 2:
Input: 2->1->3->5->6->4->7->NULL
Output: 2->3->6->7->1->5->4->NULL

Constraints:
The relative order inside both the even and odd groups should remain as it was in the input.
The first node is considered odd, the second node even and so on ...
The length of the linked list is between [0, 10^4].
*/

// Runtime: 4 ms, faster than 86.67% of Go
// Memory Usage: 3.3 MB, less than 100.00% of Go
//    Time complexity: O(n)
//    Space complexity: O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := odd.Next
	evenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next // e.g., 1st to 3rd node (skip 2nd)
		odd = odd.Next
		even.Next = odd.Next // e.g., 2nd to 4th node (skip 3rd)
		even = even.Next
	}
	odd.Next = evenHead
	return head
}
