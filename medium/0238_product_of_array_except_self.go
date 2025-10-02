/**
Given an integer array nums, return an array answer such that answer[i] is equal
to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Example 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]
**/

package medium

import "fmt"

type testCase struct {
	input  []int
	output []int
}

func compareSlices(first, second []int) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}
	return true
}

func createPrepopulatedSlice(length, value int) []int {
	slice := make([]int, length)
	for i := range slice {
		slice[i] = value
	}
	return slice
}

// Time complexity: O(n)
// Space complexity: O(1) - output array does not count as extra space as per problem description.
func productExceptSelf(nums []int) []int {
	result := createPrepopulatedSlice(len(nums), 1) // e.g., [1, 1, 1, 1]
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1] // left prod
		// e.g., for 'nums' [1, 2, 3, 4] we get 'result' [1, 1, 2, 6]
		// here last 3 items calculated as 1*1=1, 2*1=2, 3*2=6
	}

	rightProd := 1
	for i := len(nums) - 1; i >= 0; i-- {
		// while calc product from right side, we define product without current num
		result[i] *= rightProd
		rightProd *= nums[i]
	}
	// 'result' [24,12,8,6]
	return result
}

// Another implementation of the same logic
// Time complexity: O(n)
// Space complexity: O(1) - output array does not count as extra space as per problem description.
func productExceptSelf2(nums []int) []int {
    result := make([]int, len(nums))

    prefix := 1
    for i := 0; i < len(nums); i++ {
        result[i] = prefix  // first store prefix value
        prefix *= nums[i]
    }

    postfix := 1
    for j := len(nums)-1; j >= 0; j-- {
        result[j] *= postfix // here we multiply both prefix and postfix (so product of numbers before and after current number)
        postfix *= nums[j]
    }

    return result
}

func main() {
	testCases := []testCase{
		testCase{input: []int{1, 2, 3, 4}, output: []int{24, 12, 8, 6}},
		testCase{input: []int{-1, 1, 0, -3, 3}, output: []int{0, 0, 9, 0, 0}},
	}

	for _, tc := range testCases {
		result := productExceptSelf(tc.input)
		fmt.Println(compareSlices(result, tc.output))
	}
}
