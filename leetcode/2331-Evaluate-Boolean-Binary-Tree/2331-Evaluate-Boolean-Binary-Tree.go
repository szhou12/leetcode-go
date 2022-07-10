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

func evaluateTree(root *TreeNode) bool {
	//base case: Leaf node
	if root.Left == nil && root.Right == nil {
		if root.Val == 0 {
			return false
		} else {
			return true
		}
	}

	left := evaluateTree(root.Left)
	right := evaluateTree(root.Right)

	if root.Val == 2 { // OR
		return left || right
	}

	// AND
	return left && right
}
