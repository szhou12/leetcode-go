package leetcode

func largestVariance(s string) int {
	n := len(s)

	count := make([]int, 26)
	for _, char := range s {
		count[char-'a'] += 1
	}

	ret := 0
	// Try out all possible distinct letter pairs;
	// treat one as 1, the other as -1, treat all remaining letters as 0
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if count[i] == 0 || count[j] == 0 {
				continue
			}
			if i == j {
				continue
			}

			nums := make([]int, n)
			for k := 0; k < n; k++ {
				if int(s[k]) == 'a'+i {
					nums[k] = 1
				} else if int(s[k]) == 'a'+j {
					nums[k] = -1
				}
			}
			// helper() implements max sum subarray (057) from left-to-right and from right-to-left
			ret = max(ret, helper(nums))

		}
	}
	return ret

}

func helper(nums []int) int {
	n := len(nums)
	ret := 0

	dp1 := make([]int, n)
	dp1[0] = nums[0]
	// dp1: from left-to-right
	for i := 1; i < n; i++ {
		dp1[i] = max(dp1[i-1]+nums[i], nums[i])
	}

	curSum := 0
	// dp2: from right-to-left
	for i := n - 1; i >= 0; i-- {
		curSum = max(curSum+nums[i], nums[i])

		// ONLY update results at i-th element where nums[i] == -1
		// B/c we only consider subarrays that have at least one -1
		if nums[i] == -1 {
			ret = max(ret, dp1[i]+curSum-nums[i])
		}
	}

	return ret
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
