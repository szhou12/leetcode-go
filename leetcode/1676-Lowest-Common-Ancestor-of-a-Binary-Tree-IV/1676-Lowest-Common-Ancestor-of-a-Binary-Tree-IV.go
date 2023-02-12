package leetcode

import "github.com/szhou12/leetcode-go/structures"

type TreeNode = structures.TreeNode

func lowestCommonAncestor(root *TreeNode, nodes []*TreeNode) *TreeNode {
	nodeSet := make(map[*TreeNode]bool)
	for _, node := range nodes {
		nodeSet[node] = true
	}

	res := LCA(root, nodeSet)
	return res
}

func LCA(root *TreeNode, set map[*TreeNode]bool) *TreeNode {
	// base cases
	if root == nil {
		return nil
	}
	if set[root] {
		return root
	}

	// recursion call
	left := LCA(root.Left, set)
	right := LCA(root.Right, set)

	// return current layer
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	} else {
		return right
	}
}
