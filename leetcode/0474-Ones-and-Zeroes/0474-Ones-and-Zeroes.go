package leetcode

// DP - Optimal Solution - 0/1 Knapsack
func findMaxForm(strs []string, m int, n int) int {

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
