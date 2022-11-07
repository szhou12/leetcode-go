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

func minCameraCover(root *TreeNode) int {
	count := 0
	setCamera(root, false, &count)
	return count
}

// 状态:
// -1: 子节点为empty - 空节点
// 0: 子节点uncovered
// 1: 子节点covered
// 2: 子节点set camera
func setCamera(root *TreeNode, hasParent bool, count *int) int {
	// base case
	if root == nil {
		return -1
	}

	// 获取左右子节点的情况
	left := setCamera(root.Left, true, count)
	right := setCamera(root.Right, true, count)

	// 根据左右子节点的情况和父节点的情况判断当前节点应该做的事情

	// Piority Level I: uncovered
	// 左、右子节点只要有一个 uncovered, 另一个 uncovered/empty/covered/camera, 当前节点必须放置camera
	if left == 0 || right == 0 {
		*count++
		return 2
	}

	// Piority Level II: camera
	// 左、右子节点只要有一个 camera, 另一个 empty/covered/camera, 当前节点已经covered, 什么都不用管了
	if left == 2 || right == 2 {
		return 1
	}

	// Piority Level III: empty
	// 左、右子节点有一个 empty, 另一个 empty/covered, 当前节点操作取决于有没有父节点
	if left == -1 || right == -1 {
		if hasParent { // 有父节点的话，让父节点来 cover 自己
			return 0
		} else { // 没有父节点的话，自己 set camera
			*count++
			return 2
		}
	}

	// Piority Level IV: covered
	// 左、右子节点有一个 covered, 另一个 covered, 当前节点操作取决于有没有父节点
	if hasParent { // 有父节点的话，让父节点来 cover 自己
		return 0
	} else { // 没有父节点的话，自己 set camera
		*count++
		return 2
	}
}
