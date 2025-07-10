/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such
that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.

Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

Constraints:
3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

package medium

import "slices"

// Sorting + 2-pointer approach
// Time/space complexity: O(n^2)
func threeSum(nums []int) [][]int {
	result := [][]int{}

	slices.Sort(nums)

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates for the first number
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		// pointers for second and third number
		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				// found a triplet, add to result
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// skip duplicated nums (if any) from both sides
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right-1] == nums[right] {
					right--
				}

				// move both pointers to next non duplicating positions
				left++
				right--
			} else if sum < 0 {
				// sum is too small, increase it by moving left pointer forward
				left++
			} else {
				// sum is too big, decrease it by moving right pointer backward
				right--
			}
		}
	}

	return result
}
