/*
Given an integer array nums, find a subarray that has the largest product, and return the product.

The test cases are generated so that the answer will fit in a 32-bit integer.

Example 1:
Input: nums = [2,3,-2,4] Output: 6
Explanation: [2,3] has the largest product 6.

Example 2:
Input: nums = [-2,0,-1] Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:

1 <= nums.length <= 2 * 104
-10 <= nums[i] <= 10
The product of any subarray of nums is guaranteed to fit in a 32-bit integer.
 */
package main

// Kadane's algorithm is used to find the maximum product subarray.
// The algorithm keeps track of the maximum and minimum products at each position,
// because a negative number can turn a small product into a large one.
// The global maximum product is updated at each step.
// Time Complexity: O(n), Space Complexity: O(1)
func maxProduct(nums []int) int {
	minProd, maxProd, globalMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]

		// If the current number is negative, swap maxProd and minProd
		if num < 0 {
			minProd, maxProd = maxProd, minProd
		}

		// Update maxProd and minProd
		maxProd = max(num, maxProd*num)
		minProd = min(num, minProd*num)

		// Update globalMax
		globalMax = max(globalMax, maxProd)
	}

	return globalMax
}

/* Test Case Walkthrough
Input: [2, 3, -2, 4]

Initialization:
minProd = 2, maxProd = 2, globalMax = 2

Iteration 1 (num = 3):
num > 0, no swap.
maxProd = max(3, 2*3) = 6
minProd = min(3, 2*3) = 3
globalMax = max(2, 6) = 6

Iteration 2 (num = -2):
num < 0, swap maxProd and minProd.
maxProd = max(-2, 3*(-2)) = -2
minProd = min(-2, 6*(-2)) = -12
globalMax = max(6, -2) = 6

Iteration 3 (num = 4):
num > 0, no swap.
maxProd = max(4, -2*4) = 4
minProd = min(4, -12*4) = -48
globalMax = max(6, 4) = 6

Final Result: globalMax = 6
*/