package leetcode

// Solution: Prefix Sum
// O(n^2)
func subarraySum_prefixSum(nums []int, k int) int {
	n := len(nums)
	// prefixSum[i] = sum(nums[0:i]) (inclusive)
	prefixSum := make([]int, n)
	prefixSum[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	res := 0

	// try all subarrays [i:j]
	// NOTE: in this implementation - prefixSum[] of length n, prefixSum[j] - prefixSum[i] will miss all subarrays starting from 0
	for i := 0; i < n; i++ {
		if prefixSum[i] == k {
			res++
		}
		for j := i+1; j < n; j++ {
			if prefixSum[j] - prefixSum[i] == k {
				res++
			}
		}
	}

	return res
}

// Optimal Solution: Prefix Sum + Hash Map (frequency map)
// O(n)
func subarraySum(nums []int, k int) int {
	n := len(nums)
	dict := make(map[int]int) // {prefix sum: count}
	dict[0] = 1 // MUST HAVE THIS. 

	sum := 0

	res := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if cnt, ok := dict[sum-k]; ok {
			res += cnt
		}
		dict[sum]++
	}

	return res
}