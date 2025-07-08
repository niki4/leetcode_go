/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.

Example 1:
Input: head = [1,2,3,4]
Output: [1,4,2,3]

Example 2:
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]

Constraints:
The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000
*/

func findMidPoint(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var curr *ListNode = head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func mergeLists(first, second *ListNode) {
	var temp1, temp2 *ListNode

	for first != nil && second != nil {
		// Remember before re-linking
		temp1, temp2 = first.Next, second.Next

		// Linking (merging)
		first.Next = second
		second.Next = temp1

		// Move current pointers
		first, second = temp1, temp2
	}
}

// Time complexity: O(n)
// Space complexity: O(1)
func reorderList(head *ListNode) {
	// E.g., for input 1->2->3->4-5(->nil)
	// 1. Find the middle point ("3")
	// 2. Reverse the list starting from that middle point to end: 5->4->3(->nil)
	// 3. Merge two lists (the first regular part and the second reversed part) in place (1->5->2->4->3)

	if head == nil {
		return
	}

	// Step 1: Find the middle of the linked list
	middle := findMidPoint(head)

	// Step 2: Reverse the second half of the list
	first := head
	second := reverseList(middle.Next)
	middle.Next = nil

	// Step 3: Merge the two halves
	mergeLists(first, second)
}
