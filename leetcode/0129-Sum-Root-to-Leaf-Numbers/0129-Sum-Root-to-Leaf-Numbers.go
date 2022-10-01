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

func sumNumbers(root *TreeNode) int {
	sum := 0
	rootVal := root.Val
	dfs(root, rootVal, &sum)
	return sum
}

func dfs(root *TreeNode, topVal int, sum *int) {
	// Base Case: reached a leaf node
	if root.Left == nil && root.Right == nil {
		*sum += topVal
		return
	}

	// 2 branches: Left, Right
	if root.Left != nil {
		dfs(root, topVal*10+root.Left.Val, sum)
	}
	if root.Right != nil {
		dfs(root, topVal*10+root.Right.Val, sum)
	}
}
