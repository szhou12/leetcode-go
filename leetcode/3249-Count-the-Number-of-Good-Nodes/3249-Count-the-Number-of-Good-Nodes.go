package leetcode

func countGoodNodes(edges [][]int) int {
	root := buildTreeBFS(edges)
	res := []int{0}
	dfs(root, res)
	return res[0]
}

// return the number of nodes in the subtree rooted at root
func dfs(root *TreeNode, res []int) int {
	// base case
	if root == nil {
		return 0
	}

	total := 0
	dedup := make(map[int]bool) // used to determine if each child branch has same amount of nodes
	// recursion: iterate over each child and get their nodes count
	for _, child := range root.Children {
		curCount := dfs(child, res)
		total += curCount
		dedup[curCount] = true
	}

	// indicates that each branch has equal number of nodes so current subtree is a good node
	if len(dedup) <= 1 {
		res[0]++
	}

	return total + 1
}

type TreeNode struct {
	Val      int
	Children []*TreeNode
}

func buildTreeBFS(edges [][]int) *TreeNode {
	// Step 1: reconstruct adj-list repr
	// NOTE: using map instead of []int bc we don't know total number of nodes n
	next := make(map[int][]int)
	nodes := make(map[int]*TreeNode) // {node idx : TreeNode, node idx : TreeNode, ...}
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		next[a] = append(next[a], b)
		next[b] = append(next[b], a)

		if _, ok := nodes[a]; !ok {
			nodes[a] = &TreeNode{Val: a}
		}
		if _, ok := nodes[b]; !ok {
			nodes[b] = &TreeNode{Val: b}
		}
	}

	queue := make([]int, 0)
	visited := make(map[int]bool)

	// Start node: node 0
	queue = append(queue, 0)
	visited[0] = true

	// Loop
	for len(queue) > 0 {
		// cur
		cur := queue[0]
		queue = queue[1:]

		// make the next move
		for _, nei := range next[cur] {
			if visited[nei] {
				continue
			}
			nodes[cur].Children = append(nodes[cur].Children, nodes[nei])
			queue = append(queue, nei)
			visited[nei] = true
		}
	}

	return nodes[0]

}
