/*
Given the root of a binary tree, imagine yourself standing on the right side of it,
return the values of the nodes you can see ordered from top to bottom.

Example 1:
Input: root = [1,2,3,null,5,null,4]
Output: [1,3,4]
Explanation: https://assets.leetcode.com/uploads/2024/11/24/tmpd5jn43fs-1.png

Example 2:
Input: root = [1,2,3,4,null,null,null,5]
Output: [1,3,4,5]
Explanation: https://assets.leetcode.com/uploads/2024/11/24/tmpkpe40xeh-1.png

Example 3:
Input: root = [1,null,3]
Output: [1,3]

Example 4:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
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

// BFS (breadth-first-search) so that we traverse nodes level by level
func rightSideView(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	nodes := []*TreeNode{root}

	for len(nodes) > 0 {
		var rightMostLvlNode *TreeNode
		levelSize := len(nodes)

		// pop (from left in `nodes` queue) all nodes from current level
		// while adding nodes of the next level (from right)
		for i := 0; i < levelSize; i++ {
			node := nodes[0]
			nodes = nodes[1:]

			if node != nil {
				rightMostLvlNode = node
				nodes = append(nodes, node.Left, node.Right)
			}
		}

		if rightMostLvlNode != nil {
			result = append(result, rightMostLvlNode.Val)
		}
	}

	return result
}
