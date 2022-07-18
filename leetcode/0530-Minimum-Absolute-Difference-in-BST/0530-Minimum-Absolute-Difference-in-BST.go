package leetcode

import (
	"math"

	"github.com/szhou12/leetcode-go/structures"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = structures.TreeNode

func getMinimumDifference(root *TreeNode) int {
	// max val in left subtree + cur node + min val in right subtree

	var l []int
	inOrder(root, &l)

	res := math.MaxInt
	for i := 1; i < len(l); i++ {
		res = min(res, l[i]-l[i-1])
	}

	return int(res)
}

func inOrder(root *TreeNode, treeList *[]int) {
	if root == nil {
		return
	}

	inOrder(root.Left, treeList)
	*treeList = append(*treeList, root.Val)
	inOrder(root.Right, treeList)

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
