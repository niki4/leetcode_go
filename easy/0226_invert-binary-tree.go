/*
Given the root of a binary tree, invert the tree, and return its root.

Example 1:  https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg
Input: root = [4,2,7,1,3,6,9]
Output: [4,7,2,9,6,3,1]

Example 2:  https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg
Input: root = [2,1,3]
Output: [2,3,1]

Example 3:
Input: root = []
Output: []

Constraints:
The number of nodes in the tree is in the range [0, 100].
-100 <= Node.val <= 100
*/

package easy

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	// Recursive approach. Time/Space complexity: O(n)
	// The idea is to swap left and right child of each node, then recursively do DFS for its children.

	if root == nil { // base case, where to stop recursion
		return nil
	}

	// swap nodes (so to invert them)
	root.Left, root.Right = root.Right, root.Left

	// continue recurion for each subtree (DFS)
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}
