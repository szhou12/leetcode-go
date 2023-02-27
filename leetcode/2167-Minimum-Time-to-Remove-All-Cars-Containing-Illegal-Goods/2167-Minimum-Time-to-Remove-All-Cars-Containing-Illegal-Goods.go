package leetcode

// left[i] := min time to remove 1 in s[0:i] using type 1 operation
// right[i] := min time to remove 1 in s[i:n-1] using type 2 OR type 3 operation
func minimumTime(s string) int {
	n := len(s)

	left := make([]int, n)
	// base case
	if s[0] == '1' {
		left[0] = 1
	}
	// recurrence
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			left[i] = i - 0 + 1
		} else {
			left[i] = left[i-1]
		}
	}

	right := make([]int, n)
	// base case
	if s[n-1] == '1' {
		right[n-1] = 1
	}
	// recurrence
	for i := n - 2; i >= 0; i-- {
		if s[i] == '1' {
			right[i] = min((n-1)-i+1, right[i+1]+2)
		} else {
			right[i] = right[i+1]
		}
	}

	res := min(left[n-1], right[0])
	for i := 0; i+1 < n; i++ {
		res = min(res, left[i]+right[i+1])
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
