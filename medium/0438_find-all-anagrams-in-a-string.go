/*
Given two strings s and p, return an array of all the start indices of p's anagrams in s.
You may return the answer in any order.

Example 1:
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".

Example 2:
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".

Constraints:
1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.
*/

package medium

func findAnagrams(s string, p string) []int {
	result := []int{}

	// Edge case: if s is smaller than p, no anagrams are possible
	if len(s) < len(p) {
		return result
	}

	// Frequency maps for p and the current window in s
	freqS, freqP := [26]int{}, [26]int{}
	for i := range p {
		freqP[p[i]-'a']++
	}

	// Sliding window over s
	for i := 0; i < len(s); i++ {
		// Add the current character to the window
		freqS[s[i]-'a']++

		if i >= len(p) {
			// maintain window size by removing the leftmost char from our window
			freqS[s[i-len(p)]-'a']--
		}

		if freqS == freqP {
			result = append(result, i-len(p)+1)
		}
	}

	return result
}
