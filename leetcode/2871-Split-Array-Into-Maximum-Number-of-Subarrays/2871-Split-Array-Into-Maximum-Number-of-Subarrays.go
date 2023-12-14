package leetcode

func maxSubarrays(nums []int) int {
	n := len(nums)

	// # of splits
	res:= 0

	for l := 0; l < n; l++ {
		r := l
		cur := nums[l]
		for r < n - 1 && cur != 0 {
			r++
			cur &= nums[r]
		}
		// 1. r == n - 1 && cur == 0
		// 2. r < n = 1 && cur == 0
		// 3. r == n - 1 && cur != 0
		if cur == 0 {
			res++
		}
		l = r
	}

	// it's possible no split were found (res=0), then 1 being the whole array
	return max(1, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}