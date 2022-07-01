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
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base cases - at least one of input node reaches the bottom
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	// recursion
	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	if left && right && p.Val == q.Val {
		return true
	}

	return false
}
