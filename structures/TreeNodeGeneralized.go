package structures

type TreeNodeGeneralized struct {
	Val int
	Children []*TreeNodeGeneralized
}

// BFS to build a tree from edges where edges[i] = [ai, bi] indicates ai -- bi
// NOTE: [ai, bi] does not indicate order. In other words, there is no guarantee that ai is parent and bi is child
// Therefore, using BFS instead of DFS (Recursion) to build the tree can handle this case
func buildTreeBFS(edges [][]int) *TreeNodeGeneralized {
	// reconstruct adj-list repr
	next := make(map[int][]int)
	nodes := make(map[int]*TreeNodeGeneralized) // {key=node val, val=TreeNodeGeneralized}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)

		if _, ok := nodes[a]; !ok {
			nodes[a] = &TreeNodeGeneralized{Val: a}
		}
		if _, ok := nodes[b]; !ok {
			nodes[b] = &TreeNodeGeneralized{Val: b}
		}
	}

	queue := make([]int, 0)
	visited := make(map[int]bool)
	// start node
	queue = append(queue, 0)
	visited[0] = true

	// loop
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, nei := range next[cur] {
			if visited[nei] {
				continue
			}
			queue = append(queue, nei)
			visited[nei] = true
			nodes[cur].Children = append(nodes[cur].Children, nodes[nei])
		}
	}

	return nodes[0]
	
}