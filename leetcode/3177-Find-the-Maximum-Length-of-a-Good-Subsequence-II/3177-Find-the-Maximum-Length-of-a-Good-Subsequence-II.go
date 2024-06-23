package leetcode

func maximumLength(nums []int, k int) int {
	n := len(nums)

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, k+1)
	}

	// maxEqual[t][nums[i]] := max(dp[0:i-1][t] && nums[j] == nums[i]) = [{equal element: max dp value in 0:i-1} ... ]
	// maxEqual: given a cost of t, the past max dp value in [0:i-1] (inclusive) && nums[j] == nums[i]
	maxEqual := make([]map[int]int, k+1)
	for t := 0; t <= k; t++ {
		maxEqual[t] = make(map[int]int)
	}

	// maxAll[t] := max(dp[0:i-1][t])
	// maxAll: given a cost of t, the past max dp value in [0:i-1] (inclusive)
	maxAll := make([]int, k+1)

	res := 1
	for i := 0; i < n; i++ {
		for t := 0; t <= k; t++ {
			cur := 1

			// case 1: nums[j] == nums[i]
			if val, ok := maxEqual[t][nums[i]]; ok {
				cur = max(cur, val+1)
			}

			// case 2: nums[j] != nums[i]
			if t-1 >= 0 {
				cur = max(cur, maxAll[t-1]+1)
			}

			dp[i][t] = cur
			res = max(res, dp[i][t])
		}

		// update maxEqual and maxAll
		for t := 0; t <= k; t++ {
			maxEqual[t][nums[i]] = max(maxEqual[t][nums[i]], dp[i][t])
			maxAll[t] = max(maxAll[t], dp[i][t])
		}
	}
	
	return res
}
