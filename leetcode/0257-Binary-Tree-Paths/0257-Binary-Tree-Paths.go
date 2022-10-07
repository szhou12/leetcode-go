package leetcode

import (
	"strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	var path []int
	var res []string

	traverse(root, path, &res)
	return res
}

func traverse(root *TreeNode, path []int, res *[]string) {
	path = append(path, root.Val)

	// Base Case: leaf node
	if root.Left == nil && root.Right == nil {
		sPath := ""
		for i := 0; i < len(path)-1; i++ {
			sPath += strconv.Itoa(path[i]) + "->"
		}
		sPath += strconv.Itoa(path[len(path)-1])
		*res = append(*res, sPath)
		return
	}

	// left branch
	if root.Left != nil {
		traverse(root.Left, path, res)
		// 回溯不需要写！！！否则会出错
	}

	// right branch
	if root.Right != nil {
		traverse(root.Right, path, res)
		// 回溯不需要写！！！否则会出错
	}
}
