package template

/*
Example: Segment Tree For Range Sum
*/

type SegTreeNode struct {
	left  *SegTreeNode
	right *SegTreeNode
	start int
	end   int
	info  int // customize: stored value over the range [start : end].
}

/*
Constructor: 所有叶子节点的info初始化为单一值val
O(logn)
*/
func NewSegTreeNode(start int, end int, val int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
		info:  val,
	}
	if start == end {
		return node
	}

	mid := start + (end-start)/2

	node.left = NewSegTreeNode(start, mid, val)
	node.right = NewSegTreeNode(mid+1, end, val)

	node.info = node.left.info + node.right.info // customize: range sum in [node.start : node.end]

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

	return node.left.queryRange(start, end) + node.right.queryRange(start, end) // customize: range sum in [start : end]
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

	node.info = node.left.info + node.right.info // customize: range sum in [node.start : node.end]
}