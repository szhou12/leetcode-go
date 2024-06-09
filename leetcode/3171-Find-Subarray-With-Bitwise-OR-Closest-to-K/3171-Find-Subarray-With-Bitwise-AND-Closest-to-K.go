package leetcode

import "math"

// select a subarray nums[l..r] such that |k - (nums[l] AND nums[l + 1] ... AND nums[r])| is minimum
func minimumDifference_AND(nums []int, k int) int {
	n := len(nums)
	root := NewSegTreeNodeANDFromSlice(0, n-1, nums)

	res := math.MaxInt
	for i := 0; i < n; i++ {
		left, right := i, n-1 // nums[i...x]
		for left < right {
			mid := right - (right-left)/2
			if root.queryRangeAND(i, mid) < k {
				right = mid - 1
			} else {
				left = mid
			}
		}

		res1 := abs(root.queryRangeAND(i, left) - k)
		res2 := math.MaxInt
		if left+1 < n {
			res2 = abs(root.queryRangeAND(i, left+1) - k)
		}
		res = min(res, res1, res2)
	}

	return res

}

type SegTreeNodeAND struct {
	left, right *SegTreeNodeAND
	start, end  int
	info        int
	// delta       int
	// tag         bool
}

func NewSegTreeNodeANDFromSlice(start, end int, vals []int) *SegTreeNodeAND {
	node := &SegTreeNodeAND{
		start: start,
		end:   end,
	}
	if start == end {
		node.info = vals[start]
		return node
	}
	mid := (start + end) / 2
	node.left = NewSegTreeNodeANDFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeANDFromSlice(mid+1, end, vals)

	node.info = node.left.info & node.right.info
	return node
}

func (node *SegTreeNodeAND) queryRangeAND(start, end int) int {
	if end < node.start || node.end < start {
		// a & math.MaxInt = a
		return math.MaxInt
	}
	if start <= node.start && node.end <= end {
		return node.info
	}
	// node.pushDown() // no range update in this problem, thus pushDown() is not needed
	return node.left.queryRangeAND(start, end) & node.right.queryRangeAND(start, end)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
