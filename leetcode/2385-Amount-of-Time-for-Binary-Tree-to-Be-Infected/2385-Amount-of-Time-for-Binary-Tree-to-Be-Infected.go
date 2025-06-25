package leetcode

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func amountOfTime(root *TreeNode, start int) int {
	// step 1: convert tree to adj-list representation
	next := make(map[int]map[int]bool)
	getGraph(root, next)

	dist := make(map[int]int) // also works as visited
	for node, _ := range next {
		dist[node] = -1
	}

	// step 2: BFS
	queue := make([][]int, 0) // [[node, shortest time to reach]]

	// Start node
	queue = append(queue, []int{start, 0})
	dist[start] = 0 // mark as visited immediately once enqueued

	// Loop
	for len(queue) > 0 {
		// pop the current
		cur := queue[0]
		queue = queue[1:]
		node, t := cur[0], cur[1]

		// make the next move
		for nei, _ := range next[node] {
			// check if visited
			if dist[nei] != -1 {
				continue
			}

			queue = append(queue, []int{nei, t+1})
			dist[nei] = t + 1 // mark as visited immediately once enqueued
		}
	}

	res := -1
	for _, d := range dist {
		res = max(res, d)
	}

	return res
}

func getGraph(root *TreeNode, next map[int]map[int]bool) {
	// base case
	if root == nil {
		return
	}

	if _, ok := next[root.Val]; !ok {
		next[root.Val] = make(map[int]bool)
	}

	if root.Left != nil {
		next[root.Val][root.Left.Val] = true
		if _, ok := next[root.Left.Val]; !ok {
			next[root.Left.Val] = make(map[int]bool)
			next[root.Left.Val][root.Val] = true
		}
	}
	if root.Right != nil {
		next[root.Val][root.Right.Val] = true
		if _, ok := next[root.Right.Val]; !ok {
			next[root.Right.Val] = make(map[int]bool)
			next[root.Right.Val][root.Val] = true
		}
	}

	getGraph(root.Left, next)
	getGraph(root.Right, next)
}