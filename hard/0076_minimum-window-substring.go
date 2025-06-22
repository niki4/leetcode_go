/*
Given two strings s and t of lengths m and n respectively, return the minimum window substring of s
such that every character in t (including duplicates) is included in the window.
If there is no such substring, return the empty string "".

The testcases will be generated such that the answer is unique.

Example 1:
    Input: s = "ADOBECODEBANC", t = "ABC"
    Output: "BANC"
    Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.

Example 2:
    Input: s = "a", t = "a"
    Output: "a"
    Explanation: The entire string s is the minimum window.

Example 3:
    Input: s = "a", t = "aa"
    Output: ""
    Explanation: Both 'a's from t must be included in the window.
    Since the largest window of s only has one 'a', return empty string.

Constraints:
    m == s.length
    n == t.length
    1 <= m, n <= 105
    s and t consist of uppercase and lowercase English letters.

Follow up: Could you find an algorithm that runs in O(m + n) time?
*/

package main

// Sliding Window
// The idea is to compare frequencies of chars between t and window while expanding/shrinking the window.
// Time/Space complexity: O(m+n)
func minWindow(s string, t string) string {
	// Step 1. Edge cases
	if len(s) == 0 || len(t) == 0 || len(s) < len(t) {
		return ""
	}

	// Step 2. Calculate frequency of the symbols in t
	targetFreq := make(map[rune]int)
	for _, char := range t {
		targetFreq[char]++
	}

	// Step 3. Slide window vars
	windowFreq := make(map[rune]int)
	want := len(targetFreq) // number of unique letters in t
	have := 0               // number of matches with target freq (both letter and its freq)
	left := 0               // left pointer
	minLen := len(s) + 1    // minimal window (size) that would contain the match
	resStart := 0           // index of the minimal window that contain the match

	for right, rightChar := range s {
		// expand the window
		windowFreq[rightChar]++

		// check if current char meets requirements
		if targetFreq[rightChar] > 0 && targetFreq[rightChar] == windowFreq[rightChar] {
			have++
		}

		// shrink window (from 'left') while matching conditions met
		for have == want {
			// update result if current window is smaller
			windowSize := right - left + 1
			if windowSize < minLen {
				minLen = windowSize
				resStart = left
			}

			// remove leftmost char from the window
			leftChar := rune(s[left])
			windowFreq[leftChar]--
			if targetFreq[leftChar] > 0 && windowFreq[leftChar] < targetFreq[leftChar] {
				have--
			}
			left++
		}
	}

	// Step 4. Check if minLen was updated from default (so there was some match)
	if minLen == len(s)+1 {
		return ""
	}

	// Step 5. Return the minumum window substring
	return s[resStart : resStart+minLen]
}
