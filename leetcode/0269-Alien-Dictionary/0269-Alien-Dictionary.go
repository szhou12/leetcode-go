package leetcode

import "strings"

func alienOrder(words []string) string {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make(map[byte]int)
	next := make(map[byte]map[byte]bool)
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			degree[word[i]] = 0
		}
	}
	for i := 0; i < len(words)-1; i++ {
		word1 := words[i]
		word2 := words[i+1]
		m, n := len(word1), len(word2)

		// Invalid if word1 = abcz; word2 = abc
		if m > n && word1[:n] == word2[:n] {
			return ""
		}

		for j := 0; j < min(m, n); j++ {
			from, to := word1[j], word2[j]

			// Init next[from] if not existed
			if _, ok := next[from]; !ok {
				next[from] = make(map[byte]bool)
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

	var sb strings.Builder

	// Step 2: Topological Sort
	queue := make([]byte, 0)
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
		sb.WriteByte(cur)

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	res := sb.String()
	// Check if has cycle by checking if res has all chars in degree[]
	if len(res) != len(degree) {
		return ""
	}

	return res

}
