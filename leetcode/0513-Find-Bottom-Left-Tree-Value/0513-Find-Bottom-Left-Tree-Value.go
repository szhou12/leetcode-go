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

// My Solution - BFS: Level Order Traversal
func findBottomLeftValue_BFS(root *TreeNode) int {
	// level order traversal
	queue := []*TreeNode{root}

	var res int
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			if i == 0 {
				res = cur.Val
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
	return res

}

// Another Solution - DFS
func findBottomLeftValue_DFS(root *TreeNode) int {
	// Edge case
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	maxHeight := -1
	res := 0
	dfs(root, 0, &maxHeight, &res)
	return res
}

func dfs(root *TreeNode, curDepth int, maxHeight *int, res *int) {
	// Base case
	if root.Left == nil && root.Right == nil { // 到达叶子节点
		if curDepth > *maxHeight { // 保证这一层是第一次递归到达，也就保证此时这个节点一定是最底一层最左边的节点
			*maxHeight = curDepth
			*res = root.Val
		}
	}

	if root.Left != nil {
		dfs(root.Left, curDepth+1, maxHeight, res)
	}
	if root.Right != nil {
		dfs(root.Right, curDepth+1, maxHeight, res)
	}
}
