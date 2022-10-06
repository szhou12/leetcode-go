package leetcode

// DFS
// # levels = k
// # branches = 1~9
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var combo []int

	dfs(k, n, 1, combo, &res)
	return res
}

func dfs(k int, n int, num int, combo []int, res *[][]int) {
	// base case
	if len(combo) == k {
		if sum(combo) == n {
			comboCopy := make([]int, len(combo))
			copy(comboCopy, combo)
			*res = append(*res, comboCopy)
		}
		return
	}

	for i := num; i <= 9; i++ {
		combo = append(combo, i)
		dfs(k, n, i+1, combo, res)
		combo = combo[:len(combo)-1]
	}
}

func sum(nums []int) int {
	res := 0
	for _, val := range nums {
		res += val
	}
	return res
}
