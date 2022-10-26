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

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0

	inOrder(root, &sum)
	return root
}

func inOrder(root *TreeNode, sum *int) {
	// base case
	if root == nil {
		return
	}

	inOrder(root.Right, sum)
	*sum += root.Val
	root.Val = *sum
	inOrder(root.Left, sum)
}
