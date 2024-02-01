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


