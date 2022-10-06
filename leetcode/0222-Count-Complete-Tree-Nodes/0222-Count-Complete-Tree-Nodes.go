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

// 当作普通Binary Tree来做 - Post-order traversal
// Time complexity = O(n)
func countNodes_Post(root *TreeNode) int {
	// base case
	if root == nil {
		return 0
	}

	left := countNodes_Post(root.Left)
	right := countNodes_Post(root.Right)
	return 1 + left + right
}

// Optimal Solution: full binary tree
// Complete binary tree 一定会有一个 full binary subtree,
// 如果找到了它，就可以直接通过subtree height来计算节点总数
// 如果没找到，就正常递归
func countNodes(root *TreeNode) int {
	// Edge Case
	if root == nil {
		return 0
	}

	l, r := root, root
	lh, rh := 0, 0
	for l != nil {
		l = l.Left
		lh++
	}
	for r != nil {
		r = r.Right
		rh++
	}

	if lh == rh { // 如果左右子树的高度相同, find a full binary subtree: return 2^height - 1
		return int(math.Pow(2, float64(lh))) - 1
	}

	// 如果左右高度不同，则按照普通二叉树的逻辑计算
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
