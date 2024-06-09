package leetcode

import "math"

// select a subarray nums[l..r] such that |k - (nums[l] OR nums[l + 1] ... OR nums[r])| is minimum
func minimumDifference(nums []int, k int) int {
	n := len(nums)
	root := NewSegTreeNodeORFromSlice(0, n-1, nums)

	res := math.MaxInt
	// 固定左边界i, 二分法搜右边界
	for i := 0; i < n; i++ {
		// left = lower limit of 右边界, right = upper limit of 右边界
		left, right := i, n-1
		for left < right {
			mid := right - (right-left)/2
			if root.queryRangeOR(i, mid) > k { // 当前bitwise OR i到mid 已经超过k, 不能再往右扩展了, 因为bitwise OR越多只会增大结果
				right = mid - 1
			} else {
				left = mid
			}
		}

		// 收敛后的结果nums[i:left]是最接近k且 <= k的右边界
		res1 := abs(root.queryRangeOR(i, left) - k)
		// 那么, nums[i:left+1]就是最接近k且 > k的右边界。如果left+1不越界，也是一个候选结果
		res2 := math.MaxInt
		if left+1 < n {
			res2 = abs(root.queryRangeOR(i, left+1) - k)
		}
		res = min(res, res1, res2)
	}

	return res
}

type SegTreeNodeOR struct {
	left, right *SegTreeNodeOR
	start, end int
	info int
	// delta int
	// tag bool
}

func NewSegTreeNodeORFromSlice(start, end int, vals []int) *SegTreeNodeOR {
	node := &SegTreeNodeOR {
		start: start,
		end: end,
	}
	if start == end {
		node.info = vals[start]
		return node
	}
	mid := (start + end) / 2
	node.left = NewSegTreeNodeORFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeORFromSlice(mid+1, end, vals)
	node.info = node.left.info | node.right.info
	return node
}

func (node *SegTreeNodeOR) queryRangeOR(start, end int) int {
	if end < node.start || node.end < start {
		// a | 0 = a
		return 0
	}
	if start <= node.start && node.end <= end {
		return node.info
	}
	// node.pushDown() // no range update in this problem, thus pushDown() is not needed
	return node.left.queryRangeOR(start, end) | node.right.queryRangeOR(start, end)
}
