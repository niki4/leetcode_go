/*
Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure
and node values of subRoot and false otherwise.

A subtree of a binary tree `tree` is a tree that consists of a node in tree and all of this node's descendants.
The tree `tree` could also be considered as a subtree of itself.

Example 1:
            3 (root)
           /\
         4   5             4 (subroot)
        / \               / \
       1   2             1   2

Input: root = [3,4,5,1,2], subRoot = [4,1,2]
Output: true

Example 2:
            3 (root)
           /\
         4   5             4 (subroot)
        / \               / \
       1   2             1   2
          /
         0
Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
Output: false

Constraints:
    The number of nodes in the root tree is in the range [1, 2000].
    The number of nodes in the subRoot tree is in the range [1, 1000].
    -104 <= root.val <= 104
    -104 <= subRoot.val <= 104
*/

package easy

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// `root` tree can be subtree of itself, but not vice versa
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}

	if sameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func sameTree(a, b *TreeNode) bool {
	// base case for recursion
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil && a.Val == b.Val {
		return (sameTree(a.Left, b.Left) &&
			sameTree(a.Right, b.Right))
	}

	return false
}
