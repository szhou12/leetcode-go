package template

type SegTreeNode struct {
	left, right *SegTreeNode
	start, end int
	info int // sum over the range
	delta int
	tag bool
}

func NewSegTreeNode(start, end int, val int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end: end,
		info: val,
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

func (node *SegTreeNode) updateRangeBy(start int, end int, val int) {
	if end < node.start || node.end  < start {
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