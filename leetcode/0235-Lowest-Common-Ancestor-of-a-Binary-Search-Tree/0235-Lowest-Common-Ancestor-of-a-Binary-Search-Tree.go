package leetcode

import "github.com/szhou12/leetcode-go/structures"

// Definition for a binary tree node.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base cases
	if root == nil || p == nil || q == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}
