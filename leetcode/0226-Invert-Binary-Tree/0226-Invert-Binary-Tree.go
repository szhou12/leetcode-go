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

func invertTree(root *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return root
	}

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
