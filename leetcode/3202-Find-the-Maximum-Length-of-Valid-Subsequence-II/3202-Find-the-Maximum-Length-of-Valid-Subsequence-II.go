package leetcode

func maximumLength(nums []int, k int) int {
	dp := make([][]int, k)
	for i := 0; i < k; i++ {
		dp[i] = make([]int, k)
	}

	res := 1
	for _, num := range nums {
		cur := num % k
		for r := 0; r < k; r++ {
			prev := (r - cur + k) % k
			dp[cur][r] = max(dp[cur][r], dp[prev][r]+1)
			res = max(res, dp[cur][r])
		}

	}

	return res
}