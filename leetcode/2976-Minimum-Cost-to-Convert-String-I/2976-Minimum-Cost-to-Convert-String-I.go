package leetcode

import "math"

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	dp := make([][]int, 26)
	for i := 0; i < 26; i++ {
		dp[i] = make([]int, 26)
		for j := 0; j < 26; j++ {
			if i != j {
				dp[i][j] = math.MaxInt / 3
			}
		}
	}

	// Floyd
	for k, _ := range original {
		from := int(original[k] - 'a')
		to := int(changed[k] - 'a')
		weight := cost[k]
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				dp[i][j] = min(dp[i][j], dp[i][from]+weight+dp[to][j])
			}
		}
	}

	res := 0
	for i, _ := range source {
		src := int(source[i] - 'a')
		tar := int(target[i] - 'a')
		if dp[src][tar] >= math.MaxInt/3 {
			return -1
		}
		res += dp[src][tar]
	}
	return int64(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
