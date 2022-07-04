package leetcode

import (
	"github.com/szhou12/leetcode-go/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = structures.TreeNode

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, tilt := tiltHepler(root)
	return tilt

}

func tiltHepler(root *TreeNode) (int, int) {
	// base case
	if root == nil {
		return 0, 0
	}

	leftSum, leftTilt := tiltHepler(root.Left)
	rightSum, rightTilt := tiltHepler(root.Right)

	curSum := leftSum + rightSum + root.Val
	curTilt := abs(leftSum-rightSum) + leftTilt + rightTilt
	return curSum, curTilt

}

func abs(a int) int {
	if a < 0 {
		return a
	}
	return a
}
