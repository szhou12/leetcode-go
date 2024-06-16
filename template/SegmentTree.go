package template

type SegTreeNode struct {
	left, right *SegTreeNode
	start, end  int
	info        int  // stored value over the range [start, end]. e.g. stored value can be sum([a:b]), max([a:b]), min([a:b]), etc. The template given here is to sum over the range.
	delta       int  // Used for lazy propagation, it stores the value that needs to be added to all elements in this node's range
	tag         bool // A flag to indicate whether this node has a pending update (delta) that needs to be propagated to its children
}

/*
Constructor: 所有叶子节点的info初始化为单一值val
Create a new segment tree node and recursively build the entire segment tree
O(logn)
*/
func NewSegTreeNode(start, end int, val int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
		info:  val,
	}
	if start == end {
		return node
	}

	mid := (start + end) / 2
	node.left = NewSegTreeNode(start, mid, val)
	node.right = NewSegTreeNode(mid+1, end, val)
	node.info = node.left.info + node.right.info // range sum in [a : b]
	return node
}

/*
Constructor: i-th叶子节点的info初始化为vals[i]
O(logn)
*/
func NewSegTreeNodeFromSlice(start, end int, vals []int) *SegTreeNode {
	node := &SegTreeNode {
		start: start,
		end: end,
	}
	if start == end {
		node.info = vals[start]
		return node
	}
	mid := (start + end) / 2
	node.left = NewSegTreeNodeFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeFromSlice(mid+1, end, vals)

	// Update cur node's info: combine left and right child's info
	node.info = node.left.info + node.right.info // range sum in [start : end]
	return node

}

/*
单点修改 O(logn)
This method updates i-th leaf node by replacing with the val.
*/
func  (node *SegTreeNode) updateSingle(index int, val int) {
	if index < node.start || node.end < index {
		return
	}
	if index <= node.start && node.end <= index { // at leaf node
		node.info = val
		return
	}
	node.left.updateSingle(index, val)
	node.right.updateSingle(index, val)

	// Update cur node's info: merge left and right child's info
	node.info = node.left.info + node.right.info
}

/*
区间修改 O(logn)
This method updates a range [a, b] in the segment tree by adding val to each element in the range.
 1. If the current node's range is completely outside [a, b], it does nothing.
 2. If the current node's range is completely inside [a, b], it updates info and sets delta and tag for lazy propagation.
 3. If the current node's range partially overlaps with [a, b], it first propagates any pending updates to its children (using pushDown()),
    then recursively updates its children, and finally updates its info based on its children's info.
*/
func (node *SegTreeNode) updateRange(start int, end int, val int) {
	if end < node.start || node.end < start {
		return
	}
	if start <= node.start && node.end <= end {
		node.info += val * (node.end - node.start + 1)
		node.delta += val
		node.tag = true
		return
	}
	node.pushDown()
	node.left.updateRange(start, end, val)
	node.right.updateRange(start, end, val)

	// Update cur node's info: merge left and right child's info
	node.info = node.left.info + node.right.info
}

/*
区间查询
This method queries the sum of elements in the range [a, b].
1. If the current node's range is completely outside [a, b], it returns 0
2. If the current node's range is completely inside [a, b], it returns the node's info.
3. If the current node's range partially overlaps with [a, b], it first propagates any pending updates to its children (using pushDown), 
	then recursively queries its children and sums the results.
*/
func (node *SegTreeNode) queryRange(start, end int) int {
	if end < node.start || node.end < start {
		return 0
	}
	if start <= node.start && node.end <= end {
		return node.info
	}
	node.pushDown() // Needed only when we have updateRange()
	return node.left.queryRange(start, end) + node.right.queryRange(start, end)
}

/*
更新 <lazy标记>
This method is used for lazy propagation. When a node has an unpropagated update (tag==true), this method applies the update to its children.
1. The info of each child is increased by delta times the size of the range it covers.
2. The delta is added to each child's delta.
3. The tag of each child is set to true, indicating they now have an unpropagated update.
4. The current node's delta is reset to 0, and its tag is set to false, indicating the update has been propagated.

NOTE: pushDown() is ONLY needed when updateRange() needs implementation. In other words, if only updateSingle() and queryRange() need implementation, pushDown() is not necessary.
*/
func (node *SegTreeNode) pushDown() {
	if node.tag && node.left != nil {
		node.left.info += node.delta * (node.left.end - node.left.start + 1)
		node.left.delta += node.delta
		node.left.tag = true

		node.right.info += node.delta * (node.right.end - node.right.start + 1)
		node.right.delta += node.delta
		node.right.tag = true

		node.tag = false
		node.delta = 0
	}
}

/*
func main() {
	root := NewSegTreeNodeFromSlice(0, len(nums)-1, nums)
}
*/