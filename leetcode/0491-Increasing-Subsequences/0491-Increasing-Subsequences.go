package leetcode

func findSubsequences(nums []int) [][]int {
	var res [][]int
	var subset []int
	dfs(nums, 0, subset, &res)
	return res
}

func dfs(nums []int, startIndex int, subset []int, res *[][]int) {
	// 无Base Case操作
	// 沿着Recursion tree路径, 只要遇到合法的subset (ie. len(subset) >= 2), 就加入res
	if len(subset) >= 2 {
		curCopy := make([]int, len(subset))
		copy(curCopy, subset)
		*res = append(*res, curCopy)
	}

	used := make(map[int]bool) // 当前层生成一个新Map
	for i := startIndex; i < len(nums); i++ {
		// 跳过当前元素的情况: OR关系
		// Case 1: 当前元素 < 当前子集中最后一个元素
		// Case 2: 当前元素 已经在当前层用过了
		if (len(subset) > 0 && nums[i] < subset[len(subset)-1]) || used[nums[i]] {
			continue
		}

		used[nums[i]] = true
		subset = append(subset, nums[i])
		dfs(nums, i+1, subset, res)
		subset = subset[:len(subset)-1]
	}
}
