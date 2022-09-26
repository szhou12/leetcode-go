package leetcode

func combine(n int, k int) [][]int {
	var res [][]int
	var cur []int
	dfs(n, k, 1, cur, &res)

	return res
}

func dfs(n int, k int, val int, cur []int, res *[][]int) {
	// Base Case
	if len(cur) == k {
		combo := make([]int, len(cur))
		copy(combo, cur)
		*res = append(*res, combo)
		return
	}

	// Early Stopping: 有可能n个数都试完了，但是放进subset的不足k个
	if val > n {
		return
	}

	// Adding current value
	cur = append(cur, val)
	dfs(n, k, val+1, cur, res)
	cur = cur[:len(cur)-1]

	// Not adding current value
	dfs(n, k, val+1, cur, res)
}
