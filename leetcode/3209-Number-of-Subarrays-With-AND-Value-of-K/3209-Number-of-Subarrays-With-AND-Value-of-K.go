package leetcode

func countSubarrays(nums []int, k int) int64 {
	n := len(nums)

	// {key=31 possibilities : value=counts}
	// prev = occurrences from previous round
	prev := make(map[int]int)

	res := 0

	for i := 0; i < n; i++ {
		// cur = occurrences for current round
		cur := make(map[int]int)
		cur[nums[i]] += 1 // nums[i] itself can be a subarray
		for key, cnt := range prev {
			cur[key&nums[i]] += cnt
		}

		// Among the 31 possible values of subarray ending at i, find the one == k
		if cnt, ok := cur[k]; ok {
			res += cnt
		}

		// pass cur to prev
		prev = cur
	}

	return int64(res)

}
