/**
Given an array of integers nums containing n + 1 integers where each integer is in the range [1, n] inclusive.

There is only one repeated number in nums, return this repeated number.

You must solve the problem without modifying the array nums and using only constant extra space.

Example 1:
Input: nums = [1,3,4,2,2]
Output: 2

Example 2:
Input: nums = [3,1,3,4,2]
Output: 3

Example 3:
Input: nums = [3,3,3,3,3]
Output: 3
**/

// Time complexity: O(n); Space complexity: O(n) because of memoizing past nums
func findDuplicate(nums []int) int {
    seen := make(map[int]struct{})

    for _, num := range nums {
        if _, ok := seen[num]; ok {
            return num
        }
        seen[num] = struct{}{}
    }
    return -1
}

// Time complexity: O(n); Space complexity: O(1)
// Works for most inputs, but fails on the following one (doesn't fit in bitmask)
// [75,75,75,75,75,91,75,75,75,75,75,75,30,75,75,78,75,75,75,75,75,7,79,93,75,75,15,75,75,75,75,75,75,4,75,75,21,75,75,19,75,59,75,76,75,75,75,75,75,75,75,33,75,75,75,58,75,75,5,75,97,81,48,42,75,87,75,75,25,27,94,20,75,75,75,29,75,75,86,67,75,75,75,65,75,75,75,75,75,39,75,56,75,75,75,75,3,75,75,75]
func findDuplicate2(nums []int) int {
	bitMask := 1
	for _, n := range nums {
		// Inside the loop, this line checks if the nth bit in bitMask is set (i.e., if it equals 1).
		// It does this by shifting bitMask to the right by n positions and then performing a
		// bitwise AND operation with 1. If the result is not 0, it means that the nth bit was set,
		// indicating that the number n has already been seen in the slice.
		if ((bitMask >> n) & 1) != 0 {
			return n
		}
    // If the number n was not previously seen, this line sets the nth bit in bitMask.
    // It shifts 1 to the left by n positions (placing a 1 in the nth bit position) and then performs a bitwise OR operation with bitMask.
    // This operation updates bitMask to mark the number n as seen.
		bitMask |= (1 << n)
	}
	// Return a default value if no duplicate is found
	return -1
}
