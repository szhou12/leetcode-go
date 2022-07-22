package leetcode

import "github.com/szhou12/leetcode-go/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func inorderTraversal(root *TreeNode) []int {
	var result []int

	inOrder(root, &result)
	return result
}

func inOrder(root *TreeNode, result *[]int) {
	// base case
	if root == nil {
		return
	}

	inOrder(root.Left, result)
	*result = append(*result, root.Val)
	inOrder(root.Right, result)
}
