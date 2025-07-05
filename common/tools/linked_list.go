package tools

import (
	"github.com/niki4/leetcode_go/common/types"
)

// Helper function to create a linked list from an slice
func CreateLinkedList(values []int) *types.ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &types.ListNode{Val: values[0]}
	current := head

	for i := 1; i < len(values); i++ {
		current.Next = &types.ListNode{Val: values[i]}
		current = current.Next
	}

	return head
}

// CreateSliceOfLinkedList creates a slice of linked lists from 2D slice
func CreateSliceOfLinkedList(lists [][]int) []*types.ListNode {
	result := make([]*types.ListNode, len(lists))
	for i, values := range lists {
		result[i] = CreateLinkedList(values)
	}
	return result
}

// TraverseLinkedList traverses Linked List from its head till the tail and returns list of values
func TraverseLinkedList(list *types.ListNode) []int {
	result := []int{}
	for list != nil {
		result = append(result, list.Val)
		list = list.Next
	}
	return result
}

// FlattenListValues iterates over list of Linked Lists and returns flatten list of all their values  
func FlattenListValues(lists []*types.ListNode) []int {
	listsValues := []int{}
	for _, head := range lists {
		listsValues = append(listsValues, TraverseLinkedList(head)...)
	}
	return listsValues
}
