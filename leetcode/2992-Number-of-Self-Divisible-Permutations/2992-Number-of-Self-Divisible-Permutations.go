package leetcode

func selfDivisiblePermutationCount(n int) int {
	// dp[i][state] := # of permutations from the first i positions by using the digits marked as 1 in 'state'
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 1<<n)
	}

	// base case
	dp[0][0] = 1 // 表示没有任何数字的时候，state没用到任何数字的时候，一个'空'的排列是一种合法的排列
	// recurrence
	for i := 1; i <= n; i++ { // 看第i个位置
		for state := 0; state < (1<<n); state++ { // 挑一种状态 (可优化: 只从1的个数是i个的state中选)
			for digit := 1; digit <= n; digit++ {
				if GCD(digit, i) != 1 {
					continue
				}
				if state >> (digit-1) & 1 == 0 { // state表示bit array从右边开始0-indexed
					continue
				}
				dp[i][state] += dp[i-1][state-(1<<(digit-1))] // 先前的state是挖掉当前的digit
			}
		}
	}

	// (1<<n)-1 := state 1111...111
	return dp[n][(1<<n)-1]

}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// Hint from constraint: 1 <= n <= 12
// 如果全排列 = 12! = 4*10^8 太高
// 如果是 2^12 = 4096 是可以的
// 推断算法复杂度大概是 n * 2^n
// 2^n 表示什么？state: 2^n 表示的是每个数字的状态，是否使用过
// e.g. 1001000111 表示 1, 4, 7, 8, 9 位子上的数字使用过了
// 进而推导出使用DP来解: dp[i][state]
// dp[i][state] := # of self-divisible permutations from the first i positions by using the digits marked as 1 in 'state'