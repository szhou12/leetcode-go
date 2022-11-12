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

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	inOrder(root, &nums)

	// LeetCode 108
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = buildBST(nums, left, mid-1)
	root.Right = buildBST(nums, mid+1, right)
	return root

}

func inOrder(root *TreeNode, nums *[]int) {
	// base case
	if root == nil {
		return
	}

	inOrder(root.Left, nums)
	(*nums) = append((*nums), root.Val)
	inOrder(root.Right, nums)
}
