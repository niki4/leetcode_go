/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false

Constraints:
1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.

Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
*/

package main

func countChars(input string) map[rune]int {
	frequency := make(map[rune]int)

	for _, ch := range input {
		frequency[ch]++
	}
	return frequency
}

// Time complexity: O(n)
func isAnagram(s string, t string) bool {
	// assume both s and t are the same size
	if len(s) != len(t) {
		return false
	}

	sFreq := countChars(s)
	tFreq := countChars(t)

	for _, ch := range t {
		if sFreq[ch] != tFreq[ch] {
			return false
		}
	}

	return true
}
