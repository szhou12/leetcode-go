package leetcode

import (
	"github.com/szhou12/leetcode-go/structures"
)

// Definition for a binary tree node.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base cases
	// case 1
	if root == nil {
		return nil
	}

	// case 2 + 3
	if root == p || root == q {
		return root
	}

	// LCA of left subtree
	llca := lowestCommonAncestor(root.Left, p, q)
	// LCA of right subtree
	rlca := lowestCommonAncestor(root.Right, p, q)

	if llca != nil && rlca != nil {
		return root
	}

	if llca != nil {
		return llca
	} else {
		return rlca
	}

}
