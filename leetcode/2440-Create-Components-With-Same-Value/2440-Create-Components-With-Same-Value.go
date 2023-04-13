package leetcode

import "sort"

func componentValue(nums []int, edges [][]int) int {
	// Edge case: only one node
	if len(nums) == 1 {
		return 0
	}

	// Step 1: reconstruct adj-list repr + calculate degree
	n := len(nums)
	degree := make([]int, n)
	next := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		next[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		degree[a]++
		degree[b]++
		next[a][b] = true
		next[b][a] = true
	}

	// Calculate all possible component sum
	total := 0
	for _, val := range nums {
		total += val
	}
	sums := make([]int, 0)
	for s := 1; s*s <= total; s++ {
		if total%s == 0 {
			sums = append(sums, s)
			sums = append(sums, total/s)
		}
	}
	sort.Ints(sums)

	// Step 2: Topological Sort for each possible component sum
	for _, sum := range sums {
		if splitable(nums, degree, next, sum) {
			return total/sum - 1
		}
	}
	return 0

}

func splitable(nums []int, indegree []int, next []map[int]bool, targetSum int) bool {
	n := len(nums)
	// sum[i] = component sum for subtree with root = i
	sum := make([]int, len(nums))
	copy(sum, nums)

	degree := make([]int, len(indegree))
	copy(degree, indegree)

	visited := make([]bool, n)

	ok := true // return value: indicates if targetSum is split-able
	queue := make([]int, 0)

	// Start nodes: degree == 1
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
			visited[i] = true
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update
		if sum[cur] > targetSum {
			ok = false
			break
		} else if sum[cur] == targetSum {
			sum[cur] = 0 // meaning no need to percolate to higher-class node
		}
		// sum[cur] < targetSum: meaning need to percolate cur's subtree sum to higher-class node

		// Make the next move
		for nei, _ := range next[cur] {
			// Check if visited
			if visited[nei] {
				continue
			}
			sum[nei] += sum[cur]
			degree[nei]--
			if degree[nei] == 1 {
				queue = append(queue, nei)
				visited[nei] = true
			}
		}
	}

	return ok

}
