package leetcode

import "github.com/szhou12/leetcode-go/structures"

// Definition for a binary tree node.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base cases
	if root == nil || p == nil || q == nil {
		return nil
	}

	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
