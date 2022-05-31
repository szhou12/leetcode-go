package leetcode

func canJump(nums []int) bool {
	return greedyHelper(nums)
}

// Implements Greedy algorithm
func greedyHelper(nums []int) bool {
	cur := 0  // cur = farthest index I can currently jump to
	next := 0 // next = farthest index I can jump to in the next step
	for i := 0; i < len(nums); i++ {
		// i > cur means my current farthest jump cannot reach i, want to know if next will help
		if i > cur {
			if cur == next {
				// cur == next means the farthest index my next step can jump to is the same as I can currently jump to
				// i.e. the next step is not making any progress from my current step. Thus, I definitely can't reach end
				return false
			}
			cur = next
		}
		next = max(next, i+nums[i])
	}
	return true
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
