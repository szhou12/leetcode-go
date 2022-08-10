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

func postorderTraversal(root *TreeNode) []int {
	var result []int

	postorder(root, &result)
	return result
}

func postorder(root *TreeNode, result *[]int) {
	// base case
	if root == nil {
		return
	}

	postorder(root.Left, result)
	postorder(root.Right, result)
	*result = append(*result, root.Val)
}
