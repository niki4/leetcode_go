/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and
removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.

Example 1:
Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.

Example 2:
Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.

Example 3:
Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.

Constraints:
1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

import u "unicode"

// Time complexity: O(n)
// Space complexity: O(1)
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// skip non-letters from both ends
		for left < right && !u.IsLetter(rune(s[left])) && !u.IsNumber(rune(s[left])) {
			left++
		}
		for left < right && !u.IsLetter(rune(s[right])) && !u.IsNumber(rune(s[right])) {
			right--
		}

		// check if both ends are equal (so making palindrome), then continue
		if u.ToLower(rune(s[left])) != u.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}

	return true
}
