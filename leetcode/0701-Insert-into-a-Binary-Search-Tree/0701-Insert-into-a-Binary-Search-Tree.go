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

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		newNode := &TreeNode{Val: val, Left: nil, Right: nil}
		return newNode
	}
	if root.Val > val {
		left := insertIntoBST(root.Left, val)
		root.Left = left
		return root
	} else {
		right := insertIntoBST(root.Right, val)
		root.Right = right
		return root
	}
}
