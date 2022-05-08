package leetcode

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int
	var combo []int

	dfs(candidates, target, 0, combo, &result)

	return result

}

func dfs(candidates []int, moneyLeft int, index int, combo []int, result *[][]int) {
	// base case
	if len(candidates) == index {
		if moneyLeft == 0 {
			res := convert(candidates, combo)
			*result = append(*result, res)
		}
		return
	}

	for n := 0; n*candidates[index] <= moneyLeft; n++ {
		combo = append(combo, n)
		dfs(candidates, moneyLeft-n*candidates[index], index+1, combo, result)
		combo = combo[:len(combo)-1]
	}
}

func convert(candidates []int, combo []int) []int {
	var res []int
	for i, n := range combo {
		for j := 0; j < n; j++ {
			res = append(res, candidates[i])
		}
	}
	return res
}
