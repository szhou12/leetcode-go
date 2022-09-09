package leetcode

import "math"

// Optimized Solution - Binary Search + DFS
func distributeCookies_BSDFS(cookies []int, k int) int {
	plan := make([]int, 8)

	// Binary Search on max amount of cookies a person can have
	// mid := upper limit
	left := 1
	right := math.MaxInt
	for left < right {

		// need to reset plan for every trial
		for i := 0; i < 8; i++ {
			plan[i] = 0
		}

		mid := left + (right-left)/2
		if dfs2(cookies, k, 0, &plan, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func dfs2(cookies []int, k int, index int, plan *[]int, limit int) bool {
	// Base case
	if index == len(cookies) {
		return true
	}

	for person := 0; person < k; person++ {
		// if current person exceeds cookies limit, skip
		if (*plan)[person]+cookies[index] > limit {
			continue
		}
		(*plan)[person] += cookies[index]
		if dfs2(cookies, k, index+1, plan, limit) {
			return true
		}
		// if giving current person the index-th cookie results in none of workable soln, take it away
		(*plan)[person] -= cookies[index]
	}

	return false
}

// Brute force - DFS
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
