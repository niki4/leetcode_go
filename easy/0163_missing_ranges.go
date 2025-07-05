package easy

import "fmt"

func addRange(lower, upper int) string {
	if lower+2 == upper { // missed one num
		return fmt.Sprint(lower + 1)
	} else {
		return fmt.Sprintf("%d->%d", lower+1, upper-1)
	}
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2 MB, less than 52.38% of Go
// Time complexity : O(N), where N is the length of the input array.
// Space complexity : O(N) if we take the output into account and O(1) otherwise
func findMissingRanges(nums []int, lower int, upper int) []string {
	result := make([]string, 0)
	nums = append([]int{lower - 1}, nums...)
	nums = append(nums, upper+1)
	for i := 1; i < len(nums); i++ {
		prev, curr := nums[i-1], nums[i]
		if curr-prev > 1 { // missed num(s)
			result = append(result, addRange(prev, curr))
		}
	}
	return result
}
