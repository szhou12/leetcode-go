package leetcode

var M = int(1e9 + 7)

func sumCounts(nums []int) int {
	n := len(nums)

	// STEP 1: for each i, find last seen position of nums[i]
	position := make(map[int]int) // {nums[i] : recently seen index}
	prev := make([]int, n)        // prev[i] = last seen position of nums[i]
	// initialize prev: MUST init to -1 bc 0 is a valid position
	for i := 0; i < n; i++ {
		prev[i] = -1
	}
	for i := 0; i < n; i++ {
		if k, ok := position[nums[i]]; ok {
			prev[i] = k
		}
		position[nums[i]] = i
	}

	// STEP 2: DP + Segment Tree
	// dp[i] := sum of squares of distinct counts of all subarrays in num[0:i] ending at i
	root := NewSegTreeNode(0, n-1, 0)
	dp := make([]int, n)
	// recurrence loop MUST start at i=0. No need to addtionally write a base case
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] += 0
		} else {
			dp[i] += dp[i-1]
		}
		k := prev[i]
		dp[i] += 2 * root.queryRange(k+1, i-1)
		dp[i] += i - 1 - k
		dp[i] += 1
		dp[i] %= M

		// update segment tree nums[i] contributes 1 count to range [k+1:i]
		root.updateRangeBy(k+1, i, 1)
	}

	// STEP 3: collect all subarrays ending at each i
	res := 0
	for i := 0; i < n; i++ {
		res += dp[i]
		res %= M
	}
	return res

}

// Segment Tree Implementation
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