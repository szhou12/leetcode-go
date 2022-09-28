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

func generateTrees(n int) []*TreeNode {
	// edge case
	if n == 0 {
		return []*TreeNode{}
	}

	return buildBST(1, n)
}

/* 构造闭区间 [min, max] 组成的 BST */
func buildBST(min int, max int) []*TreeNode {
	trees := make([]*TreeNode, 0)
	// Base Case
	if min > max {
		trees = append(trees, nil)
	}

	// Recursion
	// 1、穷举 root 节点的所有可能。
	for i := min; i <= max; i++ {
		// 2、递归构造出左、右子树的所有合法 BST。
		leftTree := buildBST(min, i-1)
		rightTree := buildBST(i+1, max)
		// 3、给 root 节点穷举所有左、右子树的组合。
		for _, l := range leftTree { // leftTree = [第一种左子树的根节点, 第二种左子树的根节点, ... ]
			for _, r := range rightTree {
				// i 作为 root
				root := &TreeNode{Val: i, Left: l, Right: r}
				trees = append(trees, root)
			}
		}
	}

	return trees
}
