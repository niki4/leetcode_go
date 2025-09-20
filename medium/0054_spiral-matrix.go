/*
Given an m x n matrix, return all elements of the matrix in spiral order.

Example 1: https://assets.leetcode.com/uploads/2020/11/13/spiral1.jpg
Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,3,6,9,8,7,4,5]

Example 2: https://assets.leetcode.com/uploads/2020/11/13/spiral.jpg
Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]

Constraints:
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

package medium

// Time/Space complexity: O(n)
func spiralOrder(matrix [][]int) []int {
    spiral := []int{}
    
    if len(matrix) == 0 {
        return spiral
    }

    m, n := len(matrix), len(matrix[0])
    size := m * n
    r1, c1 := 0, 0        // top left
    r2, c2 := m-1, n-1    // bottom right

    for len(spiral) < size {
        // go right
        for i := c1; i <= c2 && len(spiral) < size; i++ {
            spiral = append(spiral, matrix[r1][i])
        }
        r1++ // skip that (top) row next time

        // go down
        for j := r1; j <= r2 && len(spiral) < size; j++ {
            spiral = append(spiral, matrix[j][c2])
        }
        c2-- // skip that (right) column next time

        // go left
        for k := c2; k >= c1 && len(spiral) < size; k-- {
            spiral = append(spiral, matrix[r2][k])
        }
        r2-- // skip that (bottom) row next time

        // go up
        for l := r2; l >= r1 && len(spiral) < size; l-- {
            spiral = append(spiral, matrix[l][c1])
        }
        c1++ // skip that (left) column next time
    }

    return spiral
}
