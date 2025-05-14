/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

package main

func IsValidParenthesesSolution1(s string) bool {
	return isValid(s)
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2 MB, less than 41.97% of Go
func isValid(s string) bool {
	stack := make([]rune, 0)
	counter := 0
	opposite := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}
	for _, char := range s {
		switch char {
		case '{', '[', '(':
			stack = append(stack, char)
			counter++
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != opposite[char] {
				return false
			}
			stack = stack[:len(stack)-1]
			counter--
		}
	}
	return len(stack) == counter && counter == 0
}
