/**
Count the number of prime numbers less than a non-negative number, n.

Example 1:
Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
*/

package main

// Sieve of Eratosthenes (https://en.wikipedia.org/wiki/Sieve_of_Eratosthenes)
// Time complexity: O(n log log n)
// Space complexity: O(n)
// Runtime: 4 ms, faster than 98.91% of Go.
// Memory Usage: 4.9 MB, less than 35.04% of Go.
func countPrimes(n int) (pCounter int) {
	if n <= 1 {
		return
	}

	sieve := make([]bool, n)
	sieve[0], sieve[1] = true, true

	for i := 2; i*i < n; i++ {
		if !sieve[i] {
			for j := 2 * i; j < n; j += i {
				sieve[j] = true
			}
		}
	}

	for _, isPrime := range sieve {
		if !isPrime {
			pCounter++
		}
	}
	return
}
