package leetcode

// DP - 回头看
func peopleAwareOfSecret_DP_backward(n int, delay int, forget int) int {
	modulo := int64(1e9 + 7)

	// DP[i] := NEWLY added people who know the secret on i-th day
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := i - delay; j > i-forget; j-- {
			if j >= 1 {
				dp[i] += dp[j]
				dp[i] = int(int64(dp[i]) % modulo)
			}
		}
	}

	// Result
	res := 0
	for i := 1; i <= n; i++ {
		if i+forget > n {
			res += dp[i]
			res = int(int64(res) % modulo)
		}
	}
	return res
}

// DP - 向前看
func peopleAwareOfSecret_DP_forward(n int, delay int, forget int) int {
	modulo := int64(1e9 + 7)

	// DP[i] := NEWLY added people who know the secret on i-th day
	dp := make([]int, n+1)
	dp[1] = 1

	// ONLY DIFFERENCE
	for i := 1; i <= n; i++ {
		for j := i + delay; j < i+forget; j++ {
			if j <= n {
				dp[j] += dp[i]
				dp[j] = int(int64(dp[j]) % modulo)
			}
		}
	}

	// Result
	res := 0
	for i := 1; i <= n; i++ {
		if i+forget > n {
			res += dp[i]
			res = int(int64(res) % modulo)
		}
	}
	return res
}

// Optimal solution - 差分数组
func peopleAwareOfSecret(n int, delay int, forget int) int {
	modulo := int64(1e9 + 7)

	dp := make([]int, n+1)
	// dp[1] = 1

	diff := make([]int, n+1)
	diff[1] += 1
	diff[2] -= 1

	for i := 1; i <= n; i++ {
		// for j := i + delay; j < i+forget; j++ {
		// 	if j <= n {
		// 		dp[j] += dp[i]
		// 		dp[j] = int(int64(dp[j]) % modulo)
		// 	}
		// }
		dp[i] = int((int64(dp[i-1]+diff[i]) + modulo) % modulo)
		if i+delay <= n {
			diff[i+delay] += dp[i]
		}
		if i+forget <= n {
			diff[i+forget] -= dp[i]
		}
	}

	res := 0
	for i := 1; i <= n; i++ {
		if i+forget > n {
			res = res + dp[i]
			res = int(int64(res) % modulo)
		}
	}
	return res
}
