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

func deleteNode(root *TreeNode, key int) *TreeNode {
	// 节点为空，返回
	if root == nil {
		return root
	}

	// 往右子树寻找
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 往左子树寻找
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	// root.Val == key
	// Case 1: 当前节点的左子树为空，返回当前的右子树
	if root.Left == nil {
		return root.Right
	}
	// Case 2: 当前节点的右子树为空，返回当前的左子树
	if root.Right == nil {
		return root.Left
	}
	// Case 3: 左右子树都不为空，找到右子树的最左节点 记为node
	node := root.Right
	for node.Left != nil {
		node = node.Left
	}
	// 将当前节点的左子树挂在node的左孩子上
	// 原因: node目前是当前节点的右子树中最小值, node左孩子挂上当前节点的左子树依旧保持了BST性质
	node.Left = root.Left
	// 当前节点的右子树替换掉当前节点，完成当前节点的删除
	root = root.Right
	return root
}
