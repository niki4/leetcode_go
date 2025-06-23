/* Given an array of strings strs, group the anagrams together. You can return the answer in any order.

Example 1:
	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
	Explanation:
	There is no string in strs that can be rearranged to form "bat".
	The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
	The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

Example 2:
	Input: strs = [""]
	Output: [[""]]

Example 3:
	Input: strs = ["a"]
	Output: [["a"]]

Constraints:
	1 <= strs.length <= 104
	0 <= strs[i].length <= 100
	strs[i] consists of lowercase English letters.
*/

package main

import "sort"

// sorts chars in a string, alphabetical order
func sortString(s string) string {
	chars := []rune(s)

	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	result := make([][]string, 0)

	// first define groups based on its canonical anagram
	for _, word := range strs {
		anagram := sortString(word)
		groups[anagram] = append(groups[anagram], word)
	}

	// then return groups only
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}

/*
Time Complexity: O(n * m log m), where n is the number of strings and m is the average length of the strings.
	Sorting all strings: O(n * m log m).
	Grouping strings: O(n).
	Constructing the result: O(n).
Space Complexity: O(n * m).
*/
