package leetcode

import "github.com/szhou12/leetcode-go/structures"

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 触底 leaf node's child
	if root == nil {
		return false
	}

	// left node
	if root.Val == targetSum && root.Left == nil && root.Right == nil {
		return true
	}

	left := hasPathSum(root.Left, targetSum-root.Val)
	right := hasPathSum(root.Right, targetSum-root.Val)

	return left || right
}
