package easy

/*
Given a string, find the first non-repeating character in it and return its index. If it doesn't exist, return -1.

Examples:
s = "leetcode"
return 0.
s = "loveleetcode"
return 2.

Note: You may assume the string contains only lowercase English letters.
*/

// Runtime: 24 ms, faster than 70.03% of Go
// Memory Usage: 5.7 MB, less than 50.95% of Go
// Runtime complexity: O(n)
// Space complexity: O(1) because we have constant size for dict (at most 26 eng alphabet letters)
func firstUniqChar(s string) int {
	counter := make(map[rune]int)
	for _, chr := range s {
		// if chr not exist Go will init it with default value (0) before ++ operation
		counter[chr]++
	}
	for i, ch := range s {
		if counter[ch] == 1 {
			return i
		}
	}
	return -1
}
