package leetcode

func kthCharacter_recursion(k int64, operations []int) byte {
	var n int64 = 1
	times := 0
	for n < k {
		n = 2 * n
		times++
	}

	shift := dfs(n, k, times-1, operations)

	return byte('a' + (shift % 26))
}

func dfs(n int64, k int64, t int, operations []int) int {
	// base case
	if t < 0 {
		return 0
	}

	shift := 0
	if k > n/2 {
		if operations[t] == 0 {
			shift += dfs(n/2, k-n/2, t-1, operations)
		} else {
			shift += dfs(n/2, k-n/2, t-1, operations) + 1
		}
	} else {
		shift += dfs(n/2, k, t-1, operations)
	}

	return shift
	
}
