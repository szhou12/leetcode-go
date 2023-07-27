package leetcode

import "math"

func minimizeConcatenatedLength(words []string) int {
	firstWord := words[0]
	start := int(firstWord[0] - 'a')
	end := int(firstWord[len(firstWord)-1] - 'a')

	memo := make([][][]int, len(words))
	for i := 0; i < len(words); i++ {
		memo[i] = make([][]int, 26)
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < 26; j++ {
			memo[i][j] = make([]int, 26)
		}
	}

	return len(firstWord) + dfs(1, start, end, words, memo)
}

// dfs(i, start, end) := min length when we about to join words[i] with concat string (head=start, tail=end)
// memo[i][start][end] := (i=len(words)) X (26 letters) X (26 letters)
func dfs(i int, start int, end int, words []string, memo [][][]int) int {
	n := len(words)

	// base case
	if i == n {
		// if we have no word to join, obviously concat string will be empty
		return 0
	}

	// pruning: if state(i, staart, end) already computed
	if memo[i][start][end] != 0 {
		return memo[i][start][end]
	}

	curSize := len(words[i])

	head := int(words[i][0] - 'a')
	tail := int(words[i][curSize-1] - 'a')

	res := math.MaxInt
	// add to the front: compare tail with start
	if tail == start {
		res = min(res, curSize-1+dfs(i+1, head, end, words, memo))
	} else {
		res = min(res, curSize+dfs(i+1, head, end, words, memo))
	}

	// add to the end: compare head with end
	if head == end {
		res = min(res, curSize-1+dfs(i+1, start, tail, words, memo))
	} else {
		res = min(res, curSize+dfs(i+1, start, tail, words, memo))
	}

	memo[i][start][end] = res

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
