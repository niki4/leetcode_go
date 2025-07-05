/**
Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
In Pascal's triangle, each number is the sum of the two numbers directly above it.

Example:
Input: 5
Output:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
*/

package easy

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2 MB, less than 16.23% of Go

func generatePascalTriangle(numRows int) [][]int {
	pyramid := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ { // skip first and last item as they're always "1"
			row[j] = pyramid[i-1][j-1] + pyramid[i-1][j]
		}
		pyramid[i] = row
	}
	return pyramid
}
