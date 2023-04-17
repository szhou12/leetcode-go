package leetcode

func maximumInvitations(favorite []int) int {
	// Step 1: Calculate degree (favorite is already adj-list repr)
	n := len(favorite)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		degree[favorite[i]]++ // 0 likes 2: 0 -> 2
	}

	// Step 2: Topological Sort - mark non-circular nodes + calc depth
	visited := make([]int, n) // 0: 环上节点; 1: 非环节点
	depth := make([]int, n)
	for i := 0; i < n; i++ {
		depth[i] = 1
	}
	topo(favorite, degree, visited, depth)

	// Step 3: max # of attending employees
	maxMultinaryRing := 0
	totalBinaryRing := 0
	// detect ring + calc size
	for i := 0; i < n; i++ {
		if visited[i] == 1 {
			continue
		}
		// a circle entry
		j := i
		count := 0
		for visited[j] == 0 {
			count++
			visited[j] = 1
			j = favorite[j]
		}
		// exit a circle

		if count >= 3 {
			maxMultinaryRing = max(maxMultinaryRing, count)
		} else if count == 2 {
			totalBinaryRing += depth[i] + depth[favorite[i]]
		}
	}

	return max(maxMultinaryRing, totalBinaryRing)
}

func topo(favorite []int, degree []int, visited []int, depth []int) {
	n := len(favorite)
	queue := make([]int, 0)

	// Start nodes: degree == 0
	for i := 0; i < n; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
			visited[i] = 1
			depth[i] = 1
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]

		// Make the next move
		nei := favorite[cur]
		degree[nei]--
		if degree[nei] == 0 {
			queue = append(queue, nei)
			visited[nei] = 1
		}
		depth[nei] = depth[cur] + 1 // 注意！depth更新在if之外，不能在if内，这是因为环上的节点也需要更新depth，而if内只有非环上节点才进得去
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
