package main

import (
	"fmt"
)

// Runtime: 36 ms, faster than 31.67% of Go online submissions for Two Sum.
// Memory Usage: 3 MB, less than 100.00% of Go online submissions for Two Sum.
// https://leetcode.com/submissions/detail/284883676/
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

// Runtime: 80 ms, faster than 14.90% of Go
// Memory Usage: 7.7 MB, less than 11.49% of Go
func twoSum2(nums []int, target int) []int {
	pairIdxs := make(map[int]int)
	for nIdx, nVal := range nums {
		if pIdx, ok := pairIdxs[nVal]; ok {
			return []int{pIdx, nIdx}
		} else {
			pairIdxs[target-nVal] = nIdx
		}
	}
	return []int{}
}

func main() {
	nms := []int{2, 7, 11, 15}
	trg := 9
	result := twoSum(nms, trg)
	result2 := twoSum2(nms, trg)
	fmt.Printf("twoSum res1: %#v\n", result)   // expected [0, 1]
	fmt.Printf("twoSum2 res1: %#v\n", result2) // expected [0, 1]

	nms = []int{2, 5, 5, 11}
	trg = 10
	result = twoSum(nms, trg)
	result2 = twoSum(nms, trg)
	fmt.Printf("twoSum res2: %#v\n", result)   // expected [1, 2]
	fmt.Printf("twoSum2 res2: %#v\n", result2) // expected [1, 2]
}
