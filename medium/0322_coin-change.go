/*
You are given an integer array coins representing coins of different denominations and an integer amount representing a total amount of money.

Return the fewest number of coins that you need to make up that amount.
If that amount of money cannot be made up by any combination of the coins, return -1.

You may assume that you have an infinite number of each kind of coin.

Example 1:
Input: coins = [1,2,5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1

Example 2:
Input: coins = [2], amount = 3
Output: -1

Example 3:
Input: coins = [1], amount = 0
Output: 0

Constraints:
	1 <= coins.length <= 12
	1 <= coins[i] <= 231 - 1
	0 <= amount <= 104
*/

package main

// Dynamic programming
// Time Complexity: O(len(coins) * amount)
// Space Complexity: O(amount)
func coinChange(coins []int, amount int) int {
	// Create a DP array to store the minimum coins for each amount
	dp := make([]int, amount+1)

	// Init a DP array with a large value (representing infinity)
	for i := range dp {
		dp[i] = amount + 1
	}

	// Base case. 0 coins needed to make 0 amount.
	dp[0] = 0

	// Now calculate minimal amount of coins needed for each denomination
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(
				dp[i],        // prev denomination coin amount
				dp[i-coin]+1, // curr denomination coin (1 pc) that could replace prev denomination coin amount
			)
		}
	}

	// If dp[amount] is still has init value, it means the amount cannot be formed
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]
}

// Some additional details on the Python implementation:
// https://github.com/niki4/leetcode_py3/blob/master/medium/0322_coin_change.py
