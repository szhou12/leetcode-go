package leetcode

import "github.com/szhou12/leetcode-go/structures"

type TreeNode = structures.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// edge case
	if root == nil {
		return make([][]int, 0)
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	level := 0

	queue = append(queue, root)
	qlen := len(queue)
	for qlen > 0 {
		curRes := make([]int, 0)
		for i := 0; i < qlen; i++ {
			if level == 0 {
				// pop from head, push from end (Left node first, then Right node)
				node := queue[0]
				queue = queue[1:]
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}

				curRes = append(curRes, node.Val)
			} else {
				// pop from end, push from head (Right node first, then Left node)
				node := queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				if node.Right != nil {
					queue = append([]*TreeNode{node.Right}, queue...)
				}
				if node.Left != nil {
					queue = append([]*TreeNode{node.Left}, queue...)
				}

				curRes = append(curRes, node.Val)
			}
		}
		level = 1 - level // flip level
		qlen = len(queue) // renew qlen
		res = append(res, curRes)
	}
	return res
}
