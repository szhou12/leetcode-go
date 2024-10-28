package otheralgorithms

import "sort"


// Time compelxity = O(n!)
func allPermutationsDups(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums) // Sort the input to make it easier to skip duplicates

	visited := make(map[int]bool)

	var dfs func(path []int)
	dfs = func(path []int) {
		// base case
		if len(path) == len(nums) {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
		}

		// backtracking
		// always start from nums[0]
		for i := 0; i < len(nums); i++ {
			// skip used element
			if visited[i] {
				continue
			}
			// skip duplicates
			/* Explantion of !visited[i-1]:
			*   Enforce that we can only use nums[i] if nums[i-1] is already used in the current path.
			*   If we try to use nums[1] at the first position without first using nums[0], we would get duplicate permutations.
			*   i.e. 强制相同的元素order保持不变. 不允许nums[i]排在nums[i-1]前面的情况出现, 从而避免由于交换相同元素的位置导致的重复.
			*/
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			visited[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}

	dfs([]int{})

	return result
}
