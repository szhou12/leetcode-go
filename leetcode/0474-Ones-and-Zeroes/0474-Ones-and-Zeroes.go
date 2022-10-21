package leetcode

import "strings"

// DP - Optimal Solution - 0/1 Knapsack
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}

	for _, curStr := range strs {
		zeros := strings.Count(curStr, "0")
		ones := len(curStr) - zeros

		// 遍历背包容量且从后向前遍历！
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}

	return dp[m][n]
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// DFS - All Subsets - 超时！
// # branches = 2
// # levels = len(strs)
func findMaxForm_DFS(strs []string, m int, n int) int {

	var res int
	var cur []string
	dfs(strs, m, n, 0, cur, &res)
	return res
}

func dfs(strs []string, m int, n int, index int, cur []string, res *int) {
	// base case
	if index == len(strs) {
		if valid(cur, m, n) && len(cur) > *res {
			*res = len(cur)
		}
		return
	}

	cur = append(cur, strs[index])
	dfs(strs, m, n, index+1, cur, res)
	cur = cur[:len(cur)-1]

	dfs(strs, m, n, index+1, cur, res)
}

func valid(cur []string, m int, n int) bool {
	ones, zeros := 0, 0
	for _, str := range cur {
		for _, char := range str {
			if char-'0' == 0 {
				zeros++
			} else if char-'0' == 1 {
				ones++
			}
		}
	}

	if zeros <= m && ones <= n {
		return true
	}
	return false
}
