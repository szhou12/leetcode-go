package leetcode

func countOfPeaks(nums []int, queries [][]int) []int {
	n := len(nums)

	// peak[] marks each peak in nums[]
	// peak[i] = 1 if nums[i] is a peak
	// peak[i] = 0 otherwise
	peak := make([]int, n)
	for i := 1; i < n-1; i++ {
		if nums[i-1] < nums[i] && nums[i] > nums[i+1] {
			peak[i] = 1
		}
	}

	// Segment Tree (sum) on peak[]
	root := NewSegTreeNodeFromSlice(0, n-1, peak)

	res := make([]int, 0)
	for _, q := range queries {
		if q[0] == 1 { // query peaks in range [l, r]
			l, r := q[1], q[2]

			// no peaks if l==r
			if l == r {
				res = append(res, 0)
				continue
			}

			// exclude peaks on both ends l and r
			borderPeaks := peak[l] + peak[r] // this line determines that we need to keep updating peak[]
			res = append(res, root.queryRange(l, r)-borderPeaks)
		} else if q[0] == 2 { // update single peaks at j-1, j, j+1 respectively
			j, x := q[1], q[2]
			nums[j] = x // update j-th number in nums[]

			// update peak[j]
			if 0 <= j-1 && j+1 < n {
				if nums[j-1] < nums[j] && nums[j] > nums[j+1] { // a new peak is born at j
					root.updateSingle(j, 1)
					peak[j] = 1
				} else { // a peak is gone at j
					root.updateSingle(j, 0)
					peak[j] = 0
				}
			}

			// update peak[j-1]
			if 0 <= (j-1)-1 && (j-1)+1 < n {
				if nums[(j-1)-1] < nums[j-1] && nums[j-1] > nums[(j-1)+1] { // a new peak is born at j-1
					root.updateSingle(j-1, 1)
					peak[j-1] = 1
				} else { // a peak is gone at j-1
					root.updateSingle(j-1, 0)
					peak[j-1] = 0
				}
			}

			// update peak[j+1]
			if 0 <= (j+1)-1 && (j+1)+1 < n {
				if nums[(j+1)-1] < nums[j+1] && nums[j+1] > nums[(j+1)+1] { // a new peak is born at j+1
					root.updateSingle(j+1, 1)
					peak[j+1] = 1
				} else { // a peak is gone at j+1
					root.updateSingle(j+1, 0)
					peak[j+1] = 0
				}

			}
		}
	}

	return res
}

// Segment Tree (sum)
type SegTreeNode struct {
	left, right *SegTreeNode
	start, end  int
	info        int // sum over [start:end]
}

func NewSegTreeNodeFromSlice(start, end int, vals []int) *SegTreeNode {
	node := &SegTreeNode{
		start: start,
		end:   end,
	}

	// leaf node
	if start == end {
		node.info = vals[start]
		return node
	}

	mid := start + (end-start)/2
	node.left = NewSegTreeNodeFromSlice(start, mid, vals)
	node.right = NewSegTreeNodeFromSlice(mid+1, end, vals)
	node.info = node.left.info + node.right.info
	return node

}

func (node *SegTreeNode) updateSingle(index int, val int) {
	// out of node's range
	if index < node.start || node.end < index {
		return
	}
	// completely covers node's range
	if index <= node.start && node.end <= index {
		node.info = val
		return
	}
	node.left.updateSingle(index, val)
	node.right.updateSingle(index, val)
	node.info = node.left.info + node.right.info
}

func (node *SegTreeNode) queryRange(start, end int) int {
	// out of node's range
	if end < node.start || node.end < start {
		return 0
	}
	// completely covers node's range
	if start <= node.start && node.end <= end {
		return node.info
	}
	// No need for pushDown()
	return node.left.queryRange(start, end) + node.right.queryRange(start, end)
}
