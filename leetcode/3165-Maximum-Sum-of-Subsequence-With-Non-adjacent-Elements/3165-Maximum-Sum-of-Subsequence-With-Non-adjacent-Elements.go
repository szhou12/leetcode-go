package leetcode

import "math"

var M = int(1e9 + 7)

func maximumSumSubsequence(nums []int, queries [][]int) int {
	root := NewSegTreeNodeFromSlice(0, len(nums)-1, nums)

	res := 0
	for _, q := range queries {
		root.updateSingle(q[0], q[1])
		res += max(root.info00, root.info01, root.info10, root.info11)
		res %= M
	}
	return res
}

type SegTreeNode struct {
	left, right                    *SegTreeNode
	start, end                     int
	info00, info01, info10, info11 int
}

func NewSegTreeNodeFromSlice(start, end int, vals []int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
	}
	if start == end {
		node.info00 = 0
		node.info11 = vals[start]
		node.info01 = math.MinInt / 3 // math.MinInt/2 will cause overflow
		node.info10 = math.MinInt / 3
		return node
	}
	mid := (start + end) / 2
	node.left = NewSegTreeNodeFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeFromSlice(mid+1, end, vals)
	node.info00 = max(node.left.info00+node.right.info00, node.left.info01+node.right.info00, node.left.info00+node.right.info10)
	node.info01 = max(node.left.info00+node.right.info01, node.left.info01+node.right.info01, node.left.info00+node.right.info11)
	node.info10 = max(node.left.info10+node.right.info00, node.left.info11+node.right.info00, node.left.info10+node.right.info10)
	node.info11 = max(node.left.info10+node.right.info01, node.left.info11+node.right.info01, node.left.info10+node.right.info11)
	return node
}

func (node *SegTreeNode) updateSingle(index int, val int) {
	if index < node.start || node.end < index {
		return
	}
	if index <= node.start && node.end <= index {
		node.info11 = val
		node.info00 = 0
		return
	}
	node.left.updateSingle(index, val)
	node.right.updateSingle(index, val)
	node.info00 = max(node.left.info00+node.right.info00, node.left.info01+node.right.info00, node.left.info00+node.right.info10)
	node.info01 = max(node.left.info00+node.right.info01, node.left.info01+node.right.info01, node.left.info00+node.right.info11)
	node.info10 = max(node.left.info10+node.right.info00, node.left.info11+node.right.info00, node.left.info10+node.right.info10)
	node.info11 = max(node.left.info10+node.right.info01, node.left.info11+node.right.info01, node.left.info10+node.right.info11)
}
