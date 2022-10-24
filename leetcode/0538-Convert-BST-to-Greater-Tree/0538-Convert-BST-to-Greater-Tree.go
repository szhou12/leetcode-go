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

// Optimal Solution
// 中序遍历 - 先从右子树 ==> 当前层 ==> 最后左子树 --> 就能得到降序排列
// 当前层再累加sum就会简化代码很多
func convertBST(root *TreeNode) *TreeNode {
	sum := 0

	inOrder(root, &sum)
	return root
}

func inOrder(root *TreeNode, sum *int) {
	// Base case
	if root == nil {
		return
	}

	// 先走右子树
	inOrder(root.Right, sum)
	// 走完右子树后，此时更新后的sum就是要加在当前node的值
	*sum += root.Val
	root.Val = *sum
	// 再走左子树
	inOrder(root.Left, sum)

}

// My Solution
// 中序遍历 - 把所有node放入一个array里，并且升序排列，再从后往前依次累加Val
// 不高效，要是用一个 O(n) extra space
func convertBST_mysoln(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	bstArr := make([]*TreeNode, 0)
	convertToArray(root, &bstArr)

	for i := len(bstArr) - 2; i >= 0; i-- {
		bstArr[i].Val += bstArr[i+1].Val
	}
	return root
}

func convertToArray(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}
	convertToArray(root.Left, arr)
	*arr = append(*arr, root)
	convertToArray(root.Right, arr)
}
