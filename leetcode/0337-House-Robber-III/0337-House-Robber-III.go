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

func rob(root *TreeNode) int {
	res := robTree(root)
	return max(res[0], res[1])
}

// DP := 长度为2的数组，第一个表示不偷家的总价值，第二个表示偷家的总价值
func robTree(root *TreeNode) []int {
	// Base Case
	if root == nil {
		return []int{0, 0}
	}

	left := robTree(root.Left)
	right := robTree(root.Right)

	// 当前层
	// Case 1: 当前层决定偷家，左右孩子都不能偷, 别忘了加上当前层家的价值
	robHouse := left[0] + right[0] + root.Val
	// Case 2: 当前层决定不偷家，左右孩子可以偷可以不偷，选较大的
	robHouseNot := max(left[0], left[1]) + max(right[0], right[1])

	return []int{robHouseNot, robHouse}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
