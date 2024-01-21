package leetcode

func beautifulIndices(s string, a string, b string, k int) []int {
	aIndices := KMP(s, a) // O(n)
	bIndices := KMP(s, b) // O(n)

	res := make([]int, 0)

	// early return: if aIndices or bIndices is empty, then no beautiful indices
	if len(aIndices) == 0 || len(bIndices) == 0 {
		return res
	}

	for _, i := range aIndices { // O(n)
		lowbound := lowerBound(bIndices, i-k) // O(logn)
		upbound := upperBound(bIndices, i+k) // O(logn)
		// contains at least such j in the range [i-k, i+k] (inclusive)
		if upbound-lowbound >= 1 {
			res = append(res, i)
		}
	}

	return res
}

// Find all occurrences (start index) of needle in haystack
func KMP(haystack string, needle string) []int {
	n := len(haystack)
	m := len(needle)

	res := make([]int, 0)

	// edge cases: empty needle or haystack
	if m == 0 || n == 0 {
		return res
	}

	// dp[i] := max length j at haystack[i] s.t. (prefix) needle[0:j-1] == haystack[i-j+1:i] (suffix)
	dp := make([]int, n)

	// base case: check if needle[0] == haystack[0]
	if needle[0] == haystack[0] {
		dp[0] = 1
	}

	// Early check: if len(needle) == 1 AND dp[0] = 1
	// MUST check this before the loop bc the loop is for i > 0
	if m == 1 && dp[0] == 1 {
		res = append(res, 0)
	}

	lsp := preprocess(needle)

	// recurrence
	for i := 1; i < n; i++ {
		j := dp[i-1]
		// Go的特性 (区别于C++): 因为是寻找所有位置，必须判断 j==m，当dp[i-1]刚好包含一个完整子串，第i号位要重新开始匹配，不能接上dp[i-1]
		for j > 0 && (j == m || needle[j] != haystack[i]) {
			j = lsp[j-1]
		}
		if needle[j] == haystack[i] {
			dp[i] = j + 1
		} else {
			dp[i] = j
		}

		// check if haystack[i-j+1:i] is a needle
		if dp[i] == m {
			res = append(res, i-m+1)
		}
	}

	return res
}


func preprocess(needle string) []int {
	n := len(needle)

	// dp[i] := max length j at needle[i] s.t. (prefix) needle[0:j-1] == needle[i-j+1:i] (suffix)
	dp := make([]int, n)

	// base case: we require true prefix and true suffix, so dp[0] = 0
	dp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		j := dp[i-1]
		if j > 0 && needle[j] != needle[i] {
			j = dp[j-1]
		}
		if needle[j] == needle[i] {
			dp[i] = j + 1
		} else {
			dp[i] = j
		}
	}

	return dp
}

// Find smallest index i s.t. nums[i] >= target
func lowerBound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] >= target {
		return left
	} else {
		return left + 1
	}
}

// Find smallest index i s.t. nums[i] > target
func upperBound(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if nums[left] > target {
		return left
	} else {
		return left + 1
	}
}
