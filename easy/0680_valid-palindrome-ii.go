/*
Given a string s, return true if the s can be palindrome after
deleting at most one character from it.

Example 1:
Input: s = "aba"
Output: true

Example 2:
Input: s = "abca"
Output: true
Explanation: You could delete the character 'c'.

Example 3:
Input: s = "abc"
Output: false

Constraints:
1 <= s.length <= 105
s consists of lowercase English letters.
*/

package easy

func isPalindrome_(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// Time complexity: O(n)
// Space complexity: O(1)
func validPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left < right {
		if left < right && s[left] != s[right] {
			// check both options (is palindrome after deleting at most one character from it)
			return isPalindrome_(s, left+1, right) || isPalindrome_(s, left, right-1)
		}
		left++
		right--
	}
	return true
}
