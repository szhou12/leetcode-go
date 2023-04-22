package leetcode

func largestPathValue(colors string, edges [][]int) int {
	// Step 1: Reconstruct adj-list repr + Calculate degree
	n := len(colors)
	degree := make([]int, n)
	next := make([][]int, n)
	for i := 0; i < n; i++ {
		next[i] = make([]int, 0)
	}
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		next[from] = append(next[from], to)
		degree[to]++
	}

	// Step 2 + 3: Topological Sort X 26 times + find max colors
	res := 0
	for char := 'a'; char <= 'z'; char++ { // Note: char is rune type
		colorCount := topo(byte(char), colors, degree, next, n)
		if colorCount == -1 { // detect cycle
			return -1
		}
		res = max(res, colorCount)
	}

	return res
}

func topo(target byte, colors string, inDegree []int, next [][]int, n int) int {
	degree := make([]int, n)
	copy(degree, inDegree)
	count := make([]int, n) // count[i] = # of target colors along a path from start node to node i
	queue := make([]int, 0)

	// Start nodes: degree == 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			if colors[i] == target {
				count[i]++
			}
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		for _, nei := range next[cur] {
			// update count for nei
			if colors[nei] == target {
				count[nei] = max(count[nei], count[cur]+1)
			} else {
				count[nei] = max(count[nei], count[cur])
			}

			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Check if has cycle
	for i := 0; i < n; i++ {
		if degree[i] != 0 {
			return -1
		}
	}

	// Find the node that gives max count
	res := 0
	for i := 0; i < n; i++ {
		res = max(res, count[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
