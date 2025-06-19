package amz

/*
Given the roots of two BSTs, find the greatest common integer present in both trees. 
If there is no common integer, return -1.
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func FindGreatestCommonBST(root1 *TreeNode, root2 *TreeNode) int {
	res := -1

	set := make(map[int]bool)

	var dfs1 func(root *TreeNode)
	dfs1 = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}

		// in-order traversal
		dfs1(root.Left)
		set[root.Val] = true
		dfs1(root.Right)
	}

	var dfs2 func(root *TreeNode)
	dfs2 = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}

		// in-order traversal
		dfs2(root.Left)
		if _, ok := set[root.Val]; ok {
			res = max(res, root.Val)
		}
		dfs2(root.Right)
	}

	dfs1(root1)
	dfs2(root2)

	return res
}