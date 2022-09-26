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

func levelOrder(root *TreeNode) [][]int {
	// Edge Case
	if root == nil {
		return [][]int{}
	}

	var res [][]int

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		var level []int

		size := len(queue)
		for i := 0; i < size; i++ {
			// Pop a node out
			curNode := queue[0]
			queue = queue[1:]
			// Add current node's value to current level
			level = append(level, curNode.Val)

			// Push current node's children (if not nil) to queue
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		res = append(res, level)

	}

	return res
}
