package leetcode

var dirs = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
}

func findWords(board [][]byte, words []string) []string {
	root := newTrieNode()
	// Step 1: Build Trie
	for _, word := range words {
		node := root

		for _, char := range word {
			letter := int(char - 'a')
			if node.children[letter] == nil {
				node.children[letter] = newTrieNode()
			}
			node = node.children[letter]
		}
		node.isEnd = true
	}

	// Step 2: DFS
	res := make([]string, 0)

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	appear := make(map[string]bool)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			letter := int(board[i][j] - 'a')
			if root.children[letter] == nil {
				continue
			}
			dfs(board, visited, i, j, root.children[letter], "", appear)
		}
	}

	for key := range appear {
		res = append(res, key)
	}

	return res
}

func dfs(board [][]byte, visited [][]bool, x int, y int, node *TrieNode, word string, appear map[string]bool) {
	// Base case
	if node == nil {
		return
	}

	visited[x][y] = true
	word += string(board[x][y])
	if node.isEnd {
		appear[word] = true
	}

	// Recursion
	m, n := len(board), len(board[0])
	for _, dir := range dirs {
		neiX, neiY := x+dir[0], y+dir[1]
		// check if out of bound
		if !(0 <= neiX && neiX < m && 0 <= neiY && neiY < n) {
			continue
		}
		// check if already visited
		if visited[neiX][neiY] {
			continue
		}
		next := int(board[neiX][neiY] - 'a')
		dfs(board, visited, neiX, neiY, node.children[next], word, appear)
	}

	// backtracking
	visited[x][y] = false
	// word = word[:len(word)-1] // never used

}

type TrieNode struct {
	isEnd    bool
	children [26]*TrieNode
}

func newTrieNode() *TrieNode {
	node := TrieNode{isEnd: false}

	for i := 0; i < 26; i++ {
		node.children[i] = nil
	}

	return &node
}
