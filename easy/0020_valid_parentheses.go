/**
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
*/

package easy

func IsValidParenthesesSolution1(s string) bool {
	return isValid1(s)
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2 MB, less than 41.97% of Go
func isValid1(s string) bool {
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

// A bit more concise version than `isValid1`
// Time/Space complexity: O(n)
func isValid(s string) bool {
	stack := []rune{} // to keep the unclosed parentheses
	pairs := map[rune]rune{')': '(', '}': '{', ']': '['}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else { // if closing bracket, remove the opening one if valid
			if len(stack) == 0 || stack[len(stack)-1] != pairs[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	// if there's nothing in stack, string is balanced
	return len(stack) == 0
}
