package leetcode

func countGood(nums []int, k int) int64 {
	n := len(nums)

	res := 0                  // # of good subarrays that have >= k 'good' pairs
	count := 0                // # of pairs in the window
	freq := make(map[int]int) // freq of each num in the window
	right := 0
	for left := 0; left < n; left++ {
		// 吃
		for right < n && count < k { // 左闭右开 [left, right)
			if _, ok := freq[nums[right]]; !ok {
				freq[nums[right]] = 1
			} else {
				// Order matters for the following two lines!!!
				count += freq[nums[right]]
				freq[nums[right]]++
			}
			right++
		}

		// update result: 分类讨论!!!
		if right < n && !(count < k) {
			res += (n - 1) - (right - 1) + 1 // all subarrays from ending at right-1 to ending at n-1 are good
		} else if !(right < n) && count < k {
			res += 0
		} else {
			res += 1 // only the subarray ending at n-1 is good. write res += (n - 1) - (right - 1) + 1 also works
		}

		// 吐
		// Order matters for the following two lines!!!
		freq[nums[left]]--
		count -= freq[nums[left]]
	}

	return int64(res)

}
