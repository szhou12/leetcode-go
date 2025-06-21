package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	res := -1
	count := 0

	var dfs func(root *TreeNode, k int)
	dfs = func(root *TreeNode, k int) {
		// base case
		if root == nil {
			return
		}

		// in-order traversal
		dfs(root.Left, k)
		count++
		if count == k {
			res = root.Val
			return
		}
		dfs(root.Right, k)
	}

	dfs(root, k)

	return res
}