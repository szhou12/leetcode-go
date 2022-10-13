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

func sumOfLeftLeaves(root *TreeNode) int {
	return findLeft(root)
}

func findLeft(root *TreeNode) int {
	// Base Case
	if root == nil {
		return 0
	}
	// Note: base case写在leaf node 也可以
	// if root.Left == nil && root.Right == nil {
	// 	return 0
	// }

	leftVal := findLeft(root.Left)
	rightVal := findLeft(root.Right)

	// 当前层: 运用left leaf的定义来写当前层要干啥
	addVal := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		addVal = root.Left.Val
	}

	return addVal + leftVal + rightVal
}
