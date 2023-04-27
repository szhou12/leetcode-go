package leetcode

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	// Step 0: Relabel -1 group elements
	nextGroupId := m
	for i := 0; i < n; i++ {
		if group[i] == -1 {
			group[i] = nextGroupId
			nextGroupId++
		}
	}

	// Step 1: Reconstruct adj-list repr + Calculate degree
	// 1a. next + degree within each group
	// nextPerGroup = {GroupId : {ItemId : [ItemIds]}}
	nextPerGroup := make(map[int]map[int][]int)
	// degreePerGroup only calculates in-degrees within a group
	// degreePerGroup = {GroupId : {ItemId : in-degree}}
	degreePerGroup := make(map[int]map[int]int)
	for i := 0; i < n; i++ {
		// Init
		if _, ok := nextPerGroup[group[i]]; !ok {
			nextPerGroup[group[i]] = make(map[int][]int)
		}
		if _, ok := degreePerGroup[group[i]]; !ok {
			degreePerGroup[group[i]] = make(map[int]int)
		}

		// item default degree = 0
		degreePerGroup[group[i]][i] = 0

		for _, from := range beforeItems[i] { // from -> i
			if group[from] == group[i] {
				nextPerGroup[group[from]][from] = append(nextPerGroup[group[from]][from], i)
				degreePerGroup[group[from]][i]++
			}
		}
	}
	// 1b. next + degree for groups
	// nextGroup = {GroupId : [GroupIds]}
	nextGroup := make(map[int][]int)
	// degreeGroup = {GroupId : in-degrees}
	degreeGroup := make(map[int]int)
	for i := 0; i < n; i++ {
		// Init
		if _, ok := nextGroup[group[i]]; !ok {
			nextGroup[group[i]] = make([]int, 0)
		}
		if _, ok := degreeGroup[group[i]]; !ok {
			degreeGroup[group[i]] = 0
		}

		for _, from := range beforeItems[i] {
			if group[from] != group[i] { // group[from] -> group[i]
				nextGroup[group[from]] = append(nextGroup[group[from]], group[i])
				degreeGroup[group[i]]++
			}
		}
	}

	// Step 2: Topological Sort
	// 2a. Topo sort within each group
	// sortPerGroup = {GroupId : [sorted items]}
	sortPerGroup := make(map[int][]int)
	for gid, _ := range nextPerGroup {
		cur := topo(nextPerGroup[gid], degreePerGroup[gid])
		if len(cur) == 0 {
			return []int{}
		}
		sortPerGroup[gid] = cur
	}

	// 2b. Topo sort groups
	sortGroup := topo(nextGroup, degreeGroup)
	if len(sortGroup) == 0 {
		return []int{}
	}

	// Step 3: Assemble
	res := make([]int, 0)
	for _, gid := range sortGroup {
		res = append(res, sortPerGroup[gid]...)
	}
	return res
}

func topo(next map[int][]int, degree map[int]int) []int {
	res := make([]int, 0)
	queue := make([]int, 0)

	// Start nodes: degree == 0
	for itemId, deg := range degree {
		if deg == 0 {
			queue = append(queue, itemId)
		}
	}

	// Loop
	for len(queue) > 0 {
		// Cur
		cur := queue[0]
		queue = queue[1:]
		// update sorted element
		res = append(res, cur)

		// Make the next move
		for _, nei := range next[cur] {
			degree[nei]--
			if degree[nei] == 0 {
				queue = append(queue, nei)
			}
		}
	}

	// Check if has cycle
	if len(res) != len(degree) {
		return []int{}
	}

	return res
}
