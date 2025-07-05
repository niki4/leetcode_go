package easy

// Runtime: 36 ms, faster than 31.67% of Go online submissions for Two Sum.
// Memory Usage: 3 MB, less than 100.00% of Go online submissions for Two Sum.
// https://leetcode.com/submissions/detail/284883676/
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Runtime: 80 ms, faster than 14.90% of Go
// Memory Usage: 7.7 MB, less than 11.49% of Go
func twoSum2(nums []int, target int) []int {
	pairIdxs := make(map[int]int)
	for nIdx, nVal := range nums {
		if pIdx, ok := pairIdxs[nVal]; ok {
			return []int{pIdx, nIdx}
		} else {
			pairIdxs[target-nVal] = nIdx
		}
	}
	return []int{}
}
