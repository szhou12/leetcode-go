package leetcode

// My Solution - DFS - All Subsets
// # branches = 2 -> current num is positive / current num is negative
// # levels = len(nums)
func findTargetSumWays(nums []int, target int) int {
	var res int

	dfs(nums, target, 0, 0, &res)
	return res
}

func dfs(nums []int, target int, index int, curSum int, res *int) {
	// Base case
	if index == len(nums) {
		if curSum == target {
			*res++
		}
		return
	}

	dfs(nums, target, index+1, curSum+nums[index], res)
	dfs(nums, target, index+1, curSum-nums[index], res)
}

// Optimal Solution - DP - 0/1 knapsack
// DP[i]: # of ways to form positive sequence that sums up to == i
func findTargetSumWays_DP(nums []int, target int) int {
	total := sum(nums)
	if total < target || (target+total)%2 != 0 || target+total < 0 {
		return 0
	}

	// 计算背包大小
	bagSize := (target + total) / 2
	// 定义dp数组
	dp := make([]int, bagSize+1)
	// Base Case
	dp[0] = 1 // 装满容量为0的背包，有1种方法，就是装0件物品
	// Recursion
	for i := 0; i < len(nums); i++ {
		for j := bagSize; j >= nums[i]; j-- { // 终止在nums[i] 因为容量小于nums[i]就肯定没法装nums[i]
			dp[j] += dp[j-nums[i]]
		}
	}
	return dp[bagSize]
}

func sum(nums []int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}
