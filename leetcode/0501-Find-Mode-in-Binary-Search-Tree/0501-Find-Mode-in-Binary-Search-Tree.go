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

// My Solution - naive
// Use a map to count each node value
type TreeNode = structures.TreeNode

func findMode(root *TreeNode) []int {
	countMap := make(map[int]int)

	modeHelper(root, &countMap)

	globalMax := math.MinInt
	var res []int
	for k, v := range countMap {
		if v > globalMax {
			globalMax = v
			res = []int{k}
		} else if v == globalMax {
			res = append(res, k)
		}
	}
	return res

}

func modeHelper(root *TreeNode, countMap *map[int]int) {
	if root == nil {
		return
	}

	(*countMap)[root.Val] += 1
	if root.Left != nil {
		modeHelper(root.Left, countMap)
	}
	if root.Right != nil {
		modeHelper(root.Right, countMap)
	}
}

// Better Solution - In-Order traversal
func findMode_InOrder(root *TreeNode) []int {
	curCount := 1
	maxCount := 1
	var prev *TreeNode
	var res []int

	var inOrder func(cur *TreeNode)
	inOrder = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		inOrder(cur.Left)
		if prev == nil {
			curCount = 1
			maxCount = 1
			res = append(res, cur.Val)
		} else {
			if cur.Val == prev.Val {
				curCount++
				if curCount == maxCount {
					res = append(res, cur.Val)
				} else if curCount > maxCount { // 找到了一个更大的众数, 抹掉之前所有append的
					maxCount = curCount
					res = []int{cur.Val}
				}
			}
			if cur.Val != prev.Val {
				curCount = 1
				if curCount == maxCount {
					res = append(res, cur.Val)
				}
			}
		}
		prev = cur
		inOrder(cur.Right)
	}

	inOrder(root)
	return res
}
