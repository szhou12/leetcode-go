package leetcode

func checkArray(nums []int, k int) bool {
	if k == 1 {
		return true
	}

	n := len(nums)
	diff := make([]int, n) // 长度为 n 或者 n+1 都可以

	base := 0
	for i := 0; i < n-1; i++ {
		base += diff[i]

		delta := nums[i] - base
		if delta < 0 {
			return false
		}
		diff[i] += delta // 可以不用写, 因为不会再用到diff[i]
		if i+k < n {
			diff[i+k] -= delta
		}

		base += delta
	}

	if diff[n-1]+base == nums[n-1] {
		return true
	}
	return false
}
