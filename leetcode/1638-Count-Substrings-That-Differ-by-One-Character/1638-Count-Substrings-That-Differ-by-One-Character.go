package leetcode

// left[i][j] = length (= # of chars) of longest equivalent substring between s[0:i] and t[0:j]; this substring in s[0:i] must end at i and this substring in t[0:j] must end at j
// right[i][j] = length (= # of chars) of longest equivalent substring between s[i:m-1] and t[j:n-1]; this substring in s[i:m-1] must end at i and this substring in t[j:n-1] must end at j

// base case: left[0][0] = 0
// recurrence: left[i][j] = left[i-1][j-1] + 1 if s[i] == t[j]
//                        = 0                  if s[i] != t[j]
func countSubstrings(s string, t string) int {
	m, n := len(s), len(t)
	s = "#" + s + "#"
	t = "#" + t + "#"

	left := make([][]int, m+2)
	for i := 0; i < len(left); i++ {
		left[i] = make([]int, n+2)
	}
	// recurrence
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i] == t[j] {
				left[i][j] = left[i-1][j-1] + 1
			} else {
				left[i][j] = 0
			}
		}
	}

	right := make([][]int, m+2)
	for i := 0; i < len(right); i++ {
		right[i] = make([]int, n+2)
	}
	// recurrence
	for i := m; i >= 1; i-- {
		for j := n; j >= 1; j-- {
			if s[i] == t[j] {
				right[i][j] = right[i+1][j+1] + 1
			} else {
				right[i][j] = 0
			}
		}
	}

	res := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s[i] != t[j] {
				// 注意: 这里L段和R段都有+1, 表示i,j左手边/右手边什么都没有的情况, 就只有i,j自己
				res += (left[i-1][j-1] + 1) * (right[i+1][j+1] + 1)
			}
		}
	}

	return res
}
