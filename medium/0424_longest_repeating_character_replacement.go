/*
You are given a string s and an integer k.
You can choose any character of the string and change it to any other uppercase English character.
You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing the above operations.

Example 1:
Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:
Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.
There may exists other ways to achieve this answer too.

Constraints:
1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
*/

// https://leetcode.com/problems/longest-repeating-character-replacement/

package main

// Sliding window
// Time complexity: O(n)
// Space complexity: O(1)
// n - length of the string
func characterReplacement(s string, k int) int {
	left, maxCharCount, maxLength := 0, 0, 0
	frequency := make(map[rune]int, 26) // pre-allocate space for eng alphabet

	for right, char := range s {
		frequency[char]++
		maxCharCount = max(maxCharCount, frequency[char])

		windowSize := right - left + 1
		if windowSize-maxCharCount > k {
			// invalid state, cannot make substitutions more than k, need to shrink window
			frequency[rune(s[left])]--
			left++
		} else {
			// valid state
			maxLength = max(maxLength, windowSize)
		}
	}
	return maxLength
}

/* Q: Why we don't need to decrement maxCharCount when shrinking window?
A: Because we are only interested in the maximum frequency of any character in the current window.
When we shrink the window, the frequency of the character that was removed is no longer relevant to the maximum frequency.
We only need to check if the new window size minus the maximum frequency is greater than k to determine if the window is valid.

In other words, if you'd do that, there's chance when you shrink the window from the left, you decrement the count of
the character that’s leaving the window, but you don’t know if that character was the one with the maximum frequency
(If the character being removed is not the one with the highest frequency, or if there are multiple characters with
the same frequency, this decrement will make maxCharCount inaccurate).
*/
