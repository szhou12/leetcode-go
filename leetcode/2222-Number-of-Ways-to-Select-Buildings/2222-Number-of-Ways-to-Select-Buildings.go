package leetcode

// dp[i][j][k] := after looking at s[i], there have been j (=0,1,2,3) elements selected, and the last element is k (=0,1)
func numberOfWays(s string) int64 {
	n := len(s)

	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 4)
		for j := 0; j <= 3; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	// base cases
	dp[0][0][0] = 1
	dp[0][0][1] = 1
	dp[0][1][int(s[0]-'0')] = 1

	// recurrence
	for i := 1; i < n; i++ {
		for j := 0; j <= 3; j++ {
			for k := 0; k < 2; k++ {
				// not selecting i
				dp[i][j][k] = dp[i-1][j][k]

				// selecting i
				if j-1 >= 0 && int(s[i] - '0') == k {
					dp[i][j][k] += dp[i-1][j-1][1-k]
				}
				
			}
		}
	}

	res := dp[n-1][3][0] + dp[n-1][3][1]
	return int64(res)
}

// 写法二: 前置一个虚空占位符，方便写base cases
func numberOfWays2(s string) int64 {
    n := len(s)
    s = "#"+s

    dp := make([][][]int, n+1)
    for i := 0; i <= n; i++ {
        dp[i] = make([][]int, 4)
        for j := 0; j <= 3; j++ {
            dp[i][j] = make([]int, 2)
        }
    }

    // base cases: 0号位为虚空，所以什么都不选(j=0)也是一种valid选择
    dp[0][0][0] = 1
    dp[0][0][1] = 1

    for i := 1; i <= n; i++ {
        for j := 0; j <= 3; j++ {
            for k := 0; k < 2; k++ {
                // not selecting i
                dp[i][j][k] = dp[i-1][j][k]
                // selecting i
                if j-1 >= 0 && int(s[i]-'0') == k {
                    dp[i][j][k] += dp[i-1][j-1][1-k]
                }
                
            }
        }
    }

    res := dp[n][3][0] + dp[n][3][1]
    return int64(res)

}