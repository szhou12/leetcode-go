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

func sortedArrayToBST(nums []int) *TreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, left int, right int) *TreeNode {
	// base case
	if left > right {
		return nil
	}

	// KEY: 构造根节点 - BST 节点左小右大，中间的元素就是根节点
	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}

	(*root).Left = buildTree(nums, left, mid-1)
	(*root).Right = buildTree(nums, mid+1, right)

	return root
}
