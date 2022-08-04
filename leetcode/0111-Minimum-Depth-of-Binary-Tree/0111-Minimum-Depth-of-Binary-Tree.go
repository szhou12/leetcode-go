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

func minDepth(root *TreeNode) int {
	// corner case
	if root == nil {
		return 0
	}
	// root 本身就是一层，depth 初始化为 1
	depth := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// if cur finds a leaf node, return depth
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		depth++
	}

	return depth
}
