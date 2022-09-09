package leetcode

import "math"

func distributeCookies(cookies []int, k int) int {
	res := math.MaxInt
	plan := make([]int, 8)
	dfs(cookies, k, 0, &plan, &res)
	return res
}

func dfs(cookies []int, k int, index int, plan *[]int, res *int) {
	// base case
	if index == len(cookies) {
		// find the unfairness of current permutation
		unfair := 0
		for i := 0; i < len(*plan); i++ {
			unfair = max(unfair, (*plan)[i])
		}
		//update global result
		*res = min(*res, unfair)
		return
	}

	// level = cookie i; k branches
	for person := 0; person < k; person++ {
		(*plan)[person] += cookies[index]
		dfs(cookies, k, index+1, plan, res)
		(*plan)[person] -= cookies[index]
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
