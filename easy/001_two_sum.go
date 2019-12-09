// Runtime: 36 ms, faster than 31.67% of Go online submissions for Two Sum.
// Memory Usage: 3 MB, less than 100.00% of Go online submissions for Two Sum.
// https://leetcode.com/submissions/detail/284883676/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	nms := []int{2, 7, 11, 15}
	trg := 9
	result := twoSum(nms, trg)
	fmt.Printf("res1: %#v\n", result) // expected [0, 1]

	nms = []int{2, 5, 5, 11}
	trg = 10
	result = twoSum(nms, trg)
	fmt.Printf("res2: %#v", result) // expected [1, 2]
}
