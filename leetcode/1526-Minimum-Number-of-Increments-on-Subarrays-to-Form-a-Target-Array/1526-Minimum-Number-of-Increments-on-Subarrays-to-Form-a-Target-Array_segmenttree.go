package leetcode

import "math"

func minNumberOperations_segtree(target []int) int {
	n := len(target)

	root := NewSegTreeNodeFromSlice(0, n-1, target) // O(n)

	// dfs(a, b, base) := min operations to form the range target[a:b] (inclusive) from the base value
	var dfs func(a int, b int, base int) int
	dfs = func(a int, b int, base int) int {
		// base cases
		if a > b {
			return 0 // this range has no meaning, thus return no meaning value 0
		}
		if a == b {
			return target[a] - base
		}

		// recursion
		temp := root.queryRange(a, b) // min value in range [a:b]
		val, idx := temp[0], temp[1]
		sum := val - base // we need this many more ops to reach the min element in [a:b]
		sum += dfs(a, idx-1, val)
		sum += dfs(idx+1, b, val)
		return sum
	}

	return dfs(0, n-1, 0) // O(nlogn)
}

type SegTreeNode struct {
	left  *SegTreeNode
	right *SegTreeNode
	start int
	end   int
	info  []int // [range min element's value, range min element's index]
}

func NewSegTreeNodeFromSlice(start int, end int, vals []int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
	}
	if start == end {
		node.info = []int{vals[start], start}
		return node
	}

	mid := start + (end-start)/2
	node.left = NewSegTreeNodeFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeFromSlice(mid+1, end, vals)

	if node.left.info[0] < node.right.info[0] {
		node.info = node.left.info
	} else {
		node.info = node.right.info
	}

	return node
}

func (node *SegTreeNode) queryRange(start int, end int) []int {
	// out of node's range
	if end < node.start || node.end < start {
		return []int{math.MaxInt, -1} // return no meaning value
	}
	// completely covers node's range
	if start <= node.start && node.end <= end {
		return node.info
	}

	l := node.left.queryRange(start, end)
	r := node.right.queryRange(start, end)

	if l[0] < r[0] {
		return l
	} else {
		return r
	}
}
