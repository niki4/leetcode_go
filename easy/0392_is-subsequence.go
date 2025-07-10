/*
Given two strings s and t, return true if s is a subsequence of t, or false otherwise.

A subsequence of a string is a new string that is formed from the original string by
deleting some (can be none) of the characters without disturbing the relative positions
of the remaining characters. (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

Example 1:
Input: s = "abc", t = "ahbgdc"
Output: true

Example 2:
Input: s = "axc", t = "ahbgdc"
Output: false
*/

package easy

func isSubsequence(s string, t string) bool {
	found := 0
	pos := 0

	for i := range s {
		for pos < len(t) && s[i] != t[pos] {
			pos++
		}

		if pos < len(t) && s[i] == t[pos] {
			found++
			pos++
		}
	}

	return found == len(s)
}
