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
	maxDiameter := 0 // Note: it counts number of edges, NOT nodes

	// semantic: dfs returns the max height (# of edges) of input node
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		// base case
		if root == nil {
			return 0
		}

		// recursion
		left := dfs(root.Left)
		right := dfs(root.Right)

		// update maxDiameter
		maxDiameter = max(maxDiameter, left + right) // Note: here we only have left + right because we are counting edges

		return max(left, right) + 1
	}


	dfs(root)

	return maxDiameter
}


/**
 * Extension:
 * Given a binary tree, find the longest path that contains an even number of nodes.
 */
func longestPathWithEvenNodes(root *TreeNode) int {
	maxLength := 0

	// semantic: dfs returns the max height (# of nodes) of input node
	var dfs func(root *TreeNode) int

	dfs = func(root *TreeNode) int {
		// base case
		if root == nil {
			return 0
		}

		// recursion
		left := dfs(root.Left)
		right := dfs(root.Right)

		// update maxLength
		curPath := left + right + 1 // + 1 because we are counting nodes
		if curPath % 2 == 0 {
			maxLength = max(maxLength, curPath)
		}

		return max(left, right) + 1
	}

	dfs(root)

	return maxLength
}
