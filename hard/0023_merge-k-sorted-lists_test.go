package hard

import (
	"fmt"
	"reflect"
	"testing"

	. "github.com/niki4/leetcode_go/common/tools" //lint:ignore ST1001 dotted import
	. "github.com/niki4/leetcode_go/common/types" //lint:ignore ST1001 dotted import
)

var mergeKListsSolutions = map[string]func(lists []*ListNode) *ListNode{
	"mergeKLists": mergeKLists,
	"mergeKLists2": mergeKLists2,
}

var mergeKListsTestCases = []struct {
	input  [][]int
	output []int
}{
	{
		input:  [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
		output: []int{1, 1, 2, 3, 4, 4, 5, 6},
	},
	{
		input:  [][]int{},
		output: []int{},
	},
	{
		input:  [][]int{{}},
		output: []int{},
	},
	{
		input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		output: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		input:  [][]int{{1}, {2}, {3}},
		output: []int{1, 2, 3},
	},
	{
		input:  [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}},
		output: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
	},
	{
		input:  [][]int{{-10, -5, 0}, {-8, -3, 2}, {-6, -1, 4}},
		output: []int{-10, -8, -6, -5, -3, -1, 0, 2, 4},
	},
	{
		input:  [][]int{{1, 2, 3}, {}, {4, 5, 6}},
		output: []int{1, 2, 3, 4, 5, 6},
	},
	{
		input:  [][]int{{}, {}, {}},
		output: []int{},
	},
}

func TestMergeKLists(t *testing.T) {
	for solutionName, solution := range mergeKListsSolutions {
		for i, tc := range mergeKListsTestCases {
			t.Run(fmt.Sprintf("TestCase_%d", i+1), func(t *testing.T) {
				lists := CreateSliceOfLinkedList(tc.input)
				result := solution(lists)

				// Convert result back to array for comparison
				resultArray := TraverseLinkedList(result)

				if !reflect.DeepEqual(resultArray, tc.output) {
					t.Errorf("%s(%v) = %v, want %v", solutionName, tc.input, resultArray, tc.output)
				}
			})
		}
	}
}

// Test edge cases
func TestMergeKListsEdgeCases(t *testing.T) {

	for solutionName, solution := range mergeKListsSolutions {
		t.Run("SingleList", func(t *testing.T) {
			input := [][]int{{1, 2, 3, 4, 5}}
			expected := []int{1, 2, 3, 4, 5}

			lists := CreateSliceOfLinkedList(input)
			result := solution(lists)
			resultArray := TraverseLinkedList(result)

			if !reflect.DeepEqual(resultArray, expected) {
				t.Errorf("%s(%v) = %v, want %v", solutionName, input, resultArray, expected)
			}
		})

		t.Run("AllEmptyLists", func(t *testing.T) {
			input := [][]int{{}, {}, {}, {}}
			expected := []int{}

			lists := CreateSliceOfLinkedList(input)
			result := solution(lists)
			resultArray := TraverseLinkedList(result)

			if !reflect.DeepEqual(resultArray, expected) {
				t.Errorf("%s(%v) = %v, want %v", solutionName, input, resultArray, expected)
			}
		})

		t.Run("MixedEmptyAndNonEmpty", func(t *testing.T) {
			input := [][]int{{}, {1, 2}, {}, {3, 4}, {}}
			expected := []int{1, 2, 3, 4}

			lists := CreateSliceOfLinkedList(input)
			result := solution(lists)
			resultArray := TraverseLinkedList(result)

			if !reflect.DeepEqual(resultArray, expected) {
				t.Errorf("%s(%v) = %v, want %v", solutionName, input, resultArray, expected)
			}
		})

	}

}
