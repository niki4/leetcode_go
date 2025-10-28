/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

Example 1:
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.

Example 2:
Input: nums = [0,1,0,3,2,3]
Output: 4

Example 3:
Input: nums = [7,7,7,7,7,7,7]
Output: 1

Constraints:
1 <= nums.length <= 2500
-104 <= nums[i] <= 104
*/

package medium

func maxInSlice(nums []int) int {
    maxInt := nums[0]

    for _, n := range nums {
        if n > maxInt {
            maxInt = n
        }
    }
    return maxInt
}

// DP bottom-up approach
// Time complexixty O(n*2)
// Space complexity O(n)
func lengthOfLIS(nums []int) int {
    dp := make([]int, len(nums))
    for i := range dp {
        dp[i] = 1 // there's at least one subseq consisting from single num
    }

    // for each num, start and calculate increasing subseq, take whatever is max (old one or new one)
    for i := len(nums)-1; i > -1; i-- {
        for j := i+1; j < len(nums); j++ {
            if nums[i] < nums[j] {
                dp[i] = max(dp[i], 1 + dp[j])  // dp[i] hold len of increasing subseq starting from i-th element
            }
        }
    }

    return maxInSlice(dp)
}
