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

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// Base cases
	if root1 == nil && root2 == nil {
		return root1
	}
	if root1 != nil && root2 == nil {
		return root1
	}
	if root1 == nil && root2 != nil {
		return root2
	}

	left := mergeTrees(root1.Left, root2.Left)
	right := mergeTrees(root1.Right, root2.Right)
	root1.Val += root2.Val
	root1.Left = left
	root1.Right = right
	return root1
}
