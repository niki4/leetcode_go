/*
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:

Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

*/

package easy

import (
	"fmt"

	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

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

func createLinkedList(nums []int) *ListNode {
	var head *ListNode
	var prev *ListNode

	for _, num := range nums {
		if head == nil {
			head = &ListNode{Val: num}
			prev = head
		} else {
			prev.Next = &ListNode{Val: num}
			prev = prev.Next
		}

	}
	return head
}

func printLinkedList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Print(curr.Val)
		if curr.Next != nil {
			fmt.Print("->")
		}
		curr = curr.Next
	}
	fmt.Println("")
}

// func main() {
// 	inp := []int{1, 2, 3, 4, 5}
// 	head := createLinkedList(inp)
// 	printLinkedList(head)
// 	newHead := reverseList(head)
// 	printLinkedList(newHead)
// }
