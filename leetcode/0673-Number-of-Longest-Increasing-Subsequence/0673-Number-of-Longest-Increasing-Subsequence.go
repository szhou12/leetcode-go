package leetcode

func findNumberOfLIS(nums []int) int {
	curCombo := make([]int, 0)
	var res [][]int
	dfs(nums, 0, curCombo, &res)
	count := 0
	maxLen := 0
	for _, combo := range res {
		if len(combo) > maxLen {
			count = 1
			maxLen = len(combo)
		} else if len(combo) == maxLen {
			count++
		}
	}
	return count
}

func dfs(nums []int, index int, curCombo []int, res *[][]int) {
	c := make([]int, len(curCombo))
	copy(c, curCombo)
	*res = append(*res, c)

	for i := index; i < len(nums); i++ {
		if len(curCombo) == 0 || nums[i] > curCombo[len(curCombo)-1] {
			curCombo = append(curCombo, nums[i])
			dfs(nums, i+1, curCombo, res)
			curCombo = curCombo[:len(curCombo)-1]
		}
	}
}
