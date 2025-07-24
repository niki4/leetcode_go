/*
Given an integer n, return the number of trailing zeroes in n!.

Note that n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1.

Example 1:

Input: n = 3
Output: 0
Explanation: 3! = 6, no trailing zero.
Example 2:

Input: n = 5
Output: 1
Explanation: 5! = 120, one trailing zero.
Example 3:

Input: n = 0
Output: 0


Constraints:
0 <= n <= 104

Follow up: Could you write a solution that works in logarithmic time complexity?
*/

package medium

// Time complexity: O(log(n))
// Space complexity: O(1)
// The idea behind the solution is to count multipliers of 5x2 in factorialwhich forms the trailing zeroes (5x2=10).
// As there's always more 2s than 5s, we can count only multipliers of 5.
// Best explanation is here: https://www.purplemath.com/modules/factzero.htm
func trailingZeroes(n int) int {
	zeroes := 0

	for {
		rest := n / 5
		if rest < 1 {
			break
		}

		zeroes += rest
		n = rest
	}

	return zeroes
}
