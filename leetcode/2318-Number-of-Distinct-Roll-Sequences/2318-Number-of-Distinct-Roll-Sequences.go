package leetcode

// DFS - can pass ONLY when n small
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

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
