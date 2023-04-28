package leetcode

func sequenceReconstruction(org []int, sequences [][]int) bool {
	n := len(org)

	// Step 0: Edge Cases
	if len(sequences) == 0 {
		return false
	}
	// if sequences contain element not in [1, n]
	for _, seq := range sequences {
		for _, num := range seq {
			if num < 1 || num > n {
				return false
			}
		}
	}

	// Step 1: Reconstruct adj-list + Calculate degree
	degree := make([]int, n+1)
	next := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		next[i] = make([]int, 0)
	}
	for _, seq := range sequences {
		for i := 0; i < len(seq)-1; i++ {
			from, to := seq[i], seq[i+1]
			next[from] = append(next[from], to)
			degree[to]++
		}
	}

	// Step 2: Topological Sort
	// Note: if we have unique reconstruction seq, then queue only has one element at any time
	queue := make([]int, 0)
	res := make([]int, 0)

	// Start nodes: degree == 0
	for i := 1; i <= n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Check if queue has ONLY one element
		if len(queue) > 1 {
			return false
		}
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update res
		res = append(res, cur)

		// Make the next move
		for _, nei := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Step 3: Check if res has all elements in org & if res has same order as org
	if len(res) != len(org) {
		return false
	}
	for i, num := range res {
		if num != org[i] {
			return false
		}
	}

	return true
}
