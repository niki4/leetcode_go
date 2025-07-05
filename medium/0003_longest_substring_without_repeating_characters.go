package medium

/**
Given a string s, find the length of the longest substring without repeating characters.

Example 1:
Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
*/

// Runtime: 228 ms, faster than 16.91% of Go
// Memory Usage: 6.5 MB, less than 17.54% of Go
// Time complexity: O(n**2)
// Space complexity: O(n)
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

// Runtime: 8 ms, faster than 67.23% of Go
// Memory Usage: 3.1 MB, less than 64.78% of Go
// Time complexity: O(n)
// Space complexity: O(n)
func lengthOfLongestSubstring2(s string) int {
	str := []rune(s)
	set := make(map[rune]struct{})
	ans, i, j := 0, 0, 0
	for i < len(str) && j < len(str) {
		if _, exist := set[str[j]]; !exist {
			set[str[j]] = struct{}{}
			j++
			if j-i > ans {
				ans = j - i
			}
		} else {
			delete(set, str[i])
			i++
		}
	}
	return ans
}

// Sliding window approach
// Time complexity: O(n)
// Space complexity: O(min(m, n))
func lengthOfLongestSubstring3(s string) int {
	seen := map[rune]int{} // mapping char to index it was seen
	i := 0                 // left pointer of sliding window
	maxLength := 0

	for j, char := range s {
		if index, ok := seen[char]; ok && index >= i {
			// found duplicate, shift left pointer to shrink the sliding window
			i = index + 1
		}
		seen[char] = j // remember current char position
		maxLength = max(maxLength, j-i+1)
	}
	return maxLength
}
