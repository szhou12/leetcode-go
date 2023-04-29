package leetcode

func alienOrder(words []string) string {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make(map[string]int)
	next := make(map[string]map[string]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			degree[word[i:i+1]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]

		// Check if word1 = abcd; word2 = abc
		if len(word1) > len(word2) && word1[:len(word2)] == word2 {
			return ""
		}

		for j := 0; j < min(len(word1), len(word2)); j++ {
			from, to := word1[j:j+1], word2[j:j+1]

			// Init next[from] if not existed
			if _, ok := next[from]; !ok {
				next[from] = make(map[string]bool)
			}

			if from == to {
				continue
			}
			if _, ok := next[from][to]; !ok {
				next[from][to] = true
				degree[to]++
			}
			break
		}
	}

	// Step 2: Topological Sort
	res := ""
	queue := make([]string, 0)
	// Start nodes: degree == 0
	for char, deg := range degree {
		if deg == 0 {
			queue = append(queue, char)
		}
	}
	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update res
		res += cur

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Check if has cycle
	if len(res) != len(degree) {
		return ""
	}

	return res

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
