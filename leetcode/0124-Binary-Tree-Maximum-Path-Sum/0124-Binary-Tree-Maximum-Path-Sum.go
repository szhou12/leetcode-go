package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	res := []int{math.MinInt}
	dfs(root, res)
	return res[0]
}

func dfs(root *TreeNode, res []int) int {
	// base case
	if root == nil {
		return 0
	}

	leftSum := dfs(root.Left, res)
	rightSum := dfs(root.Right, res)
	curMax := root.Val
	if leftSum > 0 {
		curMax += leftSum
	}
	if rightSum > 0 {
		curMax += rightSum
	}

	res[0] = max(res[0], curMax)

	if leftSum < 0 && rightSum < 0 {
		return root.Val
	}
	return max(leftSum, rightSum) + root.Val
}
