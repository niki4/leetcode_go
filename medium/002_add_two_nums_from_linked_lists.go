package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Runtime: 8 ms, faster than 89.57% of Go
// Memory Usage: 5 MB, less than 20.20% of Go
// Time complexity: O(n)
// Space complexity: O(n) where n - max(len(l1), len(l2))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode, sum := new(ListNode), 0
	for cur := dummyNode; l1 != nil || l2 != nil || sum != 0; cur = cur.Next {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		sum /= 10
	}
	return dummyNode.Next
}

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}
	fmt.Printf("%#v\n", l1)
	fmt.Printf("%#v\n", l1.Next)
	fmt.Printf("%#v\n\n", l1.Next.Next)

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}
	fmt.Printf("%#v\n", l2)
	fmt.Printf("%#v\n", l2.Next)
	fmt.Printf("%#v\n\n", l2.Next.Next)

	res := addTwoNumbers(l1, l2)
	fmt.Printf("%#v\n", res)           // 7
	fmt.Printf("%#v\n", res.Next)      // 0
	fmt.Printf("%#v\n", res.Next.Next) // 8
}
