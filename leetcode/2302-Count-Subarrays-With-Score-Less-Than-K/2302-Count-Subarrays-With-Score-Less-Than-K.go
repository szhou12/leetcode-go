package leetcode

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	// prepend 0 for convenience
	nums = append([]int{0}, nums...)

	// prefix sum: 方便计算区间和
	presum := make([]int64, n+1)
	for i := 1; i < n+1; i++ {
		presum[i] = presum[i-1] + int64(nums[i])
	}

	res := int64(0)
	for i := 1; i < n+1; i++ {
		// 如果i-th element自己都超过k, 以它为end的所有subarray都不行，early stop
		if int64(nums[i]) >= k {
			continue
		}
		// WHY Binary Search?
		// 因为以 i-th element作结尾的 subarray, 越往左拓展，subarray sum是单调递增的, 适合使用binary search
		left, right := 0, i
		for left < right {
			// mid := left + (right - left) / 2 // will cause dead loop when [0, 1], 故选择中间靠右的mid
			mid := right - (right-left)/2
			if (presum[i]-presum[i-mid])*int64(mid) < k {
				left = mid // mid is a qualified candidate, thus, left = mid
			} else {
				right = mid - 1 // mid is NOT a qualified candidate, thus, right = mid - 1
			}
		}
		res += int64(left)
	}
	return res
}
