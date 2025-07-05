// Runtime: 8 ms, faster than 100.00% of Go.
// Memory Usage: 5.2 MB, less than 42.48% of Go.
// https://leetcode.com/submissions/detail/227405497/

package easy

import (
	"strconv"
)

func isPalindrome(x int) bool {
	input := strconv.Itoa(x)

	for i, fullLen := 0, len(input)-1; i < fullLen; i += 1 {
		if input[i] != input[fullLen-i] {
			return false
		}
	}
	return true
}
