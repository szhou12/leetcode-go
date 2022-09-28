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

func buildTree(preorder []int, inorder []int) *TreeNode {
	// Step 1: Base case
	if len(preorder) == 0 {
		return nil
	}

	// Step 2: find root
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// Step 3: find splitting point
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}

	// Step 4: split inorder
	leftIn := inorder[:i]
	rightIn := inorder[i+1:]

	// Step 5: split preorder
	preorder = preorder[1:]
	leftPre := preorder[:len(leftIn)]
	rightPre := preorder[len(leftIn):]

	// Step 6: recursion left/right subtree
	left := buildTree(leftPre, leftIn)
	right := buildTree(rightPre, rightIn)

	// Step 7: connect
	root.Left = left
	root.Right = right

	return root
}

//  3
// / \
// 9 20
//   /\
// 15  7

// pr:[3] 9| 20 15 7
// i : 9 |3 15 20 7
// 	     3
//   /        \
//   9        20
// pr: 9    pr:[20] 15| 7
// i : 9    i : 15 |20 7
//             /    \
// 		      15     7
// 	        p: 15   p: 7
// 		    i: 15   i: 7
