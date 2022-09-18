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

func isBalanced(root *TreeNode) bool {
	balanced := true
	maxHeight(root, &balanced) // need to pass balanced as pointer
	return balanced
}

func maxHeight(root *TreeNode, balanced *bool) int {
	// base case
	if root == nil {
		return 0
	}

	// early stopping: if balanced already false, no need to proceed
	if !(*balanced) {
		return -1
	}

	leftH := maxHeight(root.Left, balanced)
	rightH := maxHeight(root.Right, balanced)

	if abs(leftH-rightH) > 1 {
		*balanced = false
	}

	return max(leftH, rightH) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}
