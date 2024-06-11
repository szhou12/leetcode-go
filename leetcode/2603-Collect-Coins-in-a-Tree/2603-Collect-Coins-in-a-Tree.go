package leetcode

func collectTheCoins(coins []int, edges [][]int) int {
	n := len(coins)

	// Step 1: Reconstruct adj-list repr + Calculate degree
	degree := make([]int, n)
	next := make([]map[int]bool, n) // find all its neighbors for each node
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

	// Step 2: Prune leaf nodes with coins = 0
	deleted := prune(degree, next, coins, n)

	// Step 3: Mark depth
	depth := mark(degree, next, deleted, n)

	// Step 4: Only nodes with depth >= 3 are must-visit nodes
	m := 0
	for i := 0; i < n; i++ {
		if depth[i] >= 3 {
			m++
		}
	}
	// m == 0: all nodes have depth <= 2, meaning that pick any node as start, we are able to collect all coins without moving (0 edge needed)
	if m-1 >= 0 {
		return 2 * (m - 1)
	}
	return 0
}

/*
Topological Sort to "mark" the depth of each surviving node as its distance to the farthest leaf node
*/
func mark(degree []int, next []map[int]bool, deleted []int, n int) []int {
	depth := make([]int, n) // leaf node <=> depth = 1
	queue := make([]int, 0)

	// Start nodes
	for i := 0; i < n; i++ {
		if degree[i] == 1 && deleted[i] != 1 {
			depth[i] = 1
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			delete(next[nei], cur) // 这里代替了 check visited, 因为过河拆桥, 把从内层走回外围的edge给删了
			depth[nei] = max(depth[nei], depth[cur]+1) // depth要在“剥洋葱”方法下取最深的值
			if degree[nei] == 1 && deleted[nei] != 1 {
				queue = append(queue, nei)
			}
		}
	}

	return depth
}

/*
Topological Sort to "peel off" all leaf nodes and potential leaf nodes whose degree = 1 and coins = 0
*/
func prune(degree []int, next []map[int]bool, coins []int, n int) []int {
	deleted := make([]int, n)
	queue := make([]int, 0)

	// Start nodes
	for i := 0; i < n; i++ {
		if degree[i] == 1 && coins[i] == 0 {
			queue = append(queue, i)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update
		deleted[cur] = 1

		// Make the next move
		for nei, _ := range next[cur] {
			degree[nei]--
			delete(next[nei], cur)
			if degree[nei] == 1 && coins[nei] == 0 {
				queue = append(queue, nei)
			}
		}

	}

	return deleted
}
