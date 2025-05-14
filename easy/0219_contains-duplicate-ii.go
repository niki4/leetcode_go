package main

/*
Given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array
such that nums[i] = nums[j] and the absolute difference between i and j is at most k.
https://en.wikipedia.org/wiki/Absolute_difference

Example 1:
Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

// Runtime: 16 ms, faster than 70.79% of Go
// Memory Usage: 7.8 MB, less than 17.42% of Go
func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)
	for currIdx, v := range nums {
		if prevIdx, ok := seen[v]; ok && (currIdx-prevIdx) <= k {
			return true
		}
		seen[v] = currIdx
	}
	return false
}
