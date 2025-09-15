/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals,
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].

Example 2:
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

Example 3:
Input: intervals = [[4,7],[1,4]]
Output: [[1,7]]
Explanation: Intervals [1,4] and [4,7] are considered overlapping.

Constraints:

1 <= intervals.length <= 104
intervals[i].length == 2
0 <= starti <= endi <= 104
*/

package medium

import "sort"

func isOverlapping(i, j []int) bool {
  return i[0] <= j[1] && i[1] >= j[0]
}

// Time complexity: O(N logN)
func merge(intervals[][]int) [][]int {
  sort.Slice(
    intervals,
    func(a, b int) bool { return intervals[a][0] <= intervals[b][0] })

  merged := [][]int{ intervals[0] }

  for i := 1; i < len(intervals); i++ {
    prev := merged[len(merged)-1]
    curr := intervals[i]

    if isOverlapping(prev, curr) {
      merged[len(merged)-1] = []int{min(prev[0], curr[0]), max(prev[1], curr[1])}
    } else {
      merged = append(merged, curr)
    }
  }

  return merged
}
