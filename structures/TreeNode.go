package structures

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NULL for testing purpose
var NULL = -1 << 63

/*
Convert int[] to a tree, return address of root
Used to write test cases for 0236-Lowest-Common-Ancestor-of-a-Binary-Tree
*/
func Ints2TreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	// length = 1; capacity = 2n
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

/*
Build a Binary Search Tree from an int array. e.g. root = [4,2,5,1,3]

		  4
		 / \
		2   5
	   / \
	  1   3
*/
func buildBST(nums []int) *TreeNode {

	// Recursive function to build the BST
	var build func(root *TreeNode, val int) *TreeNode
	build = func(root *TreeNode, val int) *TreeNode {
		if root == nil {
			return &TreeNode{Val: val}
		}

		if val < root.Val {
			root.Left = build(root.Left, val)
		} else {
			root.Right = build(root.Right, val)
		}
		return root
	}

	// Call build()
	var root *TreeNode
	for _, val := range nums {
		root = build(root, val)
	}

	return root
}

/*
Pretty-print a tree structure horizontally. i.e., leftmost is root, rightmost are leaf nodes

Args:

	root: current node
	prefix: prefix string for current node
	fromLeft: true if the current node is arrived from its parent node's left branch;
			  false if from its parent node's right branch

Call this function:

	visualizeTree(root, "", false)
*/
func visualizeTree(node *TreeNode, prefix string, fromLeft bool) {
	if node == nil {
		return
	}

	// In-order traversal + right child first for a better visualization (top-down view)
	// Right branch
	if node.Right != nil {
		newPrefix := prefix
		newPrefix += map[bool]string{true: "│   ", false: "    "}[fromLeft]
		visualizeTree(node.Right, newPrefix, false)
	}

	// Print the current node's value
	nodeLine := prefix
	nodeLine += map[bool]string{true: "└── ", false: "┌── "}[fromLeft]
	nodeLine += strconv.Itoa(node.Val)
	fmt.Println(nodeLine) // Println()来实现换行. 因为总是先走右子树，所以总是右子树的recursion cal先完成并打印，所以越靠右的节点越是在古早行数中打印

	// Left branch
	if node.Left != nil {
		newPrefix := prefix
		newPrefix += map[bool]string{true: "    ", false: "│   "}[fromLeft]
		visualizeTree(node.Left, newPrefix, true)
	}
}
