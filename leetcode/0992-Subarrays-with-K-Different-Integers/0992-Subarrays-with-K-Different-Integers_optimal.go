package leetcode

// Optimal Solution
func subarraysWithKDistinct_optimal(nums []int, k int) int {
	n := len(nums)
	res := 0

	// atMostK(k) - atMostK(k-1) = exactlyK(k)
	upperLeft := upperBound(nums, k-1)
	upperRight := upperBound(nums, k)

	for i := 0; i < n; i++ {
		if upperLeft[i] == -1 || upperRight[i] == -1 {
			continue
		}
		res += upperRight[i] - upperLeft[i]
	}
	return res
}