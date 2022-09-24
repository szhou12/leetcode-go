package leetcode

func totalNQueens(n int) int {
	var res [][]int
	var cur []int // one solution - index = row idx, value = col idx to place a Queen

	dfs(n, cur, &res)
	return len(res)
}

// # levels = # rows
// # branches = # cols
func dfs(n int, cur []int, res *[][]int) {
	// Base case
	if len(cur) == n { // all rows have been filled up
		*res = append(*res, cur)
		return
	}

	// DFS: # branches = # cols
	for col := 0; col < n; col++ {
		if valid(cur, col) {
			cur = append(cur, col)
			dfs(n, cur, res)
			cur = cur[:len(cur)-1]
		}
	}

}

func valid(cur []int, curCol int) bool {
	curRow := len(cur)
	for prevRow := 0; prevRow < curRow; prevRow++ {
		prevCol := cur[prevRow]
		if prevCol == curCol || prevRow-curRow == prevCol-curCol || prevRow-curRow == curCol-prevCol {
			return false
		}
	}
	return true
}
