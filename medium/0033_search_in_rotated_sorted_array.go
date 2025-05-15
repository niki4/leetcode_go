/*
You are given an integer array nums sorted in ascending order, and an integer target.

Suppose that nums is rotated at some pivot unknown to you beforehand (
i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

If target is found in the array return its index, otherwise, return -1.

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
Example 3:

Input: nums = [1], target = 0
Output: -1
*/

package main

// Binary Search
// Time Complexity: O(log n)
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2

        if nums[mid] == target {
            // we found it, yeah!
            return mid
        }

        // First we define in which sorted sequence we are, then define is there target.
        if nums[left] <= nums[mid] {  // left sequence sorted
            if nums[left] <= target && target <= nums[mid] {
                // target in left sequence
                right = mid - 1
            } else {
                // target in right sequence
                left = mid + 1
            }
        } else {    // right sequence sorted
            if nums[mid] <= target && target <= nums[right] {
                // target in right sequence
                left = mid + 1
            } else {
                // target in left sequence
                right = mid - 1
            }
        }
    }

    return -1
}
