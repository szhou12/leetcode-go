package leetcode

import (
	"math"

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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return checkValid(root, math.MinInt, math.MaxInt)
}

func checkValid(root *TreeNode, min int, max int) bool {
	// base case
	if root == nil {
		return true
	}

	// 若 root.Val 不符合 max 和 min 的限制，说明不是合法 BST
	// 注意！！！不能只比较 root.Val 和 root.Left.Val/root.Right.Val
	// 因为 root.Val需要比左/右子树所有的node的值都大/都小
	if root.Val <= min || root.Val >= max {
		return false
	}

	// 限定左子树的最大值是 root.Val，右子树的最小值是 root.Val
	return checkValid(root.Left, min, root.Val) && checkValid(root.Right, root.Val, max)
}
