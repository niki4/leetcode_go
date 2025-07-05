package medium

/*
Given the root of a binary tree, return the inorder traversal of its nodes' values.

// In-order traversal is to traverse the left subtree first.
// Then visit the root.
// Finally, traverse the right subtree.

For example, for the following tree result will be [A, B, C, D, E, F, G, H, I]
            F
         /     \
       B        G
     /   \       \
    A     D       I
         /  \     /
        C    E    H
*/

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2 MB, less than 31.18% of Go
// Time and Space complexity: O(n)
func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	var result []int
	node := root
	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		if len(stack) == 0 {
			return result
		}

		node = stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]

		node = node.Right
	}
}

// Runtime: 0 ms, faster than 100.00% of Go
// Memory Usage: 2.1 MB, less than 31.18% of Go
// Time and Space complexity: O(n)
func inorderTraversalRecursive(root *TreeNode) []int {
	res := new([]int)
	inorderHelper(res, root)
	return *res
}

func inorderHelper(res *[]int, node *TreeNode) {
	if node == nil {
		return
	}
	inorderHelper(res, node.Left)
	*res = append(*res, node.Val)
	inorderHelper(res, node.Right)
}
