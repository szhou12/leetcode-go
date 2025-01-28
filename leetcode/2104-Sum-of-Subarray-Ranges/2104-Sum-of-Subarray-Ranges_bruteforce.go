package leetcode

// Brute Force: find all subarrays and calculate each range
// NOTE: we CANNOT sort as the order affects results!
func subArrayRanges_bruteforce(nums []int) int64 {
	n := len(nums)

	res := 0

	// O(n^2) to find all subarrays
	for i := 0; i < n; i++ {
		mn, mx := nums[i], nums[i]

		for j := i; j < n; j++ {
			// update as subarray grows
			mn = min(mn, nums[j])
			mx = max(mx, nums[j])

			res += (mx - mn)
		}
	}

	return int64(res)
}