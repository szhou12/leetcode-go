package leetcode

import "sort"

// Soln 1: Greedy - 每次merge下一行，取前k个最小的和
func kthSmallest_greedy(mat [][]int, k int) int {

	sum := []int{0}
	for i := 0; i < len(mat); i++ { // O(m)
		var temp []int
		for _, j := range sum { // iter top k sums O(k)
			for _, t := range mat[i] { // O(n)
				temp = append(temp, j+t)
			}
		}

		// filter top k sums
		sort.Ints(temp) // O(kn*log(kn))
		sum = temp[:min(k, len(temp))]
	}

	return sum[k-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Soln 2: Binary Search 猜答案 -- Optimal Solution
func kthSmallest(mat [][]int, k int) int {
	m, n := len(mat), len(mat[0])
	var left, right int
	for r := 0; r < m; r++ {
		left += mat[r][0]
		right += mat[r][n-1]
	}

	for left < right { // O(logD)
		mid := left + (right-left)/2
		if countMoreOrEqualK(mat, mid, k) { // # of combos whose sum <= mid is >= k
			right = mid
		} else { // # of combos whose sum <= mid is < k -> mid 一定不是答案
			left = mid + 1
		}
	}
	return left
}

func countMoreOrEqualK(mat [][]int, mid int, k int) bool {
	m := len(mat)

	var sum int
	for r := 0; r < m; r++ {
		sum += mat[r][0]
	}
	count := 1

	// edge case: smallest sum already > mid
	if sum > mid {
		return false
	}

	dfs(0, mat, mid, k, sum, &count) // O(k)

	if count >= k {
		return true
	}
	return false
}

func dfs(row int, mat [][]int, target int, k int, sum int, count *int) {
	m, n := len(mat), len(mat[0])

	// base cases
	if *count >= k {
		return
	}
	if row >= m {
		return
	}

	// recursion
	for c := 0; c < n; c++ {
		if sum+mat[row][c]-mat[row][0] <= target {
			if c > 0 {
				(*count)++
			}
			dfs(row+1, mat, target, k, sum+mat[row][c]-mat[row][0], count)
		}
	}
}
