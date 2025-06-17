package leetcode

func maxFrequency(nums []int, k int) int {
	n := len(nums)

	// prefix sum: prefixSum[i] = number of k's in nums[0:i] (inclusive)
	prefixSum := make([]int, n)
	if nums[0] == k {
		prefixSum[0] = 1
	}
	for i := 1; i < n; i++ {
		if nums[i] == k {
			prefixSum[i] = prefixSum[i-1] + 1
		} else {
			prefixSum[i] = prefixSum[i-1]
		}
	}

	left, right := make(map[int]int), make(map[int]int)
	freq := make(map[int]int)
	maxGain := 0
	for i, v := range nums {
		if v == k {
			continue
		}

		if _, ok := left[v]; !ok {
			left[v] = i
		}
		freq[v]++
		right[v] = i

		// count k's between left[v] and i=right[v]
		l, r := left[v], right[v]
		var kcounts int
		if l == 0 {
			kcounts = prefixSum[r]
		} else {
			kcounts = prefixSum[r] - prefixSum[l-1]
		}
		gain := freq[v] - kcounts
		// if gain <= 0, not a good subarray to gain extra k freqs, throw away
		if gain <= 0 {
			left[v] = i
			freq[v] = 1
		}
		maxGain = max(maxGain, gain)

	}

	return maxGain + prefixSum[n-1]

}