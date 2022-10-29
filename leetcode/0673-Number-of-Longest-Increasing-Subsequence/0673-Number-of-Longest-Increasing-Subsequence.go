package leetcode

// DP - optimal solution
// 维护两个记事本: dp[] 和 count[]
// dp[i]: length of LIS ending at i
// count[i]: number of LIS ending at i
func findNumberOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	count := make([]int, n)
	// Base cases
	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
	}

	maxLen := 1
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// update count[i]
				if dp[j]+1 > dp[i] { // if find longer (new LIS), replace current count[i]
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] { // if equal length of current LIS, add to current count[i]
					count[i] += count[j]
				}
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// Need to record global LIS happened at where (ending at i)
		maxLen = max(maxLen, dp[i])
	}

	// calculate total # of LIS
	res := 0
	for i := 0; i < n; i++ {
		if dp[i] == maxLen {
			res += count[i]
		}
	}
	return res

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS - 暴力解 - 会超时
func findNumberOfLIS_DFS(nums []int) int {
	curCombo := make([]int, 0)
	var res [][]int
	dfs(nums, 0, curCombo, &res)
	count := 0
	maxLen := 0
	for _, combo := range res {
		if len(combo) > maxLen {
			count = 1
			maxLen = len(combo)
		} else if len(combo) == maxLen {
			count++
		}
	}
	return count
}

// All Subsets第二种写法
func dfs(nums []int, index int, curCombo []int, res *[][]int) {
	// 沿着recursion tree添加结果，而不是在base case处添加
	c := make([]int, len(curCombo))
	copy(c, curCombo)
	*res = append(*res, c)

	for i := index; i < len(nums); i++ {
		if len(curCombo) == 0 || nums[i] > curCombo[len(curCombo)-1] {
			curCombo = append(curCombo, nums[i])
			dfs(nums, i+1, curCombo, res)
			curCombo = curCombo[:len(curCombo)-1]
		}
	}
}
