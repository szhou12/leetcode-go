package leetcode

import "strings"

func solveNQueens(n int) [][]string {
	var result [][]string
	var cur []int // cur = one possible way = a slice of col index that places Queen
	dfs(n, cur, &result)
	return result
}

func dfs(n int, cur []int, result *[][]string) {
	// base case
	if len(cur) == n {
		res := convert(cur, n)
		*result = append(*result, res)
		return
	}

	// # branches = # cols
	// # levels = # rows
	for col := 0; col < n; col++ {
		if valid(cur, col) {
			cur = append(cur, col)
			dfs(n, cur, result)
			cur = cur[:len(cur)-1]
		}
	}
}

func valid(cur []int, curCol int) bool {
	curRow := len(cur) // curRow = the row I'm about to place a Queen
	for preRow := 0; preRow < curRow; preRow++ {
		preCol := cur[preRow]
		if curCol == preCol || curCol-preCol == curRow-preRow || curCol-preCol == preRow-curRow {
			return false
		}
	}
	return true
}

func convert(cur []int, n int) []string {
	res := make([]string, 0)
	for _, colIndex := range cur {
		row := strings.Repeat(".", n)
		row = row[0:colIndex] + "Q" + row[colIndex+1:]
		res = append(res, row)
	}
	return res
}
