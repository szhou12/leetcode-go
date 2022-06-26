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

// Definition for a binary tree node.
type TreeNode = structures.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return 1 + max(left, right)
}

func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
