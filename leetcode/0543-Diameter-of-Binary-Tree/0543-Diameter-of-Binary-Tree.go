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

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	_, diameter := diameterHelper(root)
	return diameter
}

func diameterHelper(root *TreeNode) (int, int) {
	// base case
	if root == nil {
		return 0, 0
	}

	leftDepth, leftDiameter := diameterHelper(root.Left)
	rightDepth, rightDiameter := diameterHelper(root.Right)
	curDepth := max(leftDepth, rightDepth) + 1
	// curDiameter := leftDepth + rightDepth

	depthSum := leftDepth + rightDepth
	diameterMax := max(leftDiameter, rightDiameter)
	curDiameter := max(depthSum, diameterMax)
	return curDepth, curDiameter

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
