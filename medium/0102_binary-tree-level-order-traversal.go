/*
Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).

Example 1:
		     3
		   /   \
		  9	    20
		 	   /  \
			  15   7
Input: root = [3,9,20,null,null,15,7]
Output: [[3],[9,20],[15,7]]

Example 2:
Input: root = [1]
Output: [[1]]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 2000].
-1000 <= Node.val <= 1000
*/

package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	nodes := []*TreeNode{root} // init level

	// Do BFS for each level, taking nodes from current level
	// from left and adding next level nodes to right of the `queue`
	for len(nodes) > 0 {
		levelSize := len(nodes)
		levelValues := []int{}

		for i := 0; i < levelSize; i++ {
			node := nodes[0]
			nodes = nodes[1:]

			if node != nil {
				levelValues = append(levelValues, node.Val)
				nodes = append(nodes, node.Left, node.Right)
			}
		}

		if len(levelValues) > 0 {
			result = append(result, levelValues)
		}
	}

	return result
}
