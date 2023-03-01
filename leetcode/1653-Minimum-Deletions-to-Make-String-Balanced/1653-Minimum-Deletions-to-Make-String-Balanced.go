package leetcode

// Core idea: for every index i, there can only be A before it and B after it
// i.e. count # of B before it and # of A after it => # of chars needed to be removed
// left[i] := # of B in s[0:i]
// right[i] := # of A in s[i:n-1]
func minimumDeletions(s string) int {
	n := len(s)

	left := make([]int, n)
	B := 0
	for i := 0; i < n; i++ {
		if s[i] == 'b' {
			B += 1
		}
		left[i] = B
	}

	right := make([]int, n)
	A := 0
	for i := n - 1; i >= 0; i-- {
		if s[i] == 'a' {
			A += 1
		}
		right[i] = A
	}

	res := min(left[n-1], right[0])
	for i := 1; i < n-1; i++ {
		res = min(res, left[i-1]+right[i+1])
	}
	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
