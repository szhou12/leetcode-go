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

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	res, _ := helper(root, 0)
	return res
}

// return: (LCA, depth)
// return 的物理意义: subtree中是否找到LCA, 走subtree可以到达的最深有多深
func helper(root *TreeNode, depth int) (*TreeNode, int) {
	// base case
	if root == nil {
		return nil, depth
	}

	// recursion call
	// 问问看: 是否能在左子树中找到LCA, 左子树最深可以到达哪里
	llca, leftDepth := helper(root.Left, depth+1)
	// 问问看: 是否能在右子树中找到LCA, 右子树最深可以到达哪里
	rlca, rightDepth := helper(root.Right, depth+1)

	// cur layer return

	// 如果左右子树都能走到的深度一样, 那么当前节点就是它们各自最深叶子节点的最近公共祖先
	if leftDepth == rightDepth {
		return root, leftDepth // 注意: 这里返回子树上报的depth而不是当前层的depth, 因为子树上报的depth是最深depth
	}

	// subtree哪个能走得更深就返回哪个
	if leftDepth > rightDepth {
		return llca, leftDepth
	} else {
		return rlca, rightDepth
	}

}
