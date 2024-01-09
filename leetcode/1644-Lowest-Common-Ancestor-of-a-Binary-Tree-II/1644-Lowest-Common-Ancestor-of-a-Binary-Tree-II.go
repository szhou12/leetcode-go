package leetcode

import "github.com/szhou12/leetcode-go/structures"

type TreeNode = structures.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res, count := helper(root, p, q)

	if count < 2 {
		return nil
	}

	return res
}

func helper(root, p, q *TreeNode) (*TreeNode, int) {
	// base case
	if root == nil {
		return nil, 0
	}

	// recursion call
	llca, lcount := helper(root.Left, p, q)
	rlca, rcount := helper(root.Right, p, q)

	// return
	if root == p || root == q {
		return root, 1 + lcount + rcount
	}

	if llca != nil && rlca != nil {
		return root, 2
	}

	if llca != nil {
		return llca, lcount
	} else {
		return rlca, rcount
	}
}
