package leetcode

func longestPrefix(s string) string {
	n := len(s)

	lsp := make([]int, n)

	// base case
	lsp[0] = 0

	// recurrence
	for i := 1; i < n; i++ {
		j := lsp[i-1]
		for j >= 1 && s[j] != s[i] {
			j = lsp[j-1]
		}
		if s[j] == s[i] {
			lsp[i] = j + 1
		} else {
			lsp[i] = j
		}
	}

	return s[:lsp[n-1]]
}
