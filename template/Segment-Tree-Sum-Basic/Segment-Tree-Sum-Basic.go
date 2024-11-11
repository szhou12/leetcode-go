package template

/*
Example: Segment Tree For Range Sum
*/

type SegTreeNode struct {
	left  *SegTreeNode
	right *SegTreeNode
	start int
	end   int
	info  int // TODO: stored value over the range [start : end].
}

/*
Constructor: 所有叶子节点的info初始化为单一值val
O(n)
*/
func NewSegTreeNode(start int, end int, val int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
	}
	if start == end {
		node.info = val
		return node
	}

	mid := start + (end-start)/2
	node.left = NewSegTreeNode(start, mid, val)
	node.right = NewSegTreeNode(mid+1, end, val)

	// TODO: here is range sum in [node.start : node.end]
	node.info = node.left.info + node.right.info

	return node
}

/*
Constructor: i-th叶子节点的info初始化为vals[i]
O(n)
*/
func NewSegTreeNodeFromSlice(start int, end int, vals []int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
	}
	if start == end {
		node.info = vals[start]
		return node
	}

	mid := start + (end-start)/2
	node.left = NewSegTreeNodeFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeFromSlice(mid+1, end, vals)

	// TODO: here is range sum in [node.start : node.end]
	node.info = node.left.info + node.right.info

	return node
}

/*
Range Query: 区间查询
O(logn)
*/
func (node *SegTreeNode) queryRange(start int, end int) int {
	// out of node's range -> the current node has no contribution to query range [start:end], return no meaning
	if end < node.start || node.end < start {
		return 0 // customize: return 0 for sum bc 0 has no meaning/effect in sum
	}
	// completely covers node's range -> we need all this node's info
	if start <= node.start && node.end <= end {
		return node.info
	}

	// TODO: here is range sum in [start : end]
	res := node.left.queryRange(start, end) + node.right.queryRange(start, end)

	return res
}

/*
Single Update: 单点修改
O(logn)
*/
func (node *SegTreeNode) updateSingle(index int, val int) {
	// out of node's range
	if index < node.start || node.end < index {
		return
	}
	// arrive at leaf node
	if node.start == node.end {
		node.info = val
		return
	}

	node.left.updateSingle(index, val)
	node.right.updateSingle(index, val)

	// TODO: here is range sum in [node.start : node.end]
	node.info = node.left.info + node.right.info
}
