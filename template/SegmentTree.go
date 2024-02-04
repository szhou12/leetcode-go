package template

type SegTreeNode struct {
	left, right *SegTreeNode
	start, end  int
	info        int  // sum over the range
	delta       int  // Used for lazy propagation, it stores the value that needs to be added to all elements in this node's range
	tag         bool // A flag to indicate whether this node has a pending update (delta) that needs to be propagated to its children
}

/*
Create a new segment tree node and recursively build the entire segment tree
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
	node.info = node.left.info + node.right.info // what is it doing?
	return node
}

/*
This method is used for lazy propagation. When a node has an unpropagated update (tag==true), this method applies the update to its children.
1. The info of each child is increased by delta times the size of the range it covers.
2. The delta is added to each child's delta.
3. The tag of each child is set to true, indicating they now have an unpropagated update.
4. The current node's delta is reset to 0, and its tag is set to false, indicating the update has been propagated.
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
This method updates a range [a, b] in the segment tree by adding val to each element in the range.
 1. If the current node's range is completely outside [a, b], it does nothing.
 2. If the current node's range is completely inside [a, b], it updates info and sets delta and tag for lazy propagation.
 3. If the current node's range partially overlaps with [a, b], it first propagates any pending updates to its children (using pushDown()),
    then recursively updates its children, and finally updates its info based on its children's info.
*/
func (node *SegTreeNode) updateRangeBy(start int, end int, val int) {
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
	node.left.updateRangeBy(start, end, val)
	node.right.updateRangeBy(start, end, val)
	node.info = node.left.info + node.right.info
}

/*
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
	node.pushDown()
	return node.left.queryRange(start, end) + node.right.queryRange(start, end)
}
