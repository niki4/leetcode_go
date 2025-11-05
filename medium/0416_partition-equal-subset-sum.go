/*
Given an integer array nums, return true if you can partition the array into two subsets
such that the sum of the elements in both subsets is equal or false otherwise.

Example 1:
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].

Example 2:
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.

Constraints:
1 <= nums.length <= 200
1 <= nums[i] <= 100
*/

package medium

// Dynamic Programming - 0/1 Knapsack
// The time complexity is O(N * S), where N is the number of elements and S is the total sum.
// Space complexity O(S)  where S is the total sum of the elements in the subset.
func canPartition(nums []int) bool {
    sumNums := 0
    for _, n := range nums {
        sumNums += n
    }

    if sumNums % 2 == 1 {  // odd sums cannot be devided to two equal whole int parts
        return false
    }

    halfSum := sumNums / 2  // halfSum is our target sum

    dp := make([]bool, halfSum+1)
    dp[0] = true // init value as we start top-down

    for _, num := range nums {
        for currSum := halfSum; currSum > num - 1; currSum-- {
            dp[currSum] = dp[currSum] || dp[currSum-num]
        }
    }

    return dp[halfSum]
}
