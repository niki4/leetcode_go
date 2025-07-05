/*
Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string "".

Example 1:
Input: strs = ["flower","flow","flight"]
Output: "fl"

Example 2:
Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

Constraints:
	1 <= strs.length <= 200
	0 <= strs[i].length <= 200
	strs[i] consists of only lowercase English letters if it is non-empty.
*/

package easy

// Time complexity: O(n * m^2)
// where n is the number of strings in the input array strs
// and m is the length of the shortest string in the array strs.
func longestCommonPrefix(strs []string) string {
	// Find a word with minimal length. It will be the candidate for longest prefix.
	minLengthWord := strs[0]
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(minLengthWord) {
			minLengthWord = strs[i]
		}
	}

	// Find longest common prefix by comparing minimal length word with the rest ones in input.
	// At each (outer) iteration we shrink that example word from the end.
	for i := len(minLengthWord); i > 0; i-- {
		matches := 0
		for _, word := range strs {
			if word[:i] == minLengthWord[:i] {
				matches++
			}
		}
		if matches == len(strs) { // all words in input share same prefix
			return minLengthWord[:i]
		}
	}

	return "" // seems not found a prefix that would share ALL words in input
}
