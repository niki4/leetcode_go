// Runtime: 44 ms, faster than 20.43% of Go.
// Memory Usage: 8.5 MB, less than 30.00% of Go.
// https://leetcode.com/submissions/detail/227397350/

package main

import (
	"fmt"
)

type MinStack struct {
	values []int
	minIdx int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.values = append(this.values, x)
	this.minIdx = minIntIdx(this.values)
}

func (this *MinStack) Pop() {
	this.values = this.values[:len(this.values)-1]
	if len(this.values) > 0 {
		this.minIdx = minIntIdx(this.values)
	}
}

func (this *MinStack) Top() int {
	if len(this.values) > 0 {
		return this.values[len(this.values)-1]
	}

	return -1
}

func (this *MinStack) GetMin() int {
	return this.values[this.minIdx]
}

func (this *MinStack) string() string {
	return fmt.Sprintf("this: %v\n", this)
}

func minIntIdx(nums []int) int {
	minIdx := 0
	minVal := nums[minIdx]

	for i, v := range nums {
		if v < minVal {
			minVal = v
			minIdx = i
		}
	}
	return minIdx
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)

	if x := obj.GetMin(); x != -3 {
		fmt.Println("Test fail! Expected", -3, "got", x)
	}

	obj.Pop()

	if x := obj.Top(); x != 0 {
		fmt.Println("Test fail! Expected", 0, "got", x)
	}
	if x := obj.GetMin(); x != -2 {
		fmt.Println("Test fail! Expected", -2, "got", x)
	}
}
