package leetcode

import "github.com/szhou12/leetcode-go/structures"

// Definition for a binary tree node.
type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base cases

	/* 如果p, q不一定存在于树中，这么写 */
	// if root == nil || p == nil || q == nil {
	// 	return nil
	// }

	/* 因为题目说明p, q一定存在于树中，所以这么写 */
	if root == nil {
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
