// Runtime: 12 ms, faster than 100.00% of Go.
// Memory Usage: 5.2 MB, less than 49.56% of Go.
// https://leetcode.com/submissions/detail/227401978/

package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	input := strconv.Itoa(x)
    fullLen := len(input) - 1
    halfLen := fullLen // 2

    for i := 0; i < halfLen; i += 1 {
        if input[i] != input[fullLen - i] {
            return false
        }
    }
    return true
}
