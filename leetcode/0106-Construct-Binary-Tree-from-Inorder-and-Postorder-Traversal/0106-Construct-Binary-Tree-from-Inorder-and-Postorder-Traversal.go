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

func buildTree(inorder []int, postorder []int) *TreeNode {
	// Step 1: Base Case - if len(postorder) == 0: reach nil node
	if len(postorder) == 0 {
		return nil
	}

	// Step 2: Find root - last element of postorder
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// Step 3: find splitting point - inorder index of root
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVal {
			break
		}
	}

	// Step 4: split inorder by splitting point (exclude splitting point (ie. root) in sub-inorder)
	leftIn := inorder[:i]
	rightIn := inorder[i+1:]

	// Step 5: split postorder (after split, sub-postorder should have same size as sub-inorder)
	postorder = postorder[:len(postorder)-1] // exclude last element as root
	leftPost := postorder[:i]                // 也可以写成 leftPost := postorder[:len(leftIn)]
	rightPost := postorder[i:]               // 也可以写成 rightPost := postorder[len(leftIn):]

	// Step 6: Recursion to find left/right subtrees
	left := buildTree(leftIn, leftPost)
	right := buildTree(rightIn, rightPost)

	// Step 7: connect root to left/right
	root.Left = left
	root.Right = right

	return root
}

//  3
// / \
// 9 20
//   /\
// 15  7

// p: 9 |15 7 20 3
// i: 9 |3 15 20 7
// 	     3
//   /        \
//   9        20
// p: 9      p: 15 |7 20
// i: 9      i: 15 |20 7
//             /    \
// 		      15     7
// 	        p: 15   p: 7
// 		    i: 15   i: 7
