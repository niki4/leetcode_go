/*
You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

Example 1:
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6

Example 2:
Input: lists = []
Output: []

Example 3:
Input: lists = [[]]
Output: []

Constraints:
	k == lists.length
	0 <= k <= 104
	0 <= lists[i].length <= 500
	-104 <= lists[i][j] <= 104
	lists[i] is sorted in ascending order.
	The sum of lists[i].length will not exceed 104.
*/

package hard

import (
	"slices"

	. "github.com/niki4/leetcode_go/common/tools" //lint:ignore ST1001 dotted import
	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

// Flatten, Sort, Rebuild
// Time complexity: O(n log n)
// Space complexity: O(n)
// n is the total number of nodes in all lists
func mergeKLists(lists []*ListNode) *ListNode {
	values := FlattenListValues(lists)
	slices.Sort(values)

	preHead := &ListNode{}
	curr := preHead

	// generate new list from sorted flatten values
	for _, num := range values {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	// return head of the new sorted flatten list
	return preHead.Next
}

// Iterative Merging
// Time Complexity: O(NlogK)
// Space complexity: O(1) ignoring the output list, as it merges in place
// N is the total number of nodes across all linked lists and K is the number of linked lists. 
func mergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		mergedList := MergeTwoSortedLinkedLists(lists[0], lists[1])
		lists = lists[2:]
		lists = append(lists, mergedList)
	}

	return lists[0]
}
