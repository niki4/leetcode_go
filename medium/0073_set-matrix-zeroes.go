/*
Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
You must do it in place.

Example 1:
Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]

Example 2:
Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

Constraints:
m == matrix.length
n == matrix[0].length
1 <= m, n <= 200
-231 <= matrix[i][j] <= 231 - 1

Follow up:
A straightforward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

package medium

// Time Complexity: O((m × n) ^ 2)
// Space Complexity: O(m × n)
func setZeroes(matrix [][]int)  {
    zeroes := [][]int{}

    for i := 0; i < len(matrix); i++ {         // rows
        for j := 0; j < len(matrix[i]); j++ {  // columns
            if matrix[i][j] == 0 {
                    zeroes = append(zeroes, []int{i, j})
                }
            }
        }
    
    for _, coord := range zeroes {
        i, j := coord[0], coord[1]

        zeroRow := make([]int, len(matrix[i]))
        matrix[i] = zeroRow  // rows

        for k := range matrix {
            matrix[k][j] = 0 // columns
        }
    }
}

// Time complexity: O(m*n)
// Space complexity: O(1) - no extra space, use the input array itself for tracking
func setZeroes2(matrix [][]int)  {
    zeroCol, zeroRow := false, false
    m, n := len(matrix), len(matrix[0])

    // First define if there zeroes at first row and column.
    // We need this to don't mess later with other cells.
    for i := 0; i < m; i++ {
        if matrix[i][0] == 0 {
            zeroCol = true
            break
        }
    }
    for j := 0; j < n; j++ {
        if matrix[0][j] == 0 {
            zeroRow = true
            break
        }
    }

    // Check for zeroes the rest of matrix using first row and column as
    // a tracking storage.
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][j] == 0 {
                matrix[i][0] = 0
                matrix[0][j] = 0
            }
        }
    }

    // Then using 'tracking storage' fill relevant rows/columns with zeroes
    for i := 1; i < m; i++ {
        for j := 1; j < n; j++ {
            if matrix[i][0] == 0 || matrix[0][j] == 0 {
                matrix[i][j] = 0
            }
        }
    }

    // Finally, process first row/column itself separately
    if zeroRow {
        for j := 0; j < n; j++ {
            matrix[0][j] = 0
        }
    }
    if zeroCol {
        for i := 0; i < m; i++ {
            matrix[i][0] = 0
        }
    }
}
