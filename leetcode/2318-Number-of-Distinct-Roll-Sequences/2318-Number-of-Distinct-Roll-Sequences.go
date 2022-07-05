package leetcode

// DP - Optimal solution
func distinctSequences(n int) int {
	// Edge case: n == 1
	if n == 1 {
		return 6
	}

	modulo := int64(1e9 + 7)
	dp := make([][][]int, n+1)

	for i := 0; i <= n; i++ {
		dp[i] = make([][]int, 6+1)
		for a := 0; a <= 6; a++ {
			dp[i][a] = make([]int, 6+1)
		}
	}

	count := 0
	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			if a != b && gcd(a, b) == 1 {
				dp[2][a][b] += 1
				count += 1
			}
		}
	}

	// Edge case: n == 2
	if n == 2 {
		return count
	}

	// DP
	for i := 3; i <= n; i++ {
		for a := 1; a <= 6; a++ {
			for b := 1; b <= 6; b++ {
				for c := 1; c <= 6; c++ {
					if a != b && gcd(a, b) == 1 && a != c {
						dp[i][b][a] = int(int64(dp[i][b][a]+dp[i-1][c][b]) % modulo)
					}
				}
			}
		}
	}

	// Result
	res := 0
	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			res = int(int64(res+dp[n][b][a]) % modulo)
		}
	}
	return res

}

// DFS - can pass ONLY when n small
// n levels
// each level at most 6 branches bc 6-sided dice
//     {}
// / | | | | \
// 1 2 3 4 5  6
// /
// 2-6
func distinctSequences_DFS(n int) int {
	var cur []int
	var result [][]int

	dfs(n, cur, &result)
	return len(result)
}

func dfs(n int, cur []int, result *[][]int) {
	// base case
	if n == 0 {
		res := make([]int, len(cur))
		copy(res, cur)
		*result = append(*result, res)
		return
	}

	curLen := len(cur)
	for i := 1; i <= 6; i++ {
		if curLen < 1 { // curLen == 0 ==> directly append i
			cur = append(cur, i)
			dfs(n-1, cur, result)
			cur = cur[:len(cur)-1]
		} else if curLen < 2 { // 1 <= curLen < 2 => curLen == 1 ==> check gcd and prev and cur values equal
			if gcd(i, cur[len(cur)-1]) == 1 && i != cur[len(cur)-1] {
				cur = append(cur, i)
				dfs(n-1, cur, result)
				cur = cur[:len(cur)-1]
			}
		} else { // curLen >= 2 ==> check gcd & prev prev, prev and cur values equal
			if gcd(i, cur[len(cur)-1]) == 1 && i != cur[len(cur)-1] && i != cur[len(cur)-2] {
				cur = append(cur, i)
				dfs(n-1, cur, result)
				cur = cur[:len(cur)-1]
			}
		}
	}
}

// Euclidean Algorithm
func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
