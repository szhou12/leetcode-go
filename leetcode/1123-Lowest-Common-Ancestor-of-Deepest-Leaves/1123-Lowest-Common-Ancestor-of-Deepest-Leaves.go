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

func helper(root *TreeNode, depth int) (*TreeNode, int) {
	// base case
	if root == nil {
		return nil, depth
	}

	// recursion call
	llca, leftDepth := helper(root.Left, depth+1)
	rlca, rightDepth := helper(root.Right, depth)

	// cur layer return
	if leftDepth == rightDepth {
		return root, leftDepth
	}

	if leftDepth > rightDepth {
		return llca, leftDepth
	} else {
		return rlca, rightDepth
	}

}
