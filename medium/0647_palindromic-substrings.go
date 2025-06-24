/*
Given a string s, return the number of palindromic substrings in it.
A string is a palindrome when it reads the same backward as forward.
A substring is a contiguous sequence of characters within the string.

Example 1:
Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".

Example 2:
Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".

Constraints:
1 <= s.length <= 1000
s consists of lowercase English letters.
*/

package main

// Time complexity: O(n^2)
func countSubstrings(s string) int {
	counter := 0

	expandAroundCenter := func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			counter++
			left--  // go left toward begin
			right++ // go right toward end
		}
	}

	for i := range s {
		expandAroundCenter(i, i)   // for odd-length palindromes, e.g. "a" and "aaa"
		expandAroundCenter(i, i+1) // for even-length palindromes, e.g. "aa"
	}

	return counter
}
