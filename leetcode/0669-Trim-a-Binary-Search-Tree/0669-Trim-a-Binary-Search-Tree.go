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

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	left := trimBST(root.Left, low, high)
	right := trimBST(root.Right, low, high)

	if low <= root.Val && root.Val <= high {
		root.Left = left
		root.Right = right
		return root
	} else {
		// 如果root不在合法范围, 根据BST的性质, 则一定有一边的子树会跟着被去掉 (ie. 一定有一边返回的是nil)
		// Case 1: root 小于 范围, 则root的左子树一定会被砍掉
		// Case 2: root 超过 范围，则root的右子树一定会被砍掉
		// 所以返回非空的那一边子树就行了
		if left != nil {
			return left
		} else {
			return right
		}
	}

}
