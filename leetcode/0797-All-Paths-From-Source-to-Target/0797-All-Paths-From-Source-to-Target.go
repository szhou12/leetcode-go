package leetcode

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var combo []int

	combo = append(combo, 0)
	dfs(graph, 0, combo, &res)
	return res
}

func dfs(graph [][]int, index int, combo []int, res *[][]int) {
	// 只要到达 node n-1 就是一条合法路径
	if combo[len(combo)-1] == len(graph)-1 {
		c := make([]int, len(combo))
		copy(c, combo)
		*res = append(*res, c)
		return
	}

	for _, v := range graph[index] {
		combo = append(combo, v)
		dfs(graph, v, combo, res)
		combo = combo[:len(combo)-1]
	}

}
