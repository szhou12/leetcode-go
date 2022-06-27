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
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return checkSymmetric(root.Left, root.Right)

}

func checkSymmetric(p *TreeNode, q *TreeNode) bool {
	// base cases
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	// 注意这里的递归写法
	return checkSymmetric(p.Left, q.Right) && checkSymmetric(p.Right, q.Left)
}
