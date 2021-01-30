package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceIntToInt(nums []int) int {
	numStr := strings.Trim(strings.Join(
		strings.Fields(
			fmt.Sprint(nums)), ""), "[]")
	numInt, _ := strconv.Atoi(numStr)
	return numInt
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var num1 []int
	var num2 []int
	for l1 != nil {
		num1 = append([]int{l1.Val}, num1...) // make reversed num
		l1 = l1.Next
	}
	for l2 != nil {
		num2 = append([]int{l2.Val}, num2...)
		l2 = l2.Next
	}
	sum := strconv.Itoa(sliceIntToInt(num1) + sliceIntToInt(num2))

	// create resulting linked list
	dummyNode := new(ListNode)
	prev := dummyNode
	for i := len(sum) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(sum[i]))
		fmt.Printf("n: %d (%[1]T)\n", n)
		node := &ListNode{
			Val: n,
		}
		prev.Next = node
		prev = node
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
