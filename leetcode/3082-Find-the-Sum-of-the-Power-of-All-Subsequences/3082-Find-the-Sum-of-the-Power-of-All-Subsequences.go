package leetcode

var M = int(1e9 + 7)

func sumOfPower(nums []int, k int) int {
	n := len(nums)

	// dp[i][s][l] := # of subsequences in nums[1:i] (inclusive) with their sum = s AND their length = l
	dp := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, n+1)
		}
	}
	nums = append([]int{0}, nums...)

	// base case: empty array + sum=0 + length=0 => 1 subsequence (itself)
	dp[0][0][0] = 1
	// recurrence
	// i := 1 -> n meaningful elements of nums[] starts from nums[1:]
	// s := 0 -> k
	// l := 0 -> i, because for array of length i, the length of a subseq can be 0 up to i
	for i := 1; i <= n; i++ {
		for s := 0; s <= k; s++ {
			for l := 0; l <= i; l++ { // i here represents max length of subseq in nums[1:i] that we can have for l
				dp[i][s][l] = dp[i-1][s][l]     // exclude nums[i]
				if s-nums[i] >= 0 && l-1 >= 0 { // include nums[i]
					dp[i][s][l] += dp[i-1][s-nums[i]][l-1]
				}
				dp[i][s][l] %= M
			}
		}
	}

	// total # of subseqs in array of m elements = 2^m where 0 <= m <= n
	power := make([]int, n+1)
	power[0] = 1 // 2^0 = 1
	for i := 1; i <= n; i++ {
		power[i] = 2 * power[i-1]
		power[i] %= M
	}

	// dp[n][k][l] = total # of kinds of subseqs in nums[] with their sum = k AND length = l
	// how many superseqs can contain dp[n][k][l]?
	// 2^(n-l): it means how many choices we have to select the rest of the elements (n-l) to be included in the superseq
	// Thus, for a given l, total # of superseqs = 2^(n-l) * dp[n][k][l] (total kinds of subseqs with length = l)
	res := 0
	for l := 1; l <= n; l++ {
		res += power[n-l] * dp[n][k][l]
		res %= M
	}
	return res

}
