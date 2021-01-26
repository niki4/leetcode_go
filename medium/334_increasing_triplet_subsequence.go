/**
Given an integer array nums, return true if there exists a triple of indices (i, j, k) such that
i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

Example 1:
Input: nums = [1,2,3,4,5]
Output: true
Explanation: Any triplet where i < j < k is valid.

Example 2:
Input: nums = [5,4,3,2,1]
Output: false
Explanation: No triplet exists.
*/

package main

import "math"

// Linear scan with tracking prev (min) nums
// Runtime: 4 ms, faster than 92.68% of Go
// Memory Usage: 3 MB, less than 22.65% of Go
// Time complexity: O(n)
// Space complexity: O(1)
func increasingTriplet(nums []int) bool {
	first, second := math.Inf(1), math.Inf(1)
	for i := range nums {
		n := float64(nums[i])
		if n <= first {
			first = n
		} else if n <= second {
			second = n
		} else {
			return true
		}
	}
	return false
}
