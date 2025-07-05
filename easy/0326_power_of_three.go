package easy

import "strconv"

// Runtime: 20 ms, faster than 84.00% of Go
// Memory Usage: 6.1 MB, less than 12.80% of Go
// Time complexity:  O(log3n) where 3 is desired power
// Space complexity: O(1)
func isPowerOfThree(n int) bool {
	/**
	Algorithm:
	While we can divide without remainder (n%3==0) - divide of num of its power until we get 1 (
	so 9 is 3*3, thus we doing 9/3 = 3, then 3/3 = 1) or not (45/3=15, 15/3=5, 5%3=2 -> 5 != 1 thus return False for 45)
	*/
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}

// Runtime: 16 ms, faster than 95.20% of Go
// Memory Usage: 6.1 MB, less than 12.80% of Go
// Time and Space complexity:  O(log3 n)
func isPowerOfThreeV2(n int) bool {
	/**
	Algorithm:
	In Base 10, all powers of 10 start with the digit 1 and then are followed only by 0 (e.g. 10, 100, 1000).
	This is true for other bases and their respective powers.
	Therefore if we convert our number to base 3 and the representation is of the form 100...0, then the number is a power of 3.

	e.g., FormatInt(int64(27), 3) == "1000" thus 27 is power of 3
	but   FormatInt(int64(56), 3) == "1200" since 56 is not power of 3
	*/
	base3 := strconv.FormatInt(int64(n), 3)
	if base3[:1] != "1" {
		return false
	}
	for _, d := range base3[1:] {
		if d != '0' {
			return false
		}
	}
	return true
}
