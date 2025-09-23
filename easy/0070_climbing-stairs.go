/*
You are climbing a staircase. It takes n steps to reach the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Example 1:
Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps

Example 2:
Input: n = 3
Output: 3

Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

Constraints:
1 <= n <= 45
*/

package easy

// Solution is similar to Fibonacci sequence
func climbStairs(n int) int {
	stairs := make([]int, n+1)
	for i := range stairs {
		stairs[i] = 1
	}

	// at each index `i` we store result at the specified `n`
	// so we don't need to recalculate previous results each time.
	for i := 2; i <= n; i++ {
		stairs[i] = stairs[i-2] + stairs[i-1]
	}

	return stairs[n]
}
