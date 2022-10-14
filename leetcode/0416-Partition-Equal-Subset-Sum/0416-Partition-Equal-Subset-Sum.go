package leetcode

// DP - 0/1 knapsack
func canPartition(nums []int) bool {
	totalSum := sum(nums)
	//Edge Case
	if totalSum%2 != 0 {
		return false
	}

	n := len(nums)
	target := totalSum / 2
	nums = append([]int{-1}, nums...) // prepend a placeholder for better indexing
	dp := make([][]bool, n+1)

	// Base Cases
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, target+1) // 0 -> target
		dp[i][0] = true                // always able to have empty subset
	}

	// Recurrence
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			if j >= nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
				// dp[i-1][j] 不把 nums[i] 放入背包的情况
				// dp[i-1][j-nums[i]] 把 nums[i] 放入背包的情况
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[n][target]
}

// DP - optimal solution
func canPartition_Opt(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果 nums 的总和为奇数则不可能平分成两个子集
	if sum%2 == 1 {
		return false
	}

	target := sum / 2
	dp := make([]int, target+1)

	for _, num := range nums {
		for j := target; j >= num; j-- {
			if dp[j] < dp[j-num]+num {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[target] == target
}

// DFS暴力解 - 会超时
// 逻辑: 求出所有可能的subsets, 然后看哪个的sum是total sum 的一半
func canPartition_DFS(nums []int) bool {
	subset := make([]int, 0)
	totalSum := sum(nums)
	var res [][]int
	dfs(nums, subset, 0, totalSum, &res)

	if len(res) > 0 {
		return true
	}
	return false
}

func dfs(nums []int, subset []int, index int, totalSum int, res *[][]int) {
	// Base case
	if index == len(nums) {
		if len(subset) < len(nums) && 2*sum(subset) == totalSum {
			subsetCopy := make([]int, len(subset))
			copy(subsetCopy, subset)
			*res = append(*res, subsetCopy)
		}
		return
	}

	// DFS
	subset = append(subset, nums[index])
	dfs(nums, subset, index+1, totalSum, res)
	subset = subset[:len(subset)-1]

	dfs(nums, subset, index+1, totalSum, res)

}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}

// # levels = len(nums)
// # branches = 2 (add / not add i-th element)
