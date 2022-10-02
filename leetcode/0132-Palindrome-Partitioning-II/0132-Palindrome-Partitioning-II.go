package leetcode

// DP
// dp[i] = min cuts for s[0...i] (范围包括i)
func minCut(s string) int {
	n := len(s)
	isPal := make([][]bool, n)
	for i := 0; i < n; i++ {
		isPal[i] = make([]bool, n)
	}

	// 填写 isPal 的另一种写法:
	// for length := 1; length <= n; length++ {
	//     for i := 0; i <= n-length; i++ {
	//         j := i+length-1
	//         if s[i] == s[j] {
	//             if i+1 >= j - 1 {
	//                 isPal[i][j] = true
	//             } else {
	//                isPal[i][j] = isPal[i+1][j-1]
	//             }
	//         }
	//     }
	// }
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if i >= j { // Base Cases
				isPal[i][j] = true
				continue
			}
			if s[i] == s[j] {
				isPal[i][j] = isPal[i+1][j-1]
			}
		}
	}

	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i // max cuts for length=i are i # of cuts
	}

	for i := 0; i < n; i++ {
		if isPal[0][i] {
			dp[i] = 0
			continue
		}
		for j := 0; j < i; j++ {
			if isPal[j+1][i] {
				dp[i] = min(dp[i], dp[j]+1)
			}
		}
	}

	return dp[n-1]
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
