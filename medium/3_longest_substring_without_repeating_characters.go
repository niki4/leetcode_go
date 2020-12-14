package main

/**
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

// Runtime: 228 ms, faster than 16.91% of Go
// Memory Usage: 6.5 MB, less than 17.54% of Go
func lengthOfLongestSubstring(s string) int {
	longestUniqueLen := 0
	str := []rune(s)
	seen := make(map[rune]struct{})

	for i := 0; i < len(str); i++ {
		end := i
		for j := i; j < len(str); j++ {
			if _, ok := seen[str[j]]; ok {
				seen = make(map[rune]struct{})
				break
			}
			seen[str[j]] = struct{}{}
			end++
		}
		substrLen := end - i
		if substrLen > longestUniqueLen {
			longestUniqueLen = substrLen
		}
	}
	return longestUniqueLen
}
