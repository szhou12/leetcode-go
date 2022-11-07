package leetcode

import (
	"math"
	"sort"
)

// Greedy
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) > abs(nums[j])
	})

	for i := 0; i < len(nums); i++ {
		if k > 0 && nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	// 如果k有剩余, 反复negate 最后一个元素
	// 剩余的k为偶数，就可以恢复原样
	if k%2 == 1 { // 剩余的k为奇数, negate一次
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	return sum(nums)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// DFS暴力解: 尝试所有可以变换的地方 - 会超时!!!
// # levels = k
// # branches = len(nums)
func largestSumAfterKNegations_DFS(nums []int, k int) int {
	res := math.MinInt
	dfs(nums, k, &res)
	return res
}

func dfs(nums []int, k int, res *int) {
	// base case
	if k == 0 {
		*res = max(*res, sum(nums))
		return
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = -nums[i]
		dfs(nums, k-1, res)
		nums[i] = -nums[i]
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
