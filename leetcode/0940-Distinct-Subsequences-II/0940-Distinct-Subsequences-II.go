package leetcode

var M = int(1e9 + 10)

func distinctSubseqII(s string) int {
	dp := make([]int, 26)

	for i := 0; i < len(s); i++ {
		dp[s[i]-'a'] = (sum(dp) + 1) % M // Note: LHS DP def is different from RHS DP def BUT they are equivalent
	}

	return sum(dp) % M
}

func sum(a []int) int {
	res := 0

	for _, count := range a {
		res += count % M
	}
	return res
}
