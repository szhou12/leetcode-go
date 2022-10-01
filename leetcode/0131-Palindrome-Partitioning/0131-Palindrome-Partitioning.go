package leetcode

// DFS
// # levels = len(s)
// # branches = [last cut ... end of string]
func partition(s string) [][]string {

	var res [][]string
	var combo []string
	dfs(s, 0, combo, &res)
	return res
}

func dfs(s string, index int, combo []string, res *[][]string) {
	// Base Case
	if index == len(s) {
		comboToAdd := make([]string, len(combo))
		copy(comboToAdd, combo)
		*res = append(*res, comboToAdd)
		return
	}

	// branches = last cut (index) ... end of string
	for i := index; i < len(s); i++ {
		if isValid(s, index, i) {
			combo = append(combo, s[index:i+1])
			dfs(s, i+1, combo, res)
			combo = combo[:len(combo)-1]
		}
	}
}

func isValid(s string, start int, end int) bool {
	left, right := start, end
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
