package leetcode

// leftMin[i] := min element in nums[0:i]
// rightMax[i] := max element in nums[i:n-1]

// leftMin[0] = nums[0]
// leftMin[i] = min(nums[i], leftMin[i-1])
// rightMax[n-1] = nums[n-1]
// rightMax[i] = max(nums[i], rightMax[i+1])
func increasingTriplet(nums []int) bool {
	n := len(nums)

	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for i := 1; i+1 < n; i++ {
		if leftMin[i] < nums[i] && nums[i] < rightMax[i+1] {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
