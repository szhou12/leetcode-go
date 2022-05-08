package leetcode

func generateParenthesis(n int) []string {
	var result []string
	prefix := ""

	dfs(n, 0, 0, prefix, &result)
	return result

}

func dfs(n int, l int, r int, prefix string, result *[]string) {
	// base case
	if l == n && r == n {
		*result = append(*result, prefix)
		return
	}

	if l < n {
		prefix = prefix + "("
		dfs(n, l+1, r, prefix, result)
		prefix = prefix[:len(prefix)-1]
	}
	if l > r {
		prefix = prefix + ")"
		dfs(n, l, r+1, prefix, result)
		prefix = prefix[:len(prefix)-1]
	}
}
