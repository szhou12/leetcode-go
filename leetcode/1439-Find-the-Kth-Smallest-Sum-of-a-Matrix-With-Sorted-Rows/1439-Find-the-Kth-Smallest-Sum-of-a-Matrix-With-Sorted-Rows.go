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
