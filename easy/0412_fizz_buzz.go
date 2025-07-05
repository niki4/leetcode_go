/**
Write a program that outputs the string representation of numbers from 1 to n.

But for multiples of three it should output “Fizz” instead of the number and for the multiples of five output “Buzz”.
For numbers which are multiples of both three and five output “FizzBuzz”.
*/

package easy

import "strconv"

// Runtime: 4 ms, faster than 97.62% of Go
// Memory Usage: 3.4 MB, less than 100.00% of Go
func fizzBuzz(n int) []string {
	res := make([]string, n)
	var val string
	for i := 1; i < n+1; i++ {
		if i%3 == 0 && i%5 == 0 {
			val = "FizzBuzz"
		} else if i%3 == 0 {
			val = "Fizz"
		} else if i%5 == 0 {
			val = "Buzz"
		} else {
			val = strconv.Itoa(i)
		}
		res[i-1] = val
	}
	return res
}
