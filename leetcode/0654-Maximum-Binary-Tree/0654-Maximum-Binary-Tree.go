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

func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := buildTree(nums)
	return root
}

func buildTree(nums []int) *TreeNode {
	// base case
	if len(nums) < 1 {
		return nil
	}

	maxIdx := -1
	maxVal := -1
	for i, val := range nums {
		if val > maxVal {
			maxIdx = i
			maxVal = val
		}
	}

	left := buildTree(nums[:maxIdx])
	right := buildTree(nums[maxIdx+1:])
	root := &TreeNode{Val: maxVal, Left: nil, Right: nil}
	root.Left = left
	root.Right = right
	return root
}
